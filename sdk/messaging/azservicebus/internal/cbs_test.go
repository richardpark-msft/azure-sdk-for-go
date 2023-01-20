// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/emulation"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/mock"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestNegotiateClaimWithMock(t *testing.T) {
	test.EnableStdoutLogging()

	md := emulation.NewMockData(t, &emulation.MockDataOptions{
		NewReceiverFn: func(orig *mock.MockAMQPReceiverCloser, ctx context.Context, source string, opts *amqp.ReceiverOptions) (*mock.MockAMQPReceiverCloser, error) {
			orig.EXPECT().Close(gomock.Any()).Times(1)
			orig.EXPECT().AcceptMessage(gomock.Any(), gomock.Any()).AnyTimes()
			return orig, nil
		},
		NewSenderFn: func(orig *mock.MockAMQPSenderCloser, ctx context.Context, target string, opts *amqp.SenderOptions) (*mock.MockAMQPSenderCloser, error) {
			orig.EXPECT().Close(gomock.Any()).Times(1)
			return orig, nil
		},
		NewSessionFn: func(orig *mock.MockAMQPSession, ctx context.Context, opts *amqp.SessionOptions) (*mock.MockAMQPSession, error) {
			orig.EXPECT().Close(gomock.Any()).Times(1)
			return orig, nil
		},
	})

	client, err := md.NewConnection(context.Background())
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = NegotiateClaim(ctx, "audience", client, md.NewTokenProvider())
	require.NoError(t, err)
}
