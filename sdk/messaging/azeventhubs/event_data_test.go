// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs

import (
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/go-amqp"
	"github.com/stretchr/testify/require"
)

func TestEventData_Annotations(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		re := newReceivedEventData(&amqp.Message{})

		require.Empty(t, re.Body)
		require.Empty(t, re.SystemProperties)
		require.Nil(t, re.EnqueuedTime)
		require.Equal(t, int64(0), re.SequenceNumber)
		require.Nil(t, re.Offset)
		require.Nil(t, re.PartitionKey)
	})

	t.Run("invalid types", func(t *testing.T) {
		// invalid types for properties doesn't crash us
		re := newReceivedEventData(&amqp.Message{
			Annotations: amqp.Annotations{
				"x-opt-partition-key":   99,
				"x-opt-sequence-number": "101",
				"x-opt-offset":          102,
				"x-opt-enqueued-time":   "now",
			},
		})

		require.Empty(t, re.Body)
		require.Empty(t, re.SystemProperties)
		require.Nil(t, re.EnqueuedTime)
		require.Equal(t, int64(0), re.SequenceNumber)
		require.Nil(t, re.Offset)
		require.Nil(t, re.PartitionKey)
	})
}

func TestEventData_newReceivedEventData(t *testing.T) {
	now := time.Now().UTC()

	re := newReceivedEventData(&amqp.Message{
		Properties: &amqp.MessageProperties{
			ContentType: to.Ptr("content type"),
			MessageID:   "message id",
		},
		Data: [][]byte{[]byte("hello world")},
		Annotations: map[any]any{
			"hello":                 "world",
			5:                       "ignored",
			"x-opt-partition-key":   "partition key",
			"x-opt-sequence-number": int64(101),
			"x-opt-offset":          "102",
			"x-opt-enqueued-time":   now,
		},
		ApplicationProperties: map[string]any{
			"application property 1": "application prioperty value 1",
		},
	})

	expectedRawAnnotations := map[any]any{
		"hello":                 "world",
		5:                       "ignored",
		"x-opt-partition-key":   "partition key",
		"x-opt-sequence-number": int64(101),
		"x-opt-offset":          "102",
		"x-opt-enqueued-time":   now,
	}

	expectedBody := [][]byte{
		[]byte("hello world"),
	}

	expectedAppProperties := map[string]any{
		"application property 1": "application prioperty value 1",
	}

	require.Equal(t, &ReceivedEventData{
		EventData: EventData{
			Body:        expectedBody[0],
			ContentType: to.Ptr("content type"),
			MessageID:   to.Ptr("message id"),
			Properties:  expectedAppProperties,
		},
		SystemProperties: map[string]any{
			"hello": "world",
		},
		EnqueuedTime:   &now,
		SequenceNumber: 101,
		Offset:         to.Ptr[int64](102),
		PartitionKey:   to.Ptr("partition key"),

		// this is the azeventhubs version of the "raw" AMQP message.
		// It's basically a 1:1 copy of goamqp's message.
		RawAMQPMessage: &AMQPAnnotatedMessage{
			ApplicationProperties: expectedAppProperties,
			Body: AMQPAnnotatedMessageBody{
				Data: expectedBody,
			},
			MessageAnnotations: expectedRawAnnotations,
			Properties: &AMQPAnnotatedMessageProperties{
				ContentType: to.Ptr("content type"),
				MessageID:   "message id",
			},
		},
	}, re)
}
