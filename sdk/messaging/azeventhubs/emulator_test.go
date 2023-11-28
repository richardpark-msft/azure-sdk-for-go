// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azeventhubs_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGetRuntimeProperties(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cs := os.Getenv("CS")

	hubName := fmt.Sprintf("EH-Track2-%d", time.Now().UTC().UnixNano())

	if err := createHub("127.0.0.1", hubName); err != nil {
		panic(err)
	}

	cc, err := azeventhubs.NewConsumerClientFromConnectionString(cs, hubName, azeventhubs.DefaultConsumerGroup, &azeventhubs.ConsumerClientOptions{
		RetryOptions: azeventhubs.RetryOptions{
			MaxRetries: -1,
		},
	})

	if err != nil {
		panic(err)
	}

	eventHubProps, err := cc.GetEventHubProperties(context.Background(), nil)
	require.NoError(t, err)

	t.Logf("Event Hub properties: %#v", eventHubProps)

	for _, partID := range eventHubProps.PartitionIDs {
		pp, err := cc.GetPartitionProperties(context.Background(), partID, nil)
		require.NoError(t, err)

		t.Logf("Partition properties[%s]: %#v", partID, pp)
	}
}

func TestEntireEmulator(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cs := os.Getenv("CS")

	hubName := fmt.Sprintf("EH-Track2-%d", time.Now().UTC().UnixNano())

	if err := createHub("127.0.0.1", hubName); err != nil {
		panic(err)
	}

	cc, err := azeventhubs.NewConsumerClientFromConnectionString(cs, hubName, azeventhubs.DefaultConsumerGroup, &azeventhubs.ConsumerClientOptions{
		RetryOptions: azeventhubs.RetryOptions{
			MaxRetries: -1,
		},
	})

	if err != nil {
		panic(err)
	}

	defer func() {
		err = cc.Close(context.Background())

		if err != nil {
			panic(err)
		}
	}()

	printHeader("Using hubname " + hubName)

	producerClient, err := azeventhubs.NewProducerClientFromConnectionString(cs, hubName, &azeventhubs.ProducerClientOptions{
		RetryOptions: azeventhubs.RetryOptions{
			MaxRetries: -1,
		},
	})

	if err != nil {
		panic(err)
	}

	defer producerClient.Close(context.Background())

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)

		printHeader(fmt.Sprintf("Sending round %d", i))

		batch, err := producerClient.NewEventDataBatch(context.Background(), &azeventhubs.EventDataBatchOptions{
			PartitionID: to.Ptr("0"),
		})

		if err != nil {
			panic(err)
		}

		evt := &azeventhubs.EventData{
			Body: []byte(fmt.Sprintf("[%d] hello world", i)),
			Properties: map[string]any{
				"SentAt": time.Now().UnixNano(),
			},
		}
		err = batch.AddEventData(evt, nil)

		if err != nil {
			panic(err)
		}

		// hubProps, err := producerClient.GetEventHubProperties(context.Background(), nil)

		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("Hub properties: %#v\n", hubProps)

		// unsupported as of 2023-11-20
		// t.Logf("=====> Getting partition properties before send")
		// beforePartProps, err := producerClient.GetPartitionProperties(context.Background(), "0", nil)

		// if err != nil {
		// 	panic(err)
		// }

		printHeader("Sending batch")
		if err := producerClient.SendEventDataBatch(context.Background(), batch, nil); err != nil {
			panic(err)
		}

		// unsupported as of 2023-11-20
		// t.Logf("=====> Getting partition properties _after_ send")
		// afterPartProps, err := producerClient.GetPartitionProperties(context.Background(), "0", nil)

		// if err != nil {
		// 	panic(err)
		// }

		// log.Printf("Before:\n%#v\nAfter:\n%#v\n", beforePartProps, afterPartProps)

		printHeader("Batch sent!")
	}

	err = producerClient.Close(context.Background())

	if err != nil {
		panic(err)
	}

	if err := receiveAndPrintEvents(context.Background(), cc, hubName); err != nil {
		panic(err)
	}
}

func receiveAndPrintEvents(ctx context.Context, cc *azeventhubs.ConsumerClient, hubName string) error {
	pc, err := cc.NewPartitionClient("0", &azeventhubs.PartitionClientOptions{
		StartPosition: azeventhubs.StartPosition{
			Earliest: to.Ptr(true),
			// SequenceNumber: to.Ptr[int64](4),
		},
	})

	if err != nil {
		return err
	}

	printHeader("Receiving...")

	const waitTime = time.Minute

	for {
		ctx, cancel := context.WithTimeout(ctx, waitTime)
		events, err := pc.ReceiveEvents(ctx, 1, nil)
		cancel()

		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("======> No messages received for %d seconds\n", waitTime/time.Second)
			break
		}

		if err != nil {
			return err
		}

		for _, evt := range events {
			nanos := evt.Properties["SentAt"].(int64)
			dur := time.Since(time.Unix(0, nanos))
			log.Printf("\n\n=====> seq: %d: offset: %d, body: %s, took: %dms <======\n\n", evt.SequenceNumber, evt.Offset, string(evt.Body), dur/time.Millisecond)
		}
	}

	return nil
}

func printHeader(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	log.Printf("================================== %s ======================\n", msg)
}

type ConsumerGroupInfo struct {
	Name string
}

type EventHubInfo struct {
	Name           string
	PartitionCount int
	ConsumerGroups []ConsumerGroupInfo
}

type NamespaceConfig struct {
	Type     string
	Name     string
	Entities []EventHubInfo
}

type EmulatorEventhubConfig struct {
	NamespaceConfig []NamespaceConfig
}

func createHub(host string, hubName string) error {
	ehConfig := EmulatorEventhubConfig{
		NamespaceConfig: []NamespaceConfig{
			{
				Name: "ignoredbybackend",
				Entities: []EventHubInfo{
					{
						Name:           hubName,
						PartitionCount: 3,
						ConsumerGroups: []ConsumerGroupInfo{
							{Name: azeventhubs.DefaultConsumerGroup},
						},
					},
				},
			},
		},
	}

	jsonText, err := json.Marshal(ehConfig)

	if err != nil {
		return err
	}

	log.Printf("JSON: %s", string(jsonText))

	resp, err := http.Post(fmt.Sprintf("http://%s:8090/Eventhub/Emulator/Configure", host), "application/json", bytes.NewBuffer(jsonText))

	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		text, err := httputil.DumpResponse(resp, true)

		if err != nil {
			return err
		}

		return fmt.Errorf("Failed to create hub. Response:\n%s", text)
	}

	return nil
}
