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

func (ic *IdleChecker) Start() {
	ic.lastIdleTime = time.Now()
}

func (ic *IdleChecker) NewContext(parent context.Context) (context.Context, context.CancelFunc) {
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

func (ic *IdleChecker) Debounce() {
	ic.lastIdleTime = time.Now()
}
