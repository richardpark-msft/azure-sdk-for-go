// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/sberrors"
	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

func Test_RenewLock_Failures(t *testing.T) {
	mgmt, message, cleanup := renewLockTestSetup(t)
	defer cleanup()

	//  let that message expire!
	time.Sleep(7 * time.Second)

	/*
		unhandled error link 1e7c4e10-1e57-4a37-8d5d-7a6408816ed0: status code 410 and description: The lock supplied is invalid. Either the lock expired, or the message has already been removed from the queue. Reference:34da35a1-3c28-43d3-a6c2-bbe370d13b6a, TrackingId:b6c1d77c-4dac-406a-a22a-afebcf4d17e5_B29, SystemTracker:riparkdev2:Queue:queue-16ac1f7825a0b2b4, Timestamp:2021-10-08T17:49:36
	*/
	newExpirationTimes, err := mgmt.RenewLocks(context.Background(), message.rawAMQPMessage.LinkName(), []amqp.UUID{message.LockToken})
	sbe := sberrors.AsServiceBusError(context.Background(), err)
	require.EqualValues(t, sberrors.FixNotPossible, sbe.Fix)
	require.Nil(t, newExpirationTimes)

	bogusUUID := amqp.UUID([16]byte{})
	newExpirationTimes, err = mgmt.RenewLocks(context.Background(), message.rawAMQPMessage.LinkName(), []amqp.UUID{bogusUUID})
	sbe = sberrors.AsServiceBusError(context.Background(), err)
	require.EqualValues(t, sberrors.FixNotPossible, sbe.Fix)
	require.Nil(t, newExpirationTimes)
}

func Test_RenewLock_Succeed(t *testing.T) {
	mgmt, message, cleanup := renewLockTestSetup(t)
	defer cleanup()

	newExpirationTimes, err := mgmt.RenewLocks(context.Background(), message.rawAMQPMessage.LinkName(), []amqp.UUID{message.LockToken})
	require.NoError(t, err)
	require.NotEmpty(t, newExpirationTimes)
	require.Greater(t, time.Now().UTC().Sub(newExpirationTimes[0]), 3*time.Second)
}

func renewLockTestSetup(t *testing.T) (internal.MgmtClient, *ReceivedMessage, func()) {
	client, cleanup, queueName := setupLiveTest(t, &internal.QueueDescription{
		LockDuration: to.StringPtr(internal.DurationTo8601Seconds(5 * time.Second)),
	})

	ctx := context.Background()

	receiver, err := client.NewReceiverForQueue(queueName, nil)
	require.NoError(t, err)

	sender, err := client.NewSender(queueName)
	require.NoError(t, err)

	err = sender.SendMessage(ctx, &Message{
		Body: []byte("renewal test"),
	})
	require.NoError(t, err)

	message, err := receiver.ReceiveMessage(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, message)

	_, _, mgmt, _, err := receiver.amqpLinks.Get(ctx)
	require.NoError(t, err)

	return mgmt, message, cleanup
}
