// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"time"

	"github.com/jpillora/backoff"
)

var DefaultFiniteRetrier = NewBackoffRetrier(BackoffRetrierParams{
	MaxRetries: 3,
	Factor:     2,
	Min:        time.Second,
	Max:        time.Minute,
	Jitter:     true,
})

var DefaultInfiniteRetrier = NewUnlimitedCyclingRetrier(UnlimitedCyclingRetrierOptions{
	Factor: 2,
	Min:    time.Second,
	Max:    time.Minute,
	Jitter: true,
})

// A retrier that allows you to do a basic for loop and get backoff
// and retry limits. See `Try` for more details on how to use it.
type Retrier interface {
	// Copies the retrier. Retriers are stateful and must be copied
	// before starting a set of retries.
	Copy() Retrier

	// CurrentTry is the current try (0 for the first run before retries)
	CurrentTry() int

	// Try marks an attempt to call (first call to Try() does not sleep).
	// Will return false if the `ctx` is cancelled or if we exhaust our retries.
	//
	//    rp := RetryPolicy{Backoff:defaultBackoffPolicy, MaxRetries:5}
	//
	//    for rp.Try(ctx) {
	//       <your code>
	//    }
	Try(ctx context.Context) bool
}

type unlimitedCyclingRetrier struct {
	backoff backoff.Backoff
	tries   int
}

func (r *unlimitedCyclingRetrier) Copy() Retrier {
	return &unlimitedCyclingRetrier{
		backoff: r.backoff,
		tries:   0,
	}
}

func (r *unlimitedCyclingRetrier) CurrentTry() int {
	return r.tries
}

func (r *unlimitedCyclingRetrier) Try(ctx context.Context) bool {
	defer func() { r.tries++ }()

	duration := r.backoff.Duration()

	select {
	case <-time.After(duration):
		if duration >= r.backoff.Max {
			r.backoff.Reset()
		}
		return true
	case <-ctx.Done():
		return false
	}
}

type UnlimitedCyclingRetrierOptions struct {
	Factor float64
	// Jitter eases contention by randomizing backoff steps
	Jitter bool
	// Min and Max are the minimum and maximum values of the counter
	Min, Max time.Duration
}

// NewUnlimitedCyclingRetrier creates a retrier which infinitely retries, but uses the
// backoff pol
func NewUnlimitedCyclingRetrier(options UnlimitedCyclingRetrierOptions) Retrier {
	return &unlimitedCyclingRetrier{
		backoff: backoff.Backoff{
			Factor: options.Factor,
			Jitter: options.Jitter,
			Min:    options.Min,
			Max:    options.Max,
		},
	}
}

// Encapsulates a backoff policy, which allows you to configure the amount of
// time in between retries as well as the maximum retries allowed (via MaxRetries)
// NOTE: this should be copied by the caller as it is stateful.
type backoffRetrier struct {
	backoff    backoff.Backoff
	MaxRetries int

	tries int
}

// BackoffRetrierParams are parameters for NewBackoffRetrier.
type BackoffRetrierParams struct {
	// MaxRetries is the maximum number of tries (after the first attempt)
	// that are allowed.
	MaxRetries int
	// Factor is the multiplying factor for each increment step
	Factor float64
	// Jitter eases contention by randomizing backoff steps
	Jitter bool
	// Min and Max are the minimum and maximum values of the counter
	Min, Max time.Duration
}

// NewBackoffRetrier creates a retrier that allows for configurable
// min/max times, jitter and maximum retries.
func NewBackoffRetrier(params BackoffRetrierParams) Retrier {
	return &backoffRetrier{
		backoff: backoff.Backoff{
			Factor: params.Factor,
			Jitter: params.Jitter,
			Min:    params.Min,
			Max:    params.Max,
		},
		MaxRetries: params.MaxRetries,
	}
}

// Copies the backoff retrier since it's stateful.
func (rp *backoffRetrier) Copy() Retrier {
	copy := *rp
	return &copy
}

// Exhausted is true if all the retries have been used.
func (rp *backoffRetrier) Exhausted() bool {
	return rp.tries > rp.MaxRetries
}

// CurrentTry is the current try number (0 for the first run before retries)
func (rp *backoffRetrier) CurrentTry() int {
	return rp.tries
}

// Try marks an attempt to call (first call to Try() does not sleep).
// Will return false if the `ctx` is cancelled or if we exhaust our retries.
//
//    rp := RetryPolicy{Backoff:defaultBackoffPolicy, MaxRetries:5}
//
//    for rp.Try(ctx) {
//       <your code>
//    }
//
//    if rp.Cancelled() || rp.Exhausted() {
//       // no more retries needed
//    }
//
func (rp *backoffRetrier) Try(ctx context.Context) bool {
	defer func() { rp.tries++ }()

	select {
	case <-ctx.Done():
		return false
	default:
	}

	if rp.tries == 0 {
		// first 'try' is always free
		return true
	}

	if rp.Exhausted() {
		return false
	}

	select {
	case <-time.After(rp.backoff.Duration()):
		return true
	case <-ctx.Done():
		return false
	}
}
