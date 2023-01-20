// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package emulation_test

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/emulation"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	"github.com/stretchr/testify/require"
)

func TestMockQueue(t *testing.T) {
	traffic := make(chan emulation.Operation, 1000)
	mq := emulation.NewQueue("entity", traffic)
	defer mq.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// no messages exist yet
	msg, err := mq.Receive(ctx, "link")
	require.Nil(t, msg)
	require.ErrorIs(t, err, context.DeadlineExceeded)

	err = mq.Send(context.Background(), &amqp.Message{
		Value: []byte("first message"),
	}, "link")
	require.NoError(t, err)

	// messages exist, but no credits have been added yet.
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg, err = mq.Receive(ctx, "link")
	require.Nil(t, msg)
	require.ErrorIs(t, err, context.DeadlineExceeded)

	// now issue credit
	err = mq.IssueCredit(1, "link")
	require.NoError(t, err)

	msg, err = mq.Receive(context.Background(), "link")
	require.NoError(t, err)
	require.Equal(t, []byte("first message"), msg.Value)

	// and the one message has been consumed.
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg, err = mq.Receive(ctx, "link")
	require.Nil(t, msg)
	require.ErrorIs(t, err, context.DeadlineExceeded)

	// reissue credit
	err = mq.IssueCredit(1, "link")
	require.NoError(t, err)

	err = mq.Send(context.Background(), &amqp.Message{Value: []byte("second message")}, "link")
	require.NoError(t, err)

	msg, err = mq.Receive(context.Background(), "link")
	require.NoError(t, err)
	require.Equal(t, []byte("second message"), msg.Value)
}
