// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/stretchr/testify/require"
)

// newEventDataBatchFromSlice creates batches from 'events'. If all of these events cannot fit into a single batch it asserts.
func newEventDataBatchFromSlice(t *testing.T, producer *azeventhubs.ProducerClient, events []*azeventhubs.EventData, options *azeventhubs.NewEventDataBatchOptions) *azeventhubs.EventDataBatch {
	batch, err := producer.NewEventDataBatch(context.Background(), options)
	require.NoError(t, err)

	for i := 0; i < len(events); i++ {
		err := batch.AddEventData(events[i], nil)
		require.NoError(t, err)
	}

	return batch
}

// sendEvents sends all the events, batching them as efficiently as possible.
// It returns the partition properties before the events are sent.
func sendEvents(t *testing.T, producer *azeventhubs.ProducerClient, events []*azeventhubs.EventData, partitionID string) azeventhubs.StartPosition {
	batch := newEventDataBatchFromSlice(t, producer, events, &azeventhubs.NewEventDataBatchOptions{
		PartitionID: &partitionID,
	})

	beforeSend, err := producer.GetPartitionProperties(context.Background(), partitionID, nil)
	require.NoError(t, err)

	err = producer.SendEventBatch(context.Background(), batch, nil)
	require.NoError(t, err)

	return getStartPosition(beforeSend)
}

// closeClient closes a client and asserts on the returned error. Removes some
// clutter from the test code when defer'ing the Close().
func closeClient(t *testing.T, client interface {
	Close(ctx context.Context) error
}) {
	err := client.Close(context.Background())
	require.NoError(t, err)
}
