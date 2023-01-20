// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package emulation

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/auth"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/mock"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/sbauth"
	"github.com/golang/mock/gomock"
)

type MockData struct {
	Ctrl     *gomock.Controller
	activity chan Operation

	cbs      sync.Once
	options  *MockDataOptions
	queues   map[string]*Queue
	queuesMu sync.Mutex
}

type MockDataOptions struct {
	NewConnectionFn func(orig *mock.MockAMQPClient, ctx context.Context) (amqpwrap.AMQPClient, error)
	NewReceiverFn   func(orig *mock.MockAMQPReceiverCloser, ctx context.Context, source string, opts *amqp.ReceiverOptions) (*mock.MockAMQPReceiverCloser, error)
	NewSenderFn     func(orig *mock.MockAMQPSenderCloser, ctx context.Context, target string, opts *amqp.SenderOptions) (*mock.MockAMQPSenderCloser, error)
	NewSessionFn    func(orig *mock.MockAMQPSession, ctx context.Context, opts *amqp.SessionOptions) (*mock.MockAMQPSession, error)
}

func NewMockData(t *testing.T, options *MockDataOptions) *MockData {
	if options == nil {
		options = &MockDataOptions{}
	}

	if options.NewReceiverFn == nil {
		options.NewReceiverFn = func(orig *mock.MockAMQPReceiverCloser, ctx context.Context, source string, opts *amqp.ReceiverOptions) (*mock.MockAMQPReceiverCloser, error) {
			return orig, nil
		}
	}

	if options.NewSenderFn == nil {
		options.NewSenderFn = func(orig *mock.MockAMQPSenderCloser, ctx context.Context, target string, opts *amqp.SenderOptions) (*mock.MockAMQPSenderCloser, error) {
			return orig, nil
		}
	}

	if options.NewSessionFn == nil {
		options.NewSessionFn = func(orig *mock.MockAMQPSession, ctx context.Context, opts *amqp.SessionOptions) (*mock.MockAMQPSession, error) {
			return orig, nil
		}
	}

	if options.NewConnectionFn == nil {
		options.NewConnectionFn = func(orig *mock.MockAMQPClient, ctx context.Context) (amqpwrap.AMQPClient, error) {
			return orig, nil
		}
	}

	return &MockData{
		Ctrl:     gomock.NewController(t),
		queues:   map[string]*Queue{},
		options:  options,
		activity: make(chan Operation, 1000+1000),
	}
}

func (md *MockData) NewConnection(ctx context.Context) (amqpwrap.AMQPClient, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	conn := mock.NewMockAMQPClient(md.Ctrl)

	conn.EXPECT().Name().Return(md.nextUniqueName()).AnyTimes()
	conn.EXPECT().NewSession(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, opts *amqp.SessionOptions) (amqpwrap.AMQPSession, error) {
		return md.newSession(ctx, opts, conn.Name())
	}).AnyTimes()
	conn.EXPECT().Close().DoAndReturn(func() error {
		md.activity <- Operation{
			Op: fmt.Sprintf("closeConnection(%s)", conn.Name()),
		}
		return nil
	}).AnyTimes()

	md.activity <- Operation{
		Op: fmt.Sprintf("newConnection(%s)", conn.Name()),
	}

	return conn, nil
}

func (md *MockData) newSession(ctx context.Context, opts *amqp.SessionOptions, connID string) (amqpwrap.AMQPSession, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	sess := mock.NewMockAMQPSession(md.Ctrl)

	sess.EXPECT().NewReceiver(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, source string, opts *amqp.ReceiverOptions) (amqpwrap.AMQPReceiverCloser, error) {
		return md.NewReceiver(ctx, source, opts, connID)
	}).AnyTimes()

	sess.EXPECT().NewSender(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, target string, opts *amqp.SenderOptions) (amqpwrap.AMQPSenderCloser, error) {
		return md.NewSender(ctx, target, opts, connID)
	}).AnyTimes()

	sess.EXPECT().Close(gomock.Any()).Return(nil).AnyTimes()

	return md.options.NewSessionFn(sess, ctx, opts)
}

func (md *MockData) NewReceiver(ctx context.Context, source string, opts *amqp.ReceiverOptions, connID string) (amqpwrap.AMQPReceiverCloser, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if opts == nil {
		opts = &amqp.ReceiverOptions{}
	}

	rcvr := mock.NewMockAMQPReceiverCloser(md.Ctrl)
	rcvr.EXPECT().LinkName().Return(connID + "|" + md.nextUniqueName()).AnyTimes()

	targetAddress := opts.TargetAddress

	log.Printf("NewMockAMQPSenderCloser(%s, alias: %s)", source, targetAddress)

	q := md.upsertQueue(targetAddress)
	cbs := md.upsertQueue(source)

	if source == "$cbs" {
		md.activity <- Operation{
			Op:       "newCBSReceiver",
			Entity:   targetAddress,
			LinkName: rcvr.LinkName(),
		}

		md.cbs.Do(func() {
			go func() { md.cbsRouter(context.Background(), cbs, md.getQueue) }()
		})
	} else {
		md.activity <- Operation{
			Op:       "newReceiver",
			Entity:   source,
			LinkName: rcvr.LinkName(),
		}
	}

	rcvr.EXPECT().Receive(gomock.Any()).DoAndReturn(func(ctx context.Context) (*amqp.Message, error) {
		return q.Receive(ctx, rcvr.LinkName())
	}).AnyTimes()

	rcvr.EXPECT().Close(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		md.activity <- Operation{
			Op:       "close",
			Entity:   source,
			LinkName: rcvr.LinkName(),
		}
		return nil
	}).AnyTimes()

	rcvr.EXPECT().AcceptMessage(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msg *amqp.Message) error {
		return q.AcceptMessage(ctx, msg, rcvr.LinkName())
	}).AnyTimes()

	rcvr.EXPECT().RejectMessage(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msg *amqp.Message, e *amqp.Error) error {
		return q.RejectMessage(ctx, msg, e, rcvr.LinkName())
	}).AnyTimes()

	rcvr.EXPECT().ReleaseMessage(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msg *amqp.Message) error {
		return q.ReleaseMessage(ctx, msg, rcvr.LinkName())
	}).AnyTimes()

	rcvr.EXPECT().ModifyMessage(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msg *amqp.Message, options *amqp.ModifyMessageOptions) error {
		return q.ModifyMessage(ctx, msg, options, rcvr.LinkName())
	}).AnyTimes()

	if opts.ManualCredits {
		rcvr.EXPECT().IssueCredit(gomock.Any()).DoAndReturn(func(credit uint32) error {
			return q.IssueCredit(credit, rcvr.LinkName())
		}).AnyTimes()
	} else {
		// assume unlimited credits for this receiver - the AMQP stack is going to take care of replenishing credits.
		_ = q.IssueCredit(math.MaxUint32, rcvr.LinkName())
	}

	return md.options.NewReceiverFn(rcvr, ctx, source, opts)
}

func (md *MockData) NewSender(ctx context.Context, target string, opts *amqp.SenderOptions, connID string) (amqpwrap.AMQPSenderCloser, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	log.Printf("NewMockAMQPSenderCloser(%s)", target)

	sender := mock.NewMockAMQPSenderCloser(md.Ctrl)
	sender.EXPECT().LinkName().Return(connID + "|" + md.nextUniqueName()).AnyTimes()

	// this should work fine even for RPC links like $cbs or $management
	q := md.upsertQueue(target)
	sender.EXPECT().Send(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msg *amqp.Message) error {
		return q.Send(ctx, msg, sender.LinkName())
	}).AnyTimes()

	sender.EXPECT().Close(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		md.activity <- Operation{
			Op:       "close",
			Entity:   target,
			LinkName: sender.LinkName(),
		}
		return nil
	}).AnyTimes()

	return md.options.NewSenderFn(sender, ctx, target, opts)
}

func (md *MockData) GetActivity() []Operation {
	var ops []Operation

Loop:
	for {
		select {
		case op := <-md.activity:
			ops = append(ops, op)
		default:
			break Loop
		}
	}

	return ops
}

func (md *MockData) upsertQueue(name string) *Queue {
	md.queuesMu.Lock()
	defer md.queuesMu.Unlock()

	q := md.queues[name]

	if q == nil {
		q = NewQueue(name, md.activity)
		md.queues[name] = q
	}

	return q
}

func (md *MockData) getQueue(name string) *Queue {
	md.queuesMu.Lock()
	defer md.queuesMu.Unlock()

	return md.queues[name]
}

func (md *MockData) AllQueues() map[string]*Queue {
	md.queuesMu.Lock()
	defer md.queuesMu.Unlock()

	m := map[string]*Queue{}

	for k, v := range md.queues {
		m[k] = v
	}

	return m
}

func (md *MockData) nextUniqueName() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (md *MockData) NewTokenProvider() auth.TokenProvider {
	tc := mock.NewMockTokenCredential(md.Ctrl)

	var tokenCounter int64

	tc.EXPECT().GetToken(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
		tc := atomic.AddInt64(&tokenCounter, 1)

		return azcore.AccessToken{
			Token:     fmt.Sprintf("Token-%d", tc),
			ExpiresOn: time.Now().Add(10 * time.Minute),
		}, nil
	}).AnyTimes()

	return sbauth.NewTokenProvider(tc)
}

func (md *MockData) cbsRouter(ctx context.Context, in *Queue, getQueue func(name string) *Queue) {
	_ = in.IssueCredit(math.MaxUint32, "cbs-internal")

	for {
		msg, err := in.Receive(ctx, "cbs-internal")

		if err != nil {
			log.Printf("CBS router closed: %s", err.Error())
			break
		}

		// route response to the right spot
		replyTo := *msg.Properties.ReplyTo

		out := getQueue(replyTo)

		log.Printf("Sending CBS reply to %s", replyTo)

		// assume auth is valid.
		err = out.Send(ctx, &amqp.Message{
			Properties: &amqp.MessageProperties{
				CorrelationID: msg.Properties.MessageID,
			},
			ApplicationProperties: map[string]any{
				"statusCode": int32(200),
			},
		}, "cbs-internal")

		if err != nil {
			log.Printf("CBS router closed: %s", err.Error())
			break
		}
	}
}

// func (md *mockData) mgmtRouter(ctx context.Context, in *Queue, getQueue func(name string) *Queue) {
// 	in.IssueCredit(math.MaxInt32)

// 	for {
// 		msg, err := in.Receive(ctx)

// 		if err != nil {
// 			log.Printf("CBS Processor closed: %s", err.Error())
// 			break
// 		}

// 		// route response to the right spot
// 		replyTo := *msg.Properties.ReplyTo

// 		out := getQueue(replyTo)

// 		// assume auth is valid.
// 		err = out.Send(ctx, &amqp.Message{
// 			Properties: &amqp.MessageProperties{
// 				CorrelationID: msg.Properties.MessageID,
// 			},
// 			ApplicationProperties: map[string]any{
// 				"statusCode": 200,
// 			},
// 		})

// 		if err != nil {
// 			log.Printf("CBS Processor closed: %s", err.Error())
// 			break
// 		}
// 	}
// }
