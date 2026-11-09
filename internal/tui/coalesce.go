package tui

import (
	"sync"
	"time"

	tea "charm.land/bubbletea/v2"
)

// streamCoalesceInterval is roughly one 60fps frame. Assistant-text deltas that
// arrive within this window are merged into a single agentTextMsg, so the render
// rate decouples from the token rate: a fast local model (100+ tok/s) no longer
// forces 100+ full Update→View cycles (each re-parsing the growing markdown) per
// second. Rendering stays smooth regardless of provider speed.
const streamCoalesceInterval = 16 * time.Millisecond

// textCoalescer batches agentTextMsg deltas before forwarding them to the Bubble
// Tea program. Any OTHER message flushes the pending text first, so ordering
// between streamed prose and tool-call / reasoning / row / usage messages is
// preserved. The turn's final agentResponseMsg does not pass through here (it is
// a tea.Cmd return, not a sink message), but the model drops deltas whose runID
// is no longer active, so a flush that races just past end-of-turn is harmless.
//
// Sink messages originate from the single agent goroutine and so arrive
// serially; the only concurrent caller is the flush timer. The mutex guards the
// buffer/timer. The forward (program.Send) is called AFTER releasing the mutex,
// so a slow frame render in the Bubble Tea Update goroutine no longer blocks
// the agent goroutine or the timer goroutine. Ordering is preserved because the
// drain step (under the lock) still serializes the *decision* of what to
// forward; the forward itself is just delivery, and the buffer is empty after
// drain so no other caller can forward the same content.
type textCoalescer struct {
	forward func(tea.Msg) // downstream sink (external sink + program.Send)
	// afterFunc schedules fn to run after one frame interval and returns a
	// stoppable timer. Defaults to a real time.AfterFunc(streamCoalesceInterval, …);
	// tests swap in a controllable timer so flush timing is deterministic instead of
	// racing the 16ms wall clock.
	afterFunc func(fn func()) coalesceTimer

	mu    sync.Mutex
	buf   []byte
	runID int
	timer coalesceTimer
}

// coalesceTimer is the subset of *time.Timer the coalescer needs. Abstracted
// behind afterFunc so a test can substitute a timer it controls.
type coalesceTimer interface {
	Stop() bool
}

func newTextCoalescer(forward func(tea.Msg)) *textCoalescer {
	return &textCoalescer{
		forward: forward,
		afterFunc: func(fn func()) coalesceTimer {
			return time.AfterFunc(streamCoalesceInterval, fn)
		},
	}
}

// send is the coalescing entry point installed as the RuntimeMessageSink.
func (c *textCoalescer) send(msg tea.Msg) {
	// Fast path for text deltas: buffer under lock, forward outside lock.
	text, ok := msg.(agentTextMsg)
	if !ok {
		// Non-text message: flush buffered text first (preserving order), then
		// forward the non-text message. Drain under lock, forward outside lock.
		c.mu.Lock()
		pending := c.drainLocked()
		c.mu.Unlock()
		if pending != nil {
			c.forward(pending)
		}
		c.forward(msg)
		return
	}

	c.mu.Lock()
	// A delta for a different run than the one buffered: flush the old run's text
	// before buffering the new run's. In practice runs are sequential (the prior
	// run's end already flushed via a non-text message), so this is belt-and-braces.
	if len(c.buf) > 0 && text.runID != c.runID {
		c.drainAndForwardLocked()
	}
	c.runID = text.runID
	c.buf = append(c.buf, text.delta...)
	if c.timer == nil {
		c.timer = c.afterFunc(c.flush)
	}
	c.mu.Unlock()
}

// flush forwards any buffered text as one agentTextMsg. Runs on the timer
// goroutine. Drains under the lock, then forwards outside the lock so a slow
// program.Send doesn't block the agent goroutine from acquiring the mutex.
func (c *textCoalescer) flush() {
	c.mu.Lock()
	pending := c.drainLocked()
	c.mu.Unlock()
	if pending != nil {
		c.forward(pending)
	}
}

// drainLocked flushes any buffered text as one agentTextMsg, stops the timer,
// and returns the message. Called with c.mu held. Returns nil when nothing is
// buffered. The caller is responsible for unlocking c.mu before forwarding.
func (c *textCoalescer) drainLocked() tea.Msg {
	if len(c.buf) == 0 {
		return nil
	}
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	msg := agentTextMsg{runID: c.runID, delta: string(c.buf)}
	c.buf = c.buf[:0]
	return msg
}

// drainAndForwardLocked is retained for the send() path where a run-ID change
// requires an immediate inline flush. Called with c.mu held; forwards while
// still holding the lock (acceptable here because this is a rare run-change
// edge case, not the hot path).
func (c *textCoalescer) drainAndForwardLocked() {
	if len(c.buf) == 0 {
		return
	}
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	msg := agentTextMsg{runID: c.runID, delta: string(c.buf)}
	c.buf = c.buf[:0]
	c.forward(msg)
}
