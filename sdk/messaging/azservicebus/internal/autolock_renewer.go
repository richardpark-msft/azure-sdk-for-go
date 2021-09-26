package internal

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

// renewLockFunc should match MgmtClient.RenewLocks
type renewLockFunc func(ctx context.Context, linkName string, lockTokens ...*amqp.UUID) (err error)

type LockRenewer struct {
	associatedLinkName string

	mu         sync.RWMutex
	cancellers map[amqp.UUID]func()

	// typically MgmtClient.RenewLocks. For testing.
	renewLockFunc renewLockFunc
	notifyError   func(err error)
}

func NewLockRenewer(associatedLinkName string, renewLockFunc renewLockFunc, notifyError func(err error)) *LockRenewer {
	return &LockRenewer{
		associatedLinkName: associatedLinkName,
		cancellers:         map[amqp.UUID]func(){},
		renewLockFunc:      renewLockFunc,
		notifyError:        notifyError,
	}
}

type RenewableMessage interface {
	LockToken() string
}

// Update should be called if there is a link change.
// TODO: not sure if the associated link name should be changed (or if it even matters)
// if we've detached...
func (r *LockRenewer) Update(renewLockFunc renewLockFunc, associatedLinkName string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.renewLockFunc = renewLockFunc
	r.associatedLinkName = associatedLinkName
}

func (r *LockRenewer) get() (renewLockFunc, string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.renewLockFunc, r.associatedLinkName
}

func (r *LockRenewer) Renew(ctx context.Context, m *amqp.Message) (func(), error) {
	lockToken, err := GetLockToken(m)
	initialExpirationTime := GetLockExpirationTime(m)

	if err != nil || initialExpirationTime.IsZero() {
		// no-op if they ask us to renew a lock on a message that doesn't have a
		// lock renewal token or an expiration date
		return nil, errors.New("no lock token, or invalid expiration time")
	}

	// TODO: would it make more sense to just say it has to be a specific UUID*, rather than
	// comparing by value?
	ctx, cancel := context.WithCancel(ctx)
	r.add(lockToken, cancel)

	if err := r.renew(ctx, lockToken, initialExpirationTime); err != nil {
		return nil, err
	}

	return cancel, nil
}

func (r *LockRenewer) remove(lockToken *amqp.UUID) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.cancellers, *lockToken)
}

func (r *LockRenewer) add(lockToken *amqp.UUID, cancel func()) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cancellers[*lockToken] = cancel
}

func (r *LockRenewer) renew(ctx context.Context, lockToken *amqp.UUID, initialExpirationTime time.Time) error {
	// we don't have the actual lock time from the service but we can
	// guess at it if we assume we're doing this pretty quickly after
	// receiving it.
	maxApproxLockDuration := time.Until(initialExpirationTime)
	defer r.remove(lockToken)

	// we'll reset the timer after this so just picking some arbitrary "not going to fire off time"
	timer := time.NewTimer(time.Hour * 24)
	defer timer.Stop()

	if err := r.renewLockAndScheduleNext(ctx, lockToken, maxApproxLockDuration, timer); err != nil {
		return err
	}

	go func() {
	Loop:
		for {
			select {
			case <-ctx.Done():
				break Loop
			case <-timer.C:
				if err := r.renewLockAndScheduleNext(ctx, lockToken, maxApproxLockDuration, timer); err != nil {
					r.notifyError(err)
				}
			}
		}
	}()

	return nil
}

func (r *LockRenewer) renewLockAndScheduleNext(ctx context.Context,
	lockToken *amqp.UUID,
	maxDuration time.Duration,
	timer *time.Timer) error {
	ctx, span := tab.StartSpan(ctx, "sb.renewlock")
	defer span.End()

	span.AddAttributes(
		tab.StringAttribute("message-id", lockToken.String()),
	)

	renewLockFunc, associatedLinkName := r.get()

	if err := renewLockFunc(ctx, associatedLinkName, lockToken); err != nil {
		span.Logger().Error(err)

		if isNonRetryableError(err) {
			// no point in continuing to renew lock. Something else has killed the ability to renew
			// the message.
			span.Logger().Fatal("renewlocks had a fatal error")
			span.End()
			return err
		}
	}

	// calculate the next expiration time - we'll just do it at half the interval
	// to give us a potential buffer against failure/clock drift.

	// setup the next renewal
	timer.Reset(maxDuration / 2)

	span.AddAttributes(
		tab.StringAttribute("next-renewal", time.Now().Add(maxDuration/2).String()),
	)

	timer.Reset(maxDuration / 2)
	return nil
}

func isNonRetryableError(err error) bool {
	// TODO: after some of the PR merges are done we'll have this in a centralized spot
	// and this can be deleted
	switch asType := err.(type) {
	case *net.OpError:
		return asType.Temporary() || asType.Timeout()
	default:
		// there are some other cases, like if the SB is throttled.
		// we can handle that more gracefully (and will when I throw this code away)
		return false
	}
}
