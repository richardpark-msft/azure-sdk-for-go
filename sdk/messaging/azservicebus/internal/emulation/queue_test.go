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
	mq := emulation.NewQueue()
	defer mq.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// no messages exist yet
	msg, err := mq.Receive(ctx)
	require.Nil(t, msg)
	require.ErrorIs(t, err, context.DeadlineExceeded)

	err = mq.Send(context.Background(), &amqp.Message{
		Value: []byte("first message"),
	})
	require.NoError(t, err)

	// messages exist, but no credits have been added yet.
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg, err = mq.Receive(ctx)
	require.Nil(t, msg)
	require.ErrorIs(t, err, context.DeadlineExceeded)

	// now issue credit
	mq.IssueCredit(1)

	msg, err = mq.Receive(context.Background())
	require.NoError(t, err)
	require.Equal(t, []byte("first message"), msg.Value)

	// and the one message has been consumed.
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg, err = mq.Receive(ctx)
	require.Nil(t, msg)
	require.ErrorIs(t, err, context.DeadlineExceeded)

	// reissue credit
	mq.IssueCredit(1)
	err = mq.Send(context.Background(), &amqp.Message{Value: []byte("second message")})
	require.NoError(t, err)

	msg, err = mq.Receive(context.Background())
	require.NoError(t, err)
	require.Equal(t, []byte("second message"), msg.Value)
}
