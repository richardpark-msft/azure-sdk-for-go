// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/sberrors"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

// autoRenewMessageLock renews the lock for a message.
func autoRenewMessageLock(ctx context.Context, retrier internal.Retrier, links internal.AMQPLinks, message *ReceivedMessage) {
	renewWithRetries := func(ctx context.Context) time.Time {
		for retrier.Try(ctx) {
			_, _, mgmt, _, err := links.Get(ctx)

			if err == nil {
				var newExpirationTimes []time.Time

				now := time.Now().UTC()
				newExpirationTimes, err = mgmt.RenewLocks(ctx, message.rawAMQPMessage.LinkName(), []amqp.UUID{message.LockToken})

				if err == nil {
					log.Printf("===> Got back %s, %s", newExpirationTimes[0].String(), now)
					return newExpirationTimes[0]
				}
			}

			sbe := sberrors.AsServiceBusError(ctx, err)

			if sbe.Fix == sberrors.FixNotPossible {
				break
			}
		}

		return time.Time{}
	}

	for {
		ok := func() bool {
			ctx, span := tab.StartSpan(ctx, internal.SpanRenewMessage)
			defer span.End()

			nextExpirationTime := renewWithRetries(ctx)

			if nextExpirationTime.IsZero() {
				span.AddAttributes(tab.StringAttribute("StopReason", "No next expiration time"))
				return false
			}

			waitTime := time.Now().UTC().Sub(nextExpirationTime)

			if waitTime <= 0 {
				span.AddAttributes(tab.StringAttribute("StopReason", "No time remaining (expired)"))
				return false
			}

			span.AddAttributes(tab.Int64Attribute("NextInSeconds", int64(waitTime/2/time.Second)))

			select {
			case <-ctx.Done():
				span.AddAttributes(tab.StringAttribute("StopReason", "Context cancelled)"))
				return false
			case <-time.After(waitTime / 2):
				return true
			}
		}()

		if !ok {
			break
		}
	}
}
