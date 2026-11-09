package tui

import (
	"context"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/term"
)

// forwarderBufferSize is the capacity of the buffered channel that decouples
// the agent goroutine from the Bubble Tea Update goroutine. 256 messages is
// enough to absorb a burst of tool-call + text messages during a frame render
// without blocking the producer. Non-text messages are never dropped; text
// deltas can be coalesced by the textCoalescer before reaching the forwarder,
// so the effective burst capacity is higher than 256 raw tokens.
const forwarderBufferSize = 256

// Run starts the PVYai Bubble Tea shell and returns a process-style exit code.
func Run(ctx context.Context, options Options) int {
	// The interactive shell needs a real terminal on stdin: with piped or
	// redirected input Bubble Tea blocks forever waiting for events that never
	// arrive (e.g. `echo "" | pvyai`). Fail fast with guidance toward the headless
	// path instead of hanging. term.IsTerminal is a true TTY check (it rejects
	// pipes, regular files, and non-terminal char devices like /dev/null) and
	// fails closed — anything that is not a verified terminal blocks the shell.
	if !term.IsTerminal(os.Stdin.Fd()) {
		fmt.Fprintln(os.Stderr, "pvyai: the interactive shell needs a terminal (stdin is not a TTY). For non-interactive use, run: pvyai exec \"<prompt>\"")
		return 2
	}

	externalSink := options.RuntimeMessageSink
	var program *tea.Program

	// Dedicated forwarder goroutine: decouples the agent goroutine (and the
	// coalescer's timer goroutine) from program.Send. When the Bubble Tea Update
	// goroutine is busy rendering a slow frame, program.Send blocks — but now it
	// blocks the forwarder goroutine, not the agent goroutine. The agent can
	// keep producing tokens and appending to the coalescer buffer; the forwarder
	// drains the backlog when Update frees up.
	//
	// The channel is buffered so short bursts (tool-call + text + tool-result
	// arriving in quick succession) don't block even if the forwarder is
	// momentarily stuck on a slow program.Send. If the channel fills up, the
	// producer blocks — but that's correct back-pressure: it means the TUI is
	// genuinely behind and slowing down the agent is better than dropping
	// messages (which could drop a permissionRequestMsg and deadlock the session).
	//
	// CRITICAL INVARIANT: permissionRequestMsg must NEVER be dropped. Because
	// the channel is unbounded in practice (buffer 256 + blocking send), every
	// message is guaranteed delivery. The only "loss" is coalesced text deltas
	// which are merged by the textCoalescer before reaching this channel.
	forwardCh := make(chan tea.Msg, forwarderBufferSize)
	go func() {
		for msg := range forwardCh {
			if externalSink != nil {
				externalSink(msg)
			}
			if program != nil {
				program.Send(msg)
			}
		}
	}()

	forward := func(msg tea.Msg) {
		forwardCh <- msg
	}

	// Coalesce streamed assistant-text deltas to ~one frame each so a fast provider
	// can't drive a full Update→View per token; every other message flushes pending
	// text first, keeping order intact.
	options.RuntimeMessageSink = newTextCoalescer(forward).send
	options.AltScreen = useAltScreen(options)

	programOpts := []tea.ProgramOption{
		tea.WithContext(ctx),
		tea.WithInput(os.Stdin),
		tea.WithOutput(os.Stdout),
		tea.WithFilter(mouseEventFilter()),
	}
	// Honor the no-color.org spec ourselves: NO_COLOR set to ANY non-empty value
	// disables color. bubbletea/colorprofile only respects strconv.ParseBool-style
	// values, so NO_COLOR=yes / NO_COLOR=foo would otherwise leave the UI in full
	// color. Force the Ascii (no-color, bold-still-allowed) profile. (AUDIT-M3)
	if noColorRequested(os.Getenv) {
		programOpts = append(programOpts, tea.WithColorProfile(colorprofile.Ascii))
	}
	initialModel := newModel(ctx, options)
	if initialModel.wantsMouseCapture() {
		initialModel.mouseCapture = true
	}
	program = tea.NewProgram(initialModel, programOpts...)

	if _, err := program.Run(); err != nil {
		// Surface the failure: exiting 1 with zero diagnostics left users
		// guessing why the default chat surface died.
		fmt.Fprintln(os.Stderr, "pvyai: tui error:", err)
		return 1
	}

	// Clean shutdown: close the forwarder channel so the goroutine exits.
	close(forwardCh)
	return 0
}

func useAltScreen(_ Options) bool {
	return true
}

// noColorRequested reports whether the no-color.org spec is in effect: NO_COLOR set
// to any non-empty value. Checked here rather than via the colorprofile dependency,
// whose strconv.ParseBool gate ignores common values like NO_COLOR=yes. (AUDIT-M3)
func noColorRequested(getenv func(string) string) bool {
	return getenv("NO_COLOR") != ""
}
