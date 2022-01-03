// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"math"
	"math/rand"
	"time"
)

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
	Try(ctx context.Context) bool
}

// Encapsulates a backoff policy, which allows you to configure the amount of
// time in between retries as well as the maximum retries allowed (via MaxRetries)
// NOTE: this should be copied by the caller as it is stateful.
type backoffRetrier struct {
	RetryOptions
	tries int
}

// NewBackoffRetrier creates a retrier that allows for configurable
// min/max times, jitter and maximum retries.
func NewBackoffRetrier(options RetryOptions) Retrier {
	return &backoffRetrier{
		RetryOptions: options,
	}
}

// Copies the backoff retrier since it's stateful.
func (rp *backoffRetrier) Copy() Retrier {
	copy := *rp
	return &copy
}

// CurrentTry is the current try number (0 for the first run before retries)
func (rp *backoffRetrier) CurrentTry() int {
	return rp.tries
}

// Try marks an attempt to call (first call to Try() does not sleep).
// Will return false if the `ctx` is cancelled or if we exhaust our retries.
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

// RetryOptions configures the retry policy's behavior.
// In general, if an option or argument allows a RetryOptions argument it can be nil, and will default
// to the values specified for each field below.
type RetryOptions struct {
	// MaxRetries specifies the maximum number of attempts a failed operation will be retried
	// before producing an error.
	// The default value is three.  A value less than zero means one try and no retries.
	MaxRetries int32

	// TryTimeout indicates the maximum time allowed for any single try of an HTTP request.
	// This is disabled by default.  Specify a value greater than zero to enable.
	// NOTE: Setting this to a small value might cause premature HTTP request time-outs.
	TryTimeout time.Duration

	// RetryDelay specifies the initial amount of delay to use before retrying an operation.
	// The delay increases exponentially with each retry up to the maximum specified by MaxRetryDelay.
	// The default value is four seconds.  A value less than zero means no delay between retries.
	RetryDelay time.Duration

	// MaxRetryDelay specifies the maximum delay allowed before retrying an operation.
	// Typically the value is greater than or equal to the value specified in RetryDelay.
	// The default Value is 120 seconds.  A value less than zero means there is no cap.
	MaxRetryDelay time.Duration
}

func setDefaults(o *RetryOptions) {
	if o.MaxRetries == 0 {
		o.MaxRetries = 3
	} else if o.MaxRetries < 0 {
		o.MaxRetries = 0
	}

	if o.MaxRetryDelay == 0 {
		o.MaxRetryDelay = 120 * time.Second
	} else if o.MaxRetryDelay < 0 {
		// not really an unlimited cap, but sufficiently large enough to be considered as such
		o.MaxRetryDelay = math.MaxInt64
	}

	if o.RetryDelay == 0 {
		o.RetryDelay = 4 * time.Second
	} else if o.RetryDelay < 0 {
		o.RetryDelay = 0
	}
}

func calcDelay(o RetryOptions, try int32) time.Duration { // try is >=1; never 0
	pow := func(number int64, exponent int32) int64 { // pow is nested helper function
		var result int64 = 1
		for n := int32(0); n < exponent; n++ {
			result *= number
		}
		return result
	}

	delay := time.Duration(pow(2, try)-1) * o.RetryDelay

	// Introduce some jitter:  [0.0, 1.0) / 2 = [0.0, 0.5) + 0.8 = [0.8, 1.3)
	delay = time.Duration(delay.Seconds() * (rand.Float64()/2 + 0.8) * float64(time.Second)) // NOTE: We want math/rand; not crypto/rand
	if delay > o.MaxRetryDelay {
		delay = o.MaxRetryDelay
	}
	return delay
}
