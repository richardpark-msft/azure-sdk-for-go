package internal

import (
	"context"
	"time"

	"github.com/jpillora/backoff"
)

// TODO: we should discuss what a common policy would be.
var DefaultRetryPolicy = RetryPolicy{
	Backoff: backoff.Backoff{
		Factor: 1,
		Min:    time.Second * 5,
	},
	MaxRetries: 5,
}

// Encapsulates a backoff policy, which allows you to configure the amount of
// time in between retries as well as the maximum retries allowed (via MaxRetries)
// NOTE: this should be copied by the caller as it is stateful.
type RetryPolicy struct {
	Backoff    backoff.Backoff
	MaxRetries int

	tries int
	Err   error
}

// Try marks an attempt to call (first call to Try() does not sleep).
// Will return false if the `ctx` is cancelled or if we exhaust our retries.
//
//    rp := RetryPolicy{Backoff:defaultBackoffPolicy, MaxRetries:5}
//
//    for rp.Try(ctx, errors.New("retries exhausted, can't do something")) {
//       <your code>
//    }
func (rp *RetryPolicy) Try(ctx context.Context, errOnExhaustion error) bool {
	defer func() { rp.tries++ }()

	if rp.tries == 0 {
		// first 'try' is always free
		return true
	}

	if rp.tries >= rp.MaxRetries {
		rp.Err = errOnExhaustion
		return false
	}

	select {
	case <-time.After(rp.Backoff.Duration()):
		rp.Err = nil
		return true
	case <-ctx.Done():
		rp.Err = ctx.Err()
		return false
	}
}
