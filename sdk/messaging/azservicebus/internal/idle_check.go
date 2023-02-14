package internal

import (
	"context"
	"time"
)

type IdleChecker struct {
	lastIdleTime time.Time
	maxIdleTime  time.Duration
}

func NewIdleChecker(max time.Duration) *IdleChecker {
	return &IdleChecker{
		maxIdleTime: max,
	}
}

var zeroTime = time.Time{}

func (ic *IdleChecker) reset() {
	ic.lastIdleTime = zeroTime
}

// NewContext wraps the user's context in a context with an idle timer. The deadline
// for the wrapper context is based on how long the link has been idle.
func (ic *IdleChecker) NewContext(parent context.Context) (context.Context, context.CancelFunc) {
	if ic.lastIdleTime == zeroTime {
		ic.lastIdleTime = time.Now()
	}

	expiresOn, hasDeadline := parent.Deadline()
	end := ic.lastIdleTime.Add(ic.maxIdleTime)

	if !hasDeadline {
		// they set an infinite context, no deadline to take into account but our own
		return context.WithDeadline(parent, end)
	} else {
		// take the earliest interval so we don't stomp on their
		// deadline.
		if end.After(expiresOn) {
			return context.WithDeadline(parent, expiresOn)
		} else {
			return context.WithDeadline(parent, end)
		}
	}
}

// UpdateAndCheck checks whether the parent context cancelled or if our idle
// context (created through [NewContext]) was cancelled. If the link
// is idle it returns true.
//
// The internal idle timer will be reset if:
//   - returnedErr is not a cancel error.
//   - returnedErr is a cancellation error _and_ it was due to our
//     idle context, not the userCtx.
func (ic *IdleChecker) UpdateAndCheck(userCtx context.Context, returnedErr error) bool {
	if returnedErr == nil {
		ic.reset()
		return false
	}

	if !IsCancelError(returnedErr) {
		// if any other error, aside from cancellation, we can reset since we did
		// get a response/indicator for the receiver, even if it's just dead.
		ic.reset()
		return false
	}

	// if the user's context was cancelled that just indicates that
	// they cancelled or timed out. However, if _our_ code returns a
	// cancellation error then our idle timer fired.
	isIdle := userCtx.Err() == nil

	if isIdle {
		// set it so the next check of the idle timer starts it up.
		ic.reset()
	}

	return isIdle
}
