package notify

import (
	"strings"
)

// Webhook delivery is configured from the environment (URL — which carries a
// secret token) and optionally from config.json (format, displayName,
// avatarUrl, msgtype — which are not secret). The sink is strictly opt-in:
// with EnvWebhookURL unset the wiring helper attaches nothing and the notifier
// behaves exactly as before.
const (
	// EnvWebhookURL holds the destination webhook URL. The URL itself embeds
	// the bridge token (e.g. the turt2live/matrix-appservice-webhooks hook
	// token), so it is sourced from the environment and never written to
	// config.json. Empty disables the sink.
	EnvWebhookURL = "PVYAI_NOTIFY_WEBHOOK_URL"
	// EnvWebhookSummary is an optional one-line run summary attached to every
	// payload (for example "nightly audit run").
	EnvWebhookSummary = "PVYAI_NOTIFY_WEBHOOK_SUMMARY"
	// EnvWebhookFormat overrides the text format: "plain" (default) or "html".
	// When set, takes precedence over config.json notify.webhook.format.
	EnvWebhookFormat = "PVYAI_NOTIFY_WEBHOOK_FORMAT"
	// EnvWebhookDisplayName is the sender name shown in Matrix. Takes
	// precedence over config.json notify.webhook.displayName.
	EnvWebhookDisplayName = "PVYAI_NOTIFY_WEBHOOK_DISPLAY_NAME"
	// EnvWebhookAvatarURL is the sender avatar URL shown in Matrix. Takes
	// precedence over config.json notify.webhook.avatarUrl.
	EnvWebhookAvatarURL = "PVYAI_NOTIFY_WEBHOOK_AVATAR_URL"
	// EnvWebhookMsgType is the Matrix message type: "" (normal), "notice", or
	// "emote". Takes precedence over config.json notify.webhook.msgtype.
	EnvWebhookMsgType = "PVYAI_NOTIFY_WEBHOOK_MSGTYPE"
)

// WebhookNotifyArgs carries the formatting options for the webhook sink. This
// mirrors config.WebhookNotify but lives in the notify package to avoid an
// import cycle (config imports notify for Mode/FocusMode validation).
type WebhookNotifyArgs struct {
	Format      string
	DisplayName string
	AvatarURL   string
	MsgType     string
}

// MaybeAddWebhookSink attaches a webhook sink to n when a webhook URL is present
// in the environment, and is otherwise a no-op. env resolves an environment
// variable (pass os.Getenv); logf records one redacted line per failed delivery
// (pass nil to stay silent — for example a TUI that owns the screen). It is safe
// to call unconditionally: configuration alone decides whether the sink exists.
//
// The attached sink is still subject to the notifier's Mode/focus policy, so a
// webhook only delivers when notifications are enabled (for example
// `--notify both`), matching the rest of the notification surface.
func MaybeAddWebhookSink(n *Notifier, env func(string) string, logf func(format string, args ...any)) {
	MaybeAddWebhookSinkFromConfig(n, WebhookNotifyArgs{}, env, logf)
}

// MaybeAddWebhookSinkFromConfig attaches a webhook sink using formatting options
// from the resolved config with environment variable overrides. The webhook
// URL is always sourced from the environment (EnvWebhookURL) because it carries
// a secret bridge token.
func MaybeAddWebhookSinkFromConfig(n *Notifier, cfg WebhookNotifyArgs, env func(string) string, logf func(format string, args ...any)) {
	if n == nil || env == nil {
		return
	}
	url := strings.TrimSpace(env(EnvWebhookURL))
	if url == "" {
		return
	}

	// Environment takes precedence over config for each formatting field.
	format := strings.TrimSpace(env(EnvWebhookFormat))
	if format == "" {
		format = cfg.Format
	}
	displayName := strings.TrimSpace(env(EnvWebhookDisplayName))
	if displayName == "" {
		displayName = cfg.DisplayName
	}
	avatarURL := strings.TrimSpace(env(EnvWebhookAvatarURL))
	if avatarURL == "" {
		avatarURL = cfg.AvatarURL
	}
	msgType := strings.TrimSpace(env(EnvWebhookMsgType))
	if msgType == "" {
		msgType = cfg.MsgType
	}

	n.AddSink(NewWebhookSink(WebhookConfig{
		URL:         url,
		Summary:     strings.TrimSpace(env(EnvWebhookSummary)),
		Format:      WebhookFormat(format),
		DisplayName: displayName,
		AvatarURL:   avatarURL,
		MsgType:     msgType,
		Logf:        logf,
	}))
}