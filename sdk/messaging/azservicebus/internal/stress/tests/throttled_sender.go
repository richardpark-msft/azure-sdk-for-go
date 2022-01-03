// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tests

import (
	"errors"
	"fmt"
	"log"
	"strings"	1`1`
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/stress/shared"
)

const numToSend = 500

func generateBigBatch(sc *shared.StressContext, sender *azservicebus.Sender) *azservicebus.MessageBatch {
	body := [1024]byte{0}

	batch, err := sender.NewMessageBatch(sc.Context, nil)
	sc.PanicOnError("failed to create message batch", err)

	// pack the batch as full as we can (we can ignore excess)
	for i := 0; i < 1000; i++ {
		err := batch.AddMessage(&azservicebus.Message{
			Body: body[:],
			ApplicationProperties: map[string]interface{}{
				"i": i,
			},
		})

		if errors.Is(err, azservicebus.ErrMessageTooLarge) {
			return batch
		}

		sc.PanicOnError("failed to add message to batch", err)
	}

	return batch
}

func ThrottledSender(remainingArgs []string) {
	start := time.Now()

	defer func() {
		log.Printf("Time to send %d message batches: %d seconds", numToSend, time.Since(start)/time.Second)
	}()

	// send enough that throttling should happen. Check any reported telemetry to see if
	// we properly did our retries.
	sc := shared.MustCreateStressContext("ThrottledSender")

	sc.TrackEvent("Start")
	defer sc.TrackEvent("End")

	queueName := strings.ToLower(fmt.Sprintf("throttled-sender-%X", time.Now().UnixNano()))

	log.Printf("Creating queue")
	shared.MustCreateAutoDeletingQueue(sc, queueName)

	client, err := azservicebus.NewClientFromConnectionString(sc.ConnectionString, nil)
	sc.PanicOnError("failed to create client", err)

	defer client.Close(sc.Context)

	sender, err := client.NewSender(queueName, nil)
	sc.PanicOnError("failed to create sender", err)

	wg := sync.WaitGroup{}

	log.Printf("Sending %d message batches", numToSend)

	batch := generateBigBatch(sc, sender)

	sentStats := sc.NewStat("sent")

	for i := 0; i < numToSend; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for {
				err := sender.SendMessageBatch(sc.Context, batch)

				if err != nil {
					var nonRetriable interface {
						NonRetriable()
						error
					}

					if errors.As(err, &nonRetriable) {
						panic(fmt.Errorf("non retriable error returned: %w", err))
					}

					log.Printf("Retriable error returned: %#v", err)
					continue
				}

				sentStats.AddSent(batch.NumMessages())
				break
			}
		}(i)
	}

	wg.Wait()

	sc.PanicOnError("Closing client", client.Close(sc.Context))

	log.Printf("Sent %d message batches", numToSend)
}
