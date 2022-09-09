// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs

import (
	"context"
	"errors"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/test"
	"github.com/stretchr/testify/require"
)

// NOTE: these tests are live tests, but do use internal/private
// fields for testing.

func TestProcessor_TakingOwnershipAndLinksAreAllActive(t *testing.T) {
	testParams := test.GetConnectionParamsForTest(t)

	cps := newCheckpointStoreForTest()

	processorA := createProcessorForTest(t, testParams, cps, ProcessorStrategyGreedy)
	// processorB := createProcessorForTest(t, testParams, containerName, ProcessorStrategyGreedy)

	defer closeClient(t, processorA.ConsumerClient)

	clientsSyncMap := sync.Map{}

	// since we're manually dispatching we need to initialize the next clients channel
	// ourselves.
	ehProps, err := processorA.Processor.initNextClientsCh(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, ehProps.PartitionIDs)

	err = processorA.Processor.dispatch(context.Background(), ehProps.PartitionIDs, &clientsSyncMap)
	require.NoError(t, err)

	tempMap := syncMapToMap(&clientsSyncMap)
	require.Equal(t, len(ehProps.PartitionIDs), len(tempMap))

	for partitionID, partitionClient := range tempMap {
		internalLinks := partitionClient.innerClient.links.(*internal.Links[amqpwrap.AMQPReceiverCloser])
		exists, err := internalLinks.Exists(partitionID)
		require.NoError(t, err)
		require.Truef(t, exists, "Link should be active for partition %s", partitionID)
	}

	ownerships, err := cps.ListOwnership(context.Background(), testParams.EventHubNamespace, testParams.EventHubName, DefaultConsumerGroup, nil)
	require.NoError(t, err)
	require.Equal(t, len(ehProps.PartitionIDs), len(ownerships))
}

func TestProcessor_TakingOwnershipStealing(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Logf("Running iteration %d", i)
		takingOwnershipStealingTest(t)
		t.Logf("Done running iteration %d", i)
	}
}

func takingOwnershipStealingTest(t *testing.T) {
	testParams := test.GetConnectionParamsForTest(t)

	cps := newCheckpointStoreForTest()

	processorA := createProcessorForTest(t, testParams, cps, ProcessorStrategyGreedy)
	processorB := createProcessorForTest(t, testParams, cps, ProcessorStrategyGreedy)

	defer closeClient(t, processorA.ConsumerClient)

	// since we're manually dispatching we need to initialize the next clients channel
	// ourselves.
	processorAClients := sync.Map{}
	ehProps, err := processorA.Processor.initNextClientsCh(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, ehProps.PartitionIDs)

	err = processorA.Processor.dispatch(context.Background(), ehProps.PartitionIDs, &processorAClients)
	require.NoError(t, err)
	require.Equal(t, len(ehProps.PartitionIDs), len(processorA.Processor.nextClients))

	ownerships, err := cps.ListOwnership(context.Background(), testParams.EventHubNamespace, testParams.EventHubName, DefaultConsumerGroup, nil)
	require.NoError(t, err)
	require.Equal(t, len(ehProps.PartitionIDs), len(ownerships))

	// processorA owns all the partitions. Now we will steal some using ProcessorB
	processorBClients := sync.Map{}

	_, err = processorB.Processor.initNextClientsCh(context.Background())
	require.NoError(t, err)

	err = processorB.Processor.dispatch(context.Background(), ehProps.PartitionIDs, &processorBClients)
	require.NoError(t, err)

	// now we've dispatched from two different processors.
	// ProcessorA owned _all_ partitions
	// ProcessorB came up, split the partitions into two and _must_ steal since
	// all partitions were owned.

	max := len(ehProps.PartitionIDs) / 2

	if len(ehProps.PartitionIDs)%2 != 0 {
		max++
	}

	livenessStateA := checkLiveness(t, &processorAClients, ehProps.PartitionIDs)

	require.NotZero(t, len(livenessStateA.LivePartitions))
	require.NotZero(t, len(livenessStateA.DetachedPartitions))
	require.LessOrEqual(t, len(livenessStateA.LivePartitions), max)
	require.LessOrEqual(t, len(livenessStateA.DetachedPartitions), max)

	// give a little sleep as the AMQP receivers need to pump any messages from
	// the broker to indicate they're detached
	// time.Sleep(5 * time.Second)

	livenessStateB := checkLiveness(t, &processorBClients, ehProps.PartitionIDs)

	require.Zero(t, len(livenessStateB.DetachedPartitions))
	require.NotZero(t, len(livenessStateB.LivePartitions))
	require.LessOrEqual(t, len(livenessStateB.LivePartitions), max)

	t.Logf("A: %#v\nB: %#v", livenessStateA, livenessStateB)

	// partitions will have swapped owners.
	require.Equal(t, getKeys(livenessStateA.DetachedPartitions), getKeys(livenessStateB.LivePartitions))

	// and all partitions are alive.
	require.Equal(t, len(ehProps.PartitionIDs), len(livenessStateA.LivePartitions)+len(livenessStateB.LivePartitions))

	// and all the new clients will come back from ProcessorB
	require.Equal(t, len(livenessStateB.LivePartitions), len(processorB.Processor.nextClients))
}

func getKeys(m map[string]bool) []string {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

type livenessTestData struct {
	DetachedPartitions map[string]bool
	LivePartitions     map[string]bool
}

func checkLiveness(t *testing.T, clients *sync.Map, partitionIDs []string) livenessTestData {
	m := syncMapToMap(clients)

	detachedPartitions := make(chan string, len(partitionIDs))
	livePartitions := make(chan string, len(partitionIDs))

	wg := sync.WaitGroup{}

	for partitionID, client := range m {
		wg.Add(1)

		go func(partitionID string, client *ProcessorPartitionClient) {
			defer wg.Done()

			internalLinks := client.innerClient.links.(*internal.Links[amqpwrap.AMQPReceiverCloser])
			amqpLink, err := internalLinks.GetLink(context.Background(), partitionID)
			require.NoError(t, err)

			// check if it's actually detached
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			evt, err := amqpLink.Link.Receive(ctx)
			require.Nil(t, evt)

			newError := internal.TransformError(err)

			// we expect an error - DeadLineExceeded if the link is live
			// or OwnershipLost if it was detached (expected)
			require.Error(t, err)

			var eventHubsErr *Error

			if errors.As(newError, &eventHubsErr) {
				detachedPartitions <- partitionID
				require.Equal(t, eventHubsErr.Code, CodeOwnershipLost)
			} else if errors.Is(newError, context.DeadlineExceeded) {
				// this link is still alive
				livePartitions <- partitionID
			} else {
				require.NoError(t, err, "Unexpected error - wasn't a live error or an ownership lost error")
			}
		}(partitionID, client)
	}

	wg.Wait()

	close(detachedPartitions)
	close(livePartitions)

	return struct {
		DetachedPartitions map[string]bool
		LivePartitions     map[string]bool
	}{
		DetachedPartitions: channelToMap(t, detachedPartitions),
		LivePartitions:     channelToMap(t, livePartitions),
	}
}

type processorLiveTestData struct {
	ConsumerClient *ConsumerClient
	Processor      *Processor
}

func channelToMap[T interface{ comparable }](t *testing.T, ch <-chan T) map[T]bool {
	values := map[T]bool{}

	for v := range ch {
		require.False(t, values[v])
		values[v] = true
	}

	return values
}

func syncMapToMap(clients *sync.Map) map[string]*ProcessorPartitionClient {
	m := map[string]*ProcessorPartitionClient{}

	clients.Range(func(key, value any) bool {
		m[key.(string)] = value.(*ProcessorPartitionClient)
		return true
	})

	return m
}

func createProcessorForTest(t *testing.T, testParams test.ConnectionParamsForTest, cps CheckpointStore, strategy ProcessorStrategy) processorLiveTestData {
	// Create the checkpoint store
	// NOTE: the container must exist before the checkpoint store can be used.
	consumerClient, err := NewConsumerClientFromConnectionString(testParams.ConnectionString, testParams.EventHubName, DefaultConsumerGroup, nil)
	require.NoError(t, err)

	processor, err := NewProcessor(consumerClient, cps, &NewProcessorOptions{
		LoadBalancingStrategy: strategy,
	})
	require.NoError(t, err)

	return processorLiveTestData{
		ConsumerClient: consumerClient,
		Processor:      processor,
	}
}

func closeClient(t *testing.T, client interface {
	Close(ctx context.Context) error
}) {
	err := client.Close(context.Background())
	require.NoError(t, err)
}
