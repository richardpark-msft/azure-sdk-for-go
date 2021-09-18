// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/require"
)

func TestMessageUnitTest(t *testing.T) {
	t.Run("toAMQPMessage", func(t *testing.T) {
		message := &Message{}

		// basic thing - it's totally fine to send a message nothing in it.
		amqpMessage, err := message.toAMQPMessage()
		require.NoError(t, err)
		require.Empty(t, amqpMessage.Annotations)
		require.NotEmpty(t, amqpMessage.Properties.MessageID, "MessageID is (currently) automatically filled out if you don't specify one")

		message = &Message{
			ID:                      "message id",
			Body:                    []byte("the body"),
			PartitionKey:            to.StringPtr("partition key"),
			TransactionPartitionKey: to.StringPtr("via partition key"),
			SessionID:               to.StringPtr("session id"),
		}

		amqpMessage, err = message.toAMQPMessage()
		require.NoError(t, err)

		require.EqualValues(t, "message id", amqpMessage.Properties.MessageID)
		require.EqualValues(t, "session id", amqpMessage.Properties.GroupID)

		require.EqualValues(t, "the body", string(amqpMessage.Data[0]))
		require.EqualValues(t, 1, len(amqpMessage.Data))

		require.EqualValues(t, map[interface{}]interface{}{
			annotationPartitionKey:    "partition key",
			annotationViaPartitionKey: "via partition key",
		}, amqpMessage.Annotations)
	})
}

func toInt16Ptr(i int16) *int16 {
	return &i
}
