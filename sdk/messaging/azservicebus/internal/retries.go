package internal

import (
	"time"

	"github.com/jpillora/backoff"
)

var DefaultBackoffPolicy = struct {
	backoff.Backoff
	Retries int
}{
	Backoff: backoff.Backoff{
		Factor: 1,
		Min:    time.Second * 5,
	},
	Retries: 5,
}
