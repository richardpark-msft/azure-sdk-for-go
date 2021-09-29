// Copyright (c) Microsoft Corporation. All rights reserved.

// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Azure/azure-amqp-common-go/v3/rpc"
	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

func Test_recoverableRPCLink_RetryableRPC(t *testing.T) {
	links := &FakeAMQPLinks{}
	innerRPCLink := &fakeRPCLink{}

	recoverableLink := &recoverableRPCLink{
		inner: innerRPCLink,
		links: links,
	}

	innerRPCLink.rpcReturns = []struct {
		Resp *rpc.Response
		Err  error
	}{
		{Resp: nil, Err: errors.New("RPC call failed")},
	}

	resp, err := recoverableLink.RetryableRPC(context.Background(), 1, time.Minute*3, &amqp.Message{})
	require.Nil(t, resp)
	require.EqualError(t, err, "RPC call failed")

	require.EqualError(t, links.recoverParams[0], "RPC call failed")

	require.EqualValues(t, innerRPCLink.closed, 0)
	recoverableLink.Close(context.Background())
	require.EqualValues(t, innerRPCLink.closed, 1)
}

func Test_RecoverableAMQPReceiver(t *testing.T) {
	t.Parallel()

	setup := func() (*FakeAMQPLinks, *RecoverableAMQPReceiver) {
		links := &FakeAMQPLinks{
			Receiver: &fakeAMQPReceiver{},
		}

		recoverableReceiver := &RecoverableAMQPReceiver{
			Links: links,
		}

		return links, recoverableReceiver
	}

	t.Run("DrainCredit", func(t *testing.T) {
		links, recoverableReceiver := setup()

		links.Receiver.drainCreditReturn = errors.New("bad drain call")

		err := recoverableReceiver.DrainCredit(context.Background())

		require.EqualError(t, err, "bad drain call")
		require.EqualError(t, links.recoverParams[0], "bad drain call")
	})

	t.Run("DrainCreditLinksGetFails", func(t *testing.T) {
		links := &FakeAMQPLinks{Err: errors.New("links.get failed")}
		recoverableReceiver := &RecoverableAMQPReceiver{Links: links}

		err := recoverableReceiver.DrainCredit(context.Background())

		require.EqualError(t, err, "links get failed")
		require.EqualError(t, links.recoverParams[0], "links get failed")
	})

	t.Run("IssueCredit", func(t *testing.T) {
		links, recoverableReceiver := setup()

		links.Receiver.issueCreditReturn = errors.New("bad issue credit call")

		err := recoverableReceiver.IssueCredit(context.Background(), 101)

		require.EqualError(t, err, "bad issue credit call")
		require.EqualError(t, links.recoverParams[0], "bad issue credit call")
		require.EqualValues(t, 101, links.Receiver.issueCreditParams[0])
	})

	t.Run("IssueCreditLinksGetFails", func(t *testing.T) {
		links := &FakeAMQPLinks{Err: errors.New("links.get failed")}
		recoverableReceiver := &RecoverableAMQPReceiver{Links: links}

		err := recoverableReceiver.IssueCredit(context.Background(), 101)

		require.EqualError(t, err, "links get failed")
		require.EqualError(t, links.recoverParams[0], "links get failed")
	})

	t.Run("Receive", func(t *testing.T) {
		links, recoverableReceiver := setup()

		links.Receiver.receiveReturns = []struct {
			Message *amqp.Message
			Err     error
		}{
			{Err: errors.New("bad receive call")},
		}

		msg, err := recoverableReceiver.Receive(context.Background())
		require.Nil(t, msg)
		require.EqualError(t, err, "bad receive call")
		require.EqualError(t, links.recoverParams[0], "bad receive call")
	})

	t.Run("ReceiveLinksGetFails", func(t *testing.T) {
		links := &FakeAMQPLinks{Err: errors.New("links.get failed")}
		recoverableReceiver := &RecoverableAMQPReceiver{Links: links}

		msg, err := recoverableReceiver.Receive(context.Background())

		require.Nil(t, msg)
		require.EqualError(t, err, "links.get failed")
		require.EqualError(t, links.recoverParams[0], "links get failed")
	})
}

func Test_RecoverableAMQPSender(t *testing.T) {
	setup := func() (*FakeAMQPLinks, *RecoverableAMQPSender) {
		links := &FakeAMQPLinks{
			Sender: &fakeAMQPSender{},
		}

		recoverableSender := &RecoverableAMQPSender{
			Links: links,
		}

		return links, recoverableSender
	}

	t.Run("Send", func(t *testing.T) {
		links, recoverableSender := setup()

		links.Sender.sendReturns = []error{errors.New("bad send call")}

		err := recoverableSender.Send(context.Background(), &amqp.Message{})
		require.EqualError(t, err, "bad send call")
		require.EqualError(t, links.recoverParams[0], "bad send call")
	})

	t.Run("MaxMessageSize", func(t *testing.T) {
		links, recoverableSender := setup()

		// this one is unique in that it doesn't normally do any IO (we expect the
		// link to be initialized already). However, internally, we do actually check
		// and possibly initialize the link, which means the recoverable version needs
		// to handle errors.
		//
		// But in this case the error will come links.Get(), not from the inner AMQP Sender.

		links.Err = errors.New("links.get failed")

		maxSize, err := recoverableSender.MaxMessageSize(context.Background())
		require.EqualValues(t, 0, maxSize)
		require.EqualError(t, err, "links.get failed")
	})
}

// func TestRecoverableAMQPReceiver_DrainCredit(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		r       *RecoverableAMQPReceiver
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.r.DrainCredit(tt.args.ctx); (err != nil) != tt.wantErr {
// 				t.Errorf("RecoverableAMQPReceiver.DrainCredit() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestRecoverableAMQPReceiver_Receive(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		r       *RecoverableAMQPReceiver
// 		args    args
// 		want    *amqp.Message
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.r.Receive(tt.args.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RecoverableAMQPReceiver.Receive() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("RecoverableAMQPReceiver.Receive() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRecoverableAMQPSender_Send(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		msg *amqp.Message
// 	}
// 	tests := []struct {
// 		name    string
// 		s       *RecoverableAMQPSender
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.s.Send(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
// 				t.Errorf("RecoverableAMQPSender.Send() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestRecoverableAMQPSender_MaxMessageSize(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		s       *RecoverableAMQPSender
// 		args    args
// 		want    uint64
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.s.MaxMessageSize(tt.args.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RecoverableAMQPSender.MaxMessageSize() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("RecoverableAMQPSender.MaxMessageSize() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
