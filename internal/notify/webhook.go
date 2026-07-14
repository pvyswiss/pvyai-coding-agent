package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pvyswiss/pvyai-coding-agent/internal/redaction"
)

// defaultWebhookTimeout bounds a single delivery attempt. A notification must
// never stall a finished run, so an unresponsive endpoint is abandoned quickly.
const defaultWebhookTimeout = 10 * time.Second

// maxWebhookBodyBytes caps how much of an error response we read back for the
// log line, so a hostile or chatty endpoint cannot flood the audit trail.
const maxWebhookBodyBytes = 2 << 10 // 2 KiB

// WebhookFormat selects how the "text" field is rendered in the webhook
// payload. The turt2live/matrix-appservice-webhooks bridge (used by PVYmessenger
// via whapi) accepts "plain" or "html"; Slack incoming webhooks ignore the
// field and render "text" as plain text.
type WebhookFormat string

const (
	// FormatPlain sends the text field as-is (default, backward-compatible).
	FormatPlain WebhookFormat = "plain"
	// FormatHTML wraps the text in minimal HTML tags for Matrix rendering.
	FormatHTML WebhookFormat = "html"
)

// WebhookLink is an optional labeled URL attached to a notification (for
// example a link to the CI run, the opened PR, or the session log).
type WebhookLink struct {
	Label string `json:"label,omitempty"`
	URL   string `json:"url"`
}

// webhookPayload is the JSON body POSTed to the configured webhook. It is
// shaped to be useful to a generic consumer while remaining renderable by a
// Slack incoming webhook (which reads the top-level "text" field) and by the
// turt2live/matrix-appservice-webhooks bridge (which reads "text", "format",
// "displayName", "avatarUrl", "msgtype", and "emoji"). The structured fields
// (type/message/summary/links) carry machine-readable detail ignored by both.
type webhookPayload struct {
	// Text is the human-readable line Slack/Matrix renders in the channel.
	Text string `json:"text"`
	// Format is "plain" or "html" (Matrix bridge only; Slack ignores it).
	Format string `json:"format,omitempty"`
	// DisplayName is the sender name shown in Matrix (Matrix bridge only).
	DisplayName string `json:"displayName,omitempty"`
	// AvatarURL is the sender avatar URL shown in Matrix (Matrix bridge only).
	AvatarURL string `json:"avatarUrl,omitempty"`
	// MsgType is the Matrix message type: "" (normal), "notice", or "emote".
	MsgType string `json:"msgtype,omitempty"`
	// Emoji controls auto-emoji conversion in the Matrix bridge (nil = default).
	Emoji *bool `json:"emoji,omitempty"`
	// Type is the machine-readable event kind ("completion", "awaiting_input").
	Type string `json:"type"`
	// Message is the notification body (already redacted).
	Message string `json:"message"`
	// Summary is an optional one-line run summary (already redacted).
	Summary string `json:"summary,omitempty"`
	// Links are optional labeled URLs (run page, PR, logs).
	Links []WebhookLink `json:"links,omitempty"`
}

// WebhookConfig configures a WebhookSink. URL is a Slack incoming-webhook URL,
// a turt2live/matrix-appservice-webhooks bridge URL, or any generic endpoint
// that accepts a JSON POST. The zero value (empty URL) yields an inert sink
// whose Emit is a no-op, so callers can wire a sink unconditionally and let
// configuration decide whether it fires.
type WebhookConfig struct {
	// URL is the destination. Empty disables the sink.
	URL string
	// Summary is an optional run summary attached to every emitted payload.
	Summary string
	// Format controls how the text field is rendered: "plain" (default) or
	// "html" (enables HTML formatting for the Matrix bridge).
	Format WebhookFormat
	// DisplayName is the sender name shown in Matrix (Matrix bridge only).
	DisplayName string
	// AvatarURL is the sender avatar URL shown in Matrix (Matrix bridge only).
	AvatarURL string
	// MsgType is the Matrix message type: "" (normal), "notice", or "emote".
	MsgType string
	// Links are optional labeled URLs attached to every emitted payload.
	Links []WebhookLink
	// Client is the HTTP client used for delivery. When nil a client with a
	// conservative timeout is used; its default transport honors HTTP(S)_PROXY
	// when a proxy is configured.
	Client *http.Client
	// Logf records a single line per failed delivery. Lines are passed through
	// the repo redaction before being written, so a token in the URL or message
	// is never logged in the clear. When nil, failures are silent.
	Logf func(format string, args ...any)
	// ExtraSecrets are literal values (for example the resolved API key) that must
	// be masked in the payload and logs in addition to the built-in patterns.
	ExtraSecrets []string
}

// WebhookSink delivers notifications to a webhook endpoint (Slack incoming
// webhook, turt2live/matrix-appservice-webhooks bridge, or generic HTTP POST).
// It implements Sink. Delivery is best-effort and fails soft: a non-2xx
// response or a transport error is logged (redacted) and swallowed so it can
// never disrupt the run.
type WebhookSink struct {
	url         string
	summary     string
	format      WebhookFormat
	displayName string
	avatarURL   string
	msgType     string
	links       []WebhookLink
	client      *http.Client
	logf        func(format string, args ...any)
	secrets     []string
}

// NewWebhookSink builds a WebhookSink from cfg. A blank URL produces an inert
// sink (Emit is a no-op).
func NewWebhookSink(cfg WebhookConfig) *WebhookSink {
	client := cfg.Client
	if client == nil {
		client = &http.Client{Timeout: defaultWebhookTimeout}
	}
	return &WebhookSink{
		url:         strings.TrimSpace(cfg.URL),
		summary:     cfg.Summary,
		format:      cfg.Format,
		displayName: cfg.DisplayName,
		avatarURL:   cfg.AvatarURL,
		msgType:     cfg.MsgType,
		// Copy the caller-owned slices so later mutation of cfg can neither
		// change a future payload nor race with a concurrent Emit.
		links:   append([]WebhookLink(nil), cfg.Links...),
		client:  client,
		logf:    cfg.Logf,
		secrets: append([]string(nil), cfg.ExtraSecrets...),
	}
}

// Emit builds the JSON payload for event/message and POSTs it to the webhook.
// It is fire-and-forget: every failure path logs (redacted) and returns nil-ish
// rather than propagating, satisfying the Sink contract that a notification
// must never crash the run.
func (s *WebhookSink) Emit(event Event, message string) {
	if s == nil || s.url == "" {
		return
	}

	// The webhook URL itself can carry a secret token, so treat it as a secret
	// when redacting outbound fields (matching log()), not just s.secrets.
	options := redaction.Options{ExtraSecretValues: append([]string{s.url}, s.secrets...)}
	safeMessage := redaction.RedactString(message, options)
	safeSummary := redaction.RedactString(s.summary, options)
	safeAvatar := redaction.RedactString(s.avatarURL, options)
	links := s.redactLinks(options)

	payload := webhookPayload{
		Text:        s.text(event, safeMessage),
		Format:      string(s.format),
		DisplayName: s.displayName,
		AvatarURL:   safeAvatar,
		MsgType:     s.msgType,
		Type:        eventType(event),
		Message:     safeMessage,
		Summary:     safeSummary,
		Links:       links,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		// A payload built from strings should always marshal; guard anyway.
		s.log("notify: webhook payload encode failed: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultWebhookTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, s.url, bytes.NewReader(body))
	if err != nil {
		s.log("notify: webhook request build failed: %v", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := s.client.Do(request)
	if err != nil {
		s.log("notify: webhook delivery failed: %v", err)
		return
	}
	defer func() { _ = response.Body.Close() }()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		snippet := readSnippet(response.Body)
		s.log("notify: webhook returned %d %s", response.StatusCode, snippet)
		return
	}
	// Drain the success body so the connection can be reused.
	_, _ = io.Copy(io.Discard, io.LimitReader(response.Body, maxWebhookBodyBytes))
}

// text renders the channel-facing line for an event. When format is "html" the
// message is wrapped in minimal HTML tags for the Matrix bridge.
func (s *WebhookSink) text(event Event, safeMessage string) string {
	if s.format == FormatHTML {
		return s.htmlText(event, safeMessage, safeSummaryFor(s, event))
	}
	if strings.TrimSpace(safeMessage) != "" {
		return safeMessage
	}
	return DefaultMessage(event)
}

// htmlText renders an HTML-formatted message for the Matrix bridge. The
// message body is HTML-escaped to prevent injection, then wrapped in <b> tags.
// When a summary is present it is appended as an <i> line.
func (s *WebhookSink) htmlText(event Event, message, summary string) string {
	var b strings.Builder
	if strings.TrimSpace(message) != "" {
		b.WriteString("<b>")
		b.WriteString(html.EscapeString(message))
		b.WriteString("</b>")
	} else {
		b.WriteString(htmlDefaultMessage(event))
	}
	if strings.TrimSpace(summary) != "" {
		b.WriteString("<br><i>")
		b.WriteString(html.EscapeString(summary))
		b.WriteString("</i>")
	}
	return b.String()
}

// htmlDefaultMessage returns the HTML version of DefaultMessage.
func htmlDefaultMessage(event Event) string {
	if event == AwaitingInput {
		return "<b>PVYai: needs input</b>"
	}
	return "<b>PVYai: ready</b>"
}

// safeSummaryFor returns the redacted summary for the sink, or empty string.
func safeSummaryFor(s *WebhookSink, _ Event) string {
	if s == nil {
		return ""
	}
	return s.summary
}

// redactLinks returns a copy of the configured links with their URLs and labels
// passed through redaction (a URL may carry a token in its query/userinfo).
func (s *WebhookSink) redactLinks(options redaction.Options) []WebhookLink {
	if len(s.links) == 0 {
		return nil
	}
	out := make([]WebhookLink, 0, len(s.links))
	for _, link := range s.links {
		out = append(out, WebhookLink{
			Label: redaction.RedactString(link.Label, options),
			URL:   redaction.RedactString(link.URL, options),
		})
	}
	return out
}

// log writes one redacted line. The fully-formatted message is redacted as a
// whole so a token appearing in any argument (URL, response body, error) is
// masked before it reaches the logger.
func (s *WebhookSink) log(format string, args ...any) {
	if s.logf == nil {
		return
	}
	line := fmt.Sprintf(format, args...)
	options := redaction.Options{ExtraSecretValues: append([]string{s.url}, s.secrets...)}
	s.logf("%s", redaction.RedactString(line, options))
}

// eventType maps an Event to its stable machine-readable string.
func eventType(event Event) string {
	switch event {
	case AwaitingInput:
		return "awaiting_input"
	default:
		return "completion"
	}
}

// readSnippet reads a bounded, single-line snippet of an error response body for
// logging. The caller redacts the result.
func readSnippet(body io.Reader) string {
	data, _ := io.ReadAll(io.LimitReader(body, maxWebhookBodyBytes))
	snippet := strings.TrimSpace(string(data))
	snippet = strings.ReplaceAll(snippet, "\n", " ")
	snippet = strings.ReplaceAll(snippet, "\r", " ")
	return snippet
}
