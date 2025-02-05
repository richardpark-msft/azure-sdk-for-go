package azeventhubs_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/test"
	"github.com/stretchr/testify/require"
)

func TestConsumerClient_GeoDR(t *testing.T) {
	testParams := test.GetConnectionParamsForTest(t)

	if testParams.GeoDRNamespace == "" || testParams.GeoDRHubName == "" {
		t.Skipf("Skipping GeoDR test, EVENTHUBS_GEODR_NAMESPACE or EVENTHUBS_GEODR_HUBNAME was not set")
	}

	pc, err := azeventhubs.NewProducerClient(testParams.GeoDRNamespace, testParams.GeoDRHubName, testParams.Cred, nil)
	require.NoError(t, err)

	props, err := pc.GetEventHubProperties(context.Background(), nil)
	require.NoError(t, err)
	require.True(t, props.GeoReplication)

	partProps, err := pc.GetPartitionProperties(context.Background(), "0", nil)
	require.NoError(t, err)

	t.Logf("LastEnqueuedOffset: %#v, LastEnqueuedSequenceNumber: %#v", partProps.LastEnqueuedOffset, partProps.LastEnqueuedSequenceNumber)

	{
		// we send a couple of events so the processor tests, that can't be started inclusive, will still have something
		// predictable to retrieve.
		batch, err := pc.NewEventDataBatch(context.Background(), &azeventhubs.EventDataBatchOptions{
			PartitionID: to.Ptr("0"),
		})
		require.NoError(t, err)

		err = batch.AddEventData(&azeventhubs.EventData{Body: []byte("body")}, nil)
		require.NoError(t, err)

		err = batch.AddEventData(&azeventhubs.EventData{Body: []byte("body2")}, nil)
		require.NoError(t, err)

		err = pc.SendEventDataBatch(context.Background(), batch, nil)
		require.NoError(t, err)
	}

	t.Run("StartAtLegacyOffset", func(t *testing.T) {
		cc, err := azeventhubs.NewConsumerClient(testParams.GeoDRNamespace, testParams.GeoDRHubName, azeventhubs.DefaultConsumerGroup, testParams.Cred, nil)
		require.NoError(t, err)
		defer test.RequireClose(t, cc)

		pc, err := cc.NewPartitionClient("0", &azeventhubs.PartitionClientOptions{
			StartPosition: azeventhubs.StartPosition{
				Offset: to.Ptr("0"),
			},
		})
		require.NoError(t, err)
		defer test.RequireClose(t, pc)

		events, err := pc.ReceiveEvents(context.Background(), 1, nil)
		require.Contains(t, err.Error(), "Condition: com.microsoft:georeplication:invalid-offset")
		require.Empty(t, events)
	})

	t.Run("StartAtSequenceNumber", func(t *testing.T) {
		// NOTE: GetPartitionProperties returns an offset that can't be used for starting.
		cc, err := azeventhubs.NewConsumerClient(testParams.GeoDRNamespace, testParams.GeoDRHubName, azeventhubs.DefaultConsumerGroup, testParams.Cred, nil)
		require.NoError(t, err)
		defer test.RequireClose(t, cc)

		pc, err := cc.NewPartitionClient("0", &azeventhubs.PartitionClientOptions{
			StartPosition: azeventhubs.StartPosition{
				SequenceNumber: &partProps.LastEnqueuedSequenceNumber,
				Inclusive:      true,
			},
		})
		require.NoError(t, err)
		defer test.RequireClose(t, pc)

		events, err := pc.ReceiveEvents(context.Background(), 1, nil)
		require.Nil(t, err)
		require.NotEmpty(t, events)
		require.Equal(t, partProps.LastEnqueuedSequenceNumber, events[0].SequenceNumber)
	})

	// there'll be at least two events in the event hub
	getEarliestEvents := func() []*azeventhubs.ReceivedEventData {
		cc, err := azeventhubs.NewConsumerClient(testParams.GeoDRNamespace, testParams.GeoDRHubName, azeventhubs.DefaultConsumerGroup, testParams.Cred, nil)
		require.NoError(t, err)
		defer test.RequireClose(t, cc)

		pc, err := cc.NewPartitionClient("0", &azeventhubs.PartitionClientOptions{
			StartPosition: azeventhubs.StartPosition{
				Earliest: to.Ptr(true),
			},
		})
		require.NoError(t, err)
		defer test.RequireClose(t, pc)

		// get an event, so we have a stroffset to use.
		origEvents, err := pc.ReceiveEvents(context.Background(), 2, nil)
		require.NoError(t, err)

		return origEvents
	}

	t.Run("StartAtStroffset", func(t *testing.T) {
		events := getEarliestEvents()

		cc, err := azeventhubs.NewConsumerClient(testParams.GeoDRNamespace, testParams.GeoDRHubName, azeventhubs.DefaultConsumerGroup, testParams.Cred, nil)
		require.NoError(t, err)
		defer test.RequireClose(t, cc)

		pcFromOffset, err := cc.NewPartitionClient("0", &azeventhubs.PartitionClientOptions{
			StartPosition: azeventhubs.StartPosition{
				Offset:    &events[0].Offset,
				Inclusive: true,
			},
		})
		require.NoError(t, err)
		defer test.RequireClose(t, pc)

		newEvents, err := pcFromOffset.ReceiveEvents(context.Background(), 1, nil)
		require.NoError(t, err)

		require.Equal(t, newEvents[0].SequenceNumber, events[0].SequenceNumber)
	})

	t.Run("ProcessorRestoreWithStroffset", func(t *testing.T) {
		events := getEarliestEvents()

		td := setupProcessorTest(t, true)

		err = td.CheckpointStore.SetCheckpoint(context.Background(), azeventhubs.Checkpoint{
			ConsumerGroup:           azeventhubs.DefaultConsumerGroup,
			FullyQualifiedNamespace: testParams.GeoDRNamespace,
			EventHubName:            testParams.GeoDRHubName,
			PartitionID:             "0",
			Offset:                  &events[0].Offset, // this is _exclusive_, so the next event we get for this partition should be events[1]
			SequenceNumber:          &events[0].SequenceNumber,
		}, nil)
		require.NoError(t, err)

		proc := td.Create(nil)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			err := proc.Run(ctx)
			require.NoError(t, err)
		}()

		gotEvent := false

		for {
			pc := proc.NextPartitionClient(context.Background())

			if pc.PartitionID() == "0" {
				newEvents, err := pc.ReceiveEvents(context.Background(), 1, nil)
				require.NoError(t, err)
				require.Equal(t, events[1].SequenceNumber, newEvents[0].SequenceNumber)
				cancel()
				gotEvent = true
				break
			}
		}

		require.True(t, gotEvent)
	})

	t.Run("ProcessorRestoreWithLegacyOffset", func(t *testing.T) {
		events := getEarliestEvents()

		td := setupProcessorTest(t, true)

		err = td.CheckpointStore.SetCheckpoint(context.Background(), azeventhubs.Checkpoint{
			ConsumerGroup:           azeventhubs.DefaultConsumerGroup,
			FullyQualifiedNamespace: testParams.GeoDRNamespace,
			EventHubName:            testParams.GeoDRHubName,
			PartitionID:             "0",
			Offset:                  to.Ptr("0"), // this is invalid - you can't use old offsets with a new GeoDR-enabled event hub.
			SequenceNumber:          &events[0].SequenceNumber,
		}, nil)
		require.NoError(t, err)

		proc := td.Create(nil)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			err := proc.Run(ctx)
			require.NoError(t, err)
		}()

		gotEvent := false

		for {
			pc := proc.NextPartitionClient(context.Background())

			if pc.PartitionID() == "0" {

				// TODO: it seems like we have an issue here - we need to "restart" the partition client at earliest, if this happens.
				// it might be that we have to leave this up to the user to handle, and (if so), we can export a predictable error for this
				// this also puts the control into the user's hands as to how they want to handle this specific situation.They'll need to use
				// one of the earliest/latest versions to do it properly.
				//
				// I think, given the way we've designed the library, that it's best to provide the user the information and let them design the strategy
				// or fallback value (perhaps, using StartPosition).
				require.Fail(t, "Not sure how to handle this condition in here yet.")

				newEvents, err := pc.ReceiveEvents(context.Background(), 1, nil)
				require.NoError(t, err)
				require.Equal(t, events[1].SequenceNumber, newEvents[0].SequenceNumber)
				cancel()
				gotEvent = true
				break
			}
		}

		require.True(t, gotEvent)

		require.Fail(t, "asdfasdf")
	})

	t.Run("BUG: StartWithOffsetFromGetPartitionProperties", func(t *testing.T) {
		// NOTE: GetPartitionProperties returns an offset that can't be used as a starting point in GeoDR.
		// My expectation is that I can take the offset, directly from GetPartitionProperties, and start a client
		// at that position.
		cc, err := azeventhubs.NewConsumerClient(testParams.GeoDRNamespace, testParams.GeoDRHubName, azeventhubs.DefaultConsumerGroup, testParams.Cred, nil)
		require.NoError(t, err)
		defer test.RequireClose(t, cc)

		pc, err := cc.NewPartitionClient("0", &azeventhubs.PartitionClientOptions{
			StartPosition: azeventhubs.StartPosition{
				Offset: &partProps.LastEnqueuedOffset,
			},
		})
		require.NoError(t, err)
		defer test.RequireClose(t, pc)

		events, err := pc.ReceiveEvents(context.Background(), 1, nil)
		require.Contains(t, err.Error(), "Condition: com.microsoft:georeplication:invalid-offset")
		require.Empty(t, events)
	})
}
