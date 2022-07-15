// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tests

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	azlog "github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/admin"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/stress/shared"
)

func SendAndReceiveDrain(remainingArgs []string) {
	sc := shared.MustCreateStressContext("SendAndReceiveDrainTest")

	sc.TrackEvent("Start")
	defer sc.End()

	queueName := strings.ToLower(fmt.Sprintf("queue-%X", time.Now().UnixNano()))

	log.Printf("Creating queue")

	// set a long lock duration to make it obvious when a message is being lost in our
	// internal buffer or somewhere along the way.
	// This mimics the scenario mentioned in this issue filed by a customer:
	// https://github.com/Azure/azure-sdk-for-go/issues/17853
	lockDuration := "PT5M"

	shared.MustCreateAutoDeletingQueue(sc, queueName, &admin.QueueProperties{
		LockDuration: &lockDuration,
	})

	client, err := azservicebus.NewClientFromConnectionString(sc.ConnectionString, nil)
	sc.PanicOnError("failed to create client", err)

	sender, err := client.NewSender(queueName, nil)
	sc.PanicOnError("failed to create sender", err)

	receiver, err := client.NewReceiverForQueue(queueName, nil)
	sc.PanicOnError("Failed to create receiver", err)

	stat := sc.NewStat("test")

	for i := 0; i < 1000; i++ {
		log.Printf("=====> Round [%d] <====", i)

		const numToSend = 2000
		const bodyLen = 4096
		shared.MustGenerateMessages(sc, sender, numToSend, bodyLen, nil)

		var totalCompleted int64

		for totalCompleted < numToSend {
			log.Printf("Receiving messages [%d/%d]...", totalCompleted, numToSend)
			ctx, cancel := context.WithTimeout(sc.Context, time.Minute)
			defer cancel()

			messages, err := receiver.ReceiveMessages(ctx, numToSend+100, nil)

			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				// this is bad - it means we didn't get _any_ messages within an entire
				// minute and might indicate that we're hitting the customer bug.

				log.Printf("Exceeded the timeout, trying one more time real fast")

				// let's see if there is some other momentary issue happening here by doing a quick receive again.
				ctx, cancel := context.WithTimeout(sc.Context, time.Minute)
				defer cancel()
				messages, err = receiver.ReceiveMessages(ctx, numToSend+100, nil)
				sc.PanicOnError("Exceeded a minute while waiting for messages", err)
			}

			stat.AddReceived(int32(len(messages)))
			sc.PrintStats()

			azlog.Writef("stress", "Got %d messages, completing...", len(messages))

			wg := sync.WaitGroup{}

			for _, m := range messages {
				wg.Add(1)

				func(m *azservicebus.ReceivedMessage) {
					if len(m.Body) != bodyLen {
						sc.PanicOnError("Body length issue", fmt.Errorf("Invalid body length - expected %d, got %d", bodyLen, len(m.Body)))
					}

					if err := receiver.CompleteMessage(sc.Context, m, nil); err != nil {
						sc.PanicOnError("Failed to complete message", err)
					}

					stat.AddCompleted(1)
				}(m)
			}

			wg.Wait()

			azlog.Writef("stress", "Done completing messages")
			sc.PrintStats()
		}

		log.Printf("[end] Receiving messages (all received)")
	}
}
