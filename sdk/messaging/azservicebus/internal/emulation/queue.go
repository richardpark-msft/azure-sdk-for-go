// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package emulation

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
)

// type MockConnState struct {
// 	client *mock.MockAMQPClient

// 	// NOTE: some links can show up in multiple places. For instance, with Receivers
// 	// we'll set the TargetAddress which means the receiver is addressable via it's old
// 	// link ID and also the supplied link ID.
// 	linksMu   sync.Mutex
// 	senders   map[string]chan *amqp.Message
// 	receivers map[string]chan *amqp.Message
// }

// type MockServiceBus struct {
// 	ctrl *gomock.Controller

// 	newClientFn func(ctx context.Context) (amqpwrap.AMQPClient, error)

// 	// conns maps all active connections to their unique  connection ID
// 	connsMu sync.Mutex
// 	conns   []*MockConnState

// 	// for now readers and writers will use the same channel
// 	// TODO: expand this, allow for credits, etc...
// 	entities map[string]chan *amqp.Message

// 	UpdateReceiverMock func(mock *mock.MockAMQPReceiverCloser, source string, opts *amqp.ReceiverOptions) error
// }

// func (msb *MockServiceBus) getConn(idx int) *MockConnState {
// 	msb.connsMu.Lock()
// 	defer msb.connsMu.Unlock()

// 	return msb.conns[idx]
// }

// func (msb *MockServiceBus) getSenderChannel(connIdx int, entity string) chan<- *amqp.Message {
// 	// if entity == "$cbs" {
// 	// 	cs := msb.getConn(connIdx)
// 	// } else if entity == "$management" {
// 	// 	cs := msb.getConn(connIdx)
// 	// } else {
// 	// 	cs := msb.getConn(connIdx)
// 	// }

// 	return nil
// }

// func (msb *MockServiceBus) getReceiverChannel(entity string, opts *amqp.ReceiverOptions) <-chan *amqp.Message {
// 	// if entity == "$cbs" {
// 	// 	// get the $cbs channel for this connection.
// 	// }

// 	// connState.receivers[rcvr.LinkName()] = ch

// 	// // from what I can tell this is how the RPC links can address a specific link
// 	// // to respond to. This also appears optional, so for now I'm allowing both.
// 	// if opts.TargetAddress != "" {
// 	// 	connState.receivers[opts.TargetAddress] = ch
// 	// }

// 	return nil
// }

// func (msb *MockServiceBus) NewClient(ctx context.Context) (amqpwrap.AMQPClient, error) {
// 	msb.connsMu.Lock()
// 	connState := &MockConnState{}
// 	msb.conns = append(msb.conns, connState)
// 	connIdx := len(msb.conns) - 1
// 	msb.connsMu.Unlock()

// 	conn := mock.NewMockAMQPClient(msb.ctrl)

// 	newReceiverFn := func(ctx context.Context, source string, opts *amqp.ReceiverOptions) (amqpwrap.AMQPReceiverCloser, error) {
// 		select {
// 		case <-ctx.Done():
// 			return nil, ctx.Err()
// 		default:
// 		}

// 		rcvr := mock.NewMockAMQPReceiverCloser(msb.ctrl)
// 		ch := msb.getReceiverChannel(source, opts)

// 		connState.linksMu.Lock()
// 		linkIdx := len(connState.receivers) - 1
// 		connState.linksMu.Unlock()

// 		linkID := msb.formatLinkID(connIdx, linkIdx)
// 		rcvr.EXPECT().LinkName().Return(linkID).AnyTimes()

// 		if msb.UpdateReceiverMock != nil {
// 			if err := msb.UpdateReceiverMock(rcvr, source, opts); err != nil {
// 				return nil, err
// 			}
// 		}

// 		creditsCh := make(chan uint32, 1000)

// 		rcvr.EXPECT().IssueCredit(gomock.Any()).DoAndReturn(func(credits uint32) error {
// 			creditsCh <- credits
// 			return nil
// 		})

// 		rcvr.EXPECT().Receive(gomock.Any()).DoAndReturn(func(ctx context.Context) (*amqp.Message, error) {
// 			select {
// 			case <-ctx.Done():
// 				return nil, ctx.Err()
// 			case msg := <-receivedCh:
// 				log.Writef("Mock", "Received %#v from %s\n", msg, source)
// 				return msg, nil
// 			}
// 		}).AnyTimes()

// 		return rcvr, nil
// 	}

// 	newSenderFn := func(ctx context.Context, target string, opts *amqp.SenderOptions) (amqpwrap.AMQPSenderCloser, error) {
// 		select {
// 		case <-ctx.Done():
// 			return nil, ctx.Err()
// 		default:
// 		}

// 		sender := mock.NewMockAMQPSenderCloser(msb.ctrl)

// 		connState.linksMu.Lock()
// 		connState.senders = append(connState.senders, sender)
// 		linkIdx := len(connState.senders) - 1
// 		connState.linksMu.Unlock()

// 		linkID := msb.formatLinkID(connIdx, linkIdx)
// 		sender.EXPECT().LinkName().Return(linkID).AnyTimes()

// 		ch := msb.getSenderChannel(target)

// 		sender.EXPECT().Send(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msg *amqp.Message) error {
// 			select {
// 			case <-ctx.Done():
// 				return ctx.Err()
// 			default:
// 				ch <- msg
// 				return nil
// 			}
// 		}).AnyTimes()

// 		return sender, nil
// 	}

// 	conn.EXPECT().NewSession(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, opts *amqp.SessionOptions) (amqpwrap.AMQPSession, error) {
// 		select {
// 		case <-ctx.Done():
// 			return nil, ctx.Err()
// 		default:
// 		}

// 		sess := mock.NewMockAMQPSession(msb.ctrl)

// 		if args.UpdateSessionMock != nil {
// 			if err := args.UpdateSessionMock(sess); err != nil {
// 				return nil, err
// 			}
// 		}

// 		sess.EXPECT().NewReceiver(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(newReceiverFn).AnyTimes()
// 		sess.EXPECT().NewSender(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(newSenderFn).AnyTimes()

// 		return sess, nil
// 	}).AnyTimes()

// 	return conn, nil
// }

// func (msb *MockServiceBus) NewNamespace(t *testing.T, ctrl *gomock.Controller, args MockNamespaceArgs) *Namespace {
// 	tc := mock.NewMockTokenCredential(ctrl)

// 	var tokenCounter int64

// 	tc.EXPECT().GetToken(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
// 		tc := atomic.AddInt64(&tokenCounter, 1)

// 		return azcore.AccessToken{
// 			Token:     fmt.Sprintf("Token-%d", tc),
// 			ExpiresOn: time.Now().Add(10 * time.Minute),
// 		}, nil
// 	}).AnyTimes()

// 	ns, err := NewNamespace(NamespaceWithTokenCredential("example.servicebus.windows.net", tc))
// 	require.NoError(t, err)

// 	var connCounter int64
// 	var linkCounter int64

// 	m := sync.Map{}

// 	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {

// 		return conn, nil
// 	}

// 	return ns
// }

// func (msb *MockServiceBus) formatLinkID(connIdx int, linkIdx int) string {
// 	return fmt.Sprintf("c:%d l:%d receiver", connIdx, linkIdx)
// }

// func (msb *MockServiceBus) parseLinkID(linkID string) (connIdx int, linkIdx int, linkType string) {
// 	pieces := strings.Split(linkID, " ")

// 	connIdx, err := strconv.ParseInt(pieces[0], 10, 64)

// 	if err != nil {
// 		panic(err)
// 	}

// 	linkIdx, err := strconv.ParseInt(pieces[1], 10, 64)

// 	if err != nil {
// 		panic(err)
// 	}

// 	linkType := pieces[2]

// 	return int(connIdx), int(linkIdx), linkType
// }

// func NewManagementRPC() *mockRPC {
// 	return &mockRPC{
// 		recvChannelMap: map[string]chan *amqp.Message{},
// 		sendCh:         make(chan *amqp.Message, 10000),
// 		processFn:      func(request, response *amqp.Message) {},
// 	}
// }

// func NewCBSRPC() *mockRPC {
// 	return &mockRPC{
// 		recvChannelMap: map[string]chan *amqp.Message{},
// 		sendCh:         make(chan *amqp.Message, 10000),
// 		processFn: func(request, response *amqp.Message) {
// 		},
// 	}
// }

// type mockRPC struct {
// 	// recvChannelMap is a list of link IDs to the channel we
// 	// gave them when they were created. This allows us to route
// 	// sent messages to the appropriate receiver
// 	mu             sync.Mutex
// 	recvChannelMap map[string]chan *amqp.Message
// 	sendCh         chan *amqp.Message

// 	processFn func(request *amqp.Message, response *amqp.Message)
// }

// func (c *mockRPC) Process(ctx context.Context) {
// Loop:
// 	for {
// 		select {
// 		case request := <-c.sendCh:
// 			// TODO: process and validate auth request.
// 			response := &amqp.Message{
// 				Properties: &amqp.MessageProperties{
// 					CorrelationID: request.Properties.MessageID,
// 				},
// 			}

// 			// find the right channel to send to
// 			c.mu.Lock()
// 			ch, ok := c.recvChannelMap[*response.Properties.ReplyTo]

// 			if !ok {
// 				panic(fmt.Sprintf("No reply channel for %s", response.Properties.ReplyTo))
// 			}
// 			c.mu.Unlock()

// 			c.processFn(request, response)
// 			ch <- response
// 		case <-ctx.Done():
// 			break Loop
// 		}
// 	}
// }

// func (c *mockRPC) NewReceiver(ctrl *gomock.Controller, ctx context.Context, source string, opts *amqp.ReceiverOptions) (*mock.MockAMQPReceiverCloser, error) {
// 	c.mu.Lock()
// 	ch, ok := c.recvChannelMap[opts.TargetAddress]

// 	if !ok {
// 		ch = make(chan *amqp.Message, opts.Credit)
// 		c.recvChannelMap[opts.TargetAddress] = ch
// 	}

// 	c.mu.Unlock()

// 	return NewMockAMQPReceiverCloser(ctrl, ch, ctx, source, opts)
// }

// func (c *mockRPC) NewSender(ctrl *gomock.Controller, ctx context.Context, target string, opts *amqp.SenderOptions) (*mock.MockAMQPSenderCloser, error) {
// 	return NewMockAMQPSenderCloser(ctrl, c.sendCh, ctx, target, opts)
// }

// func NewMockAMQPReceiverCloser(ctrl *gomock.Controller, ch <-chan *amqp.Message, ctx context.Context, source string, opts *amqp.ReceiverOptions) (*mock.MockAMQPReceiverCloser, error) {
// 	rcvr := mock.NewMockAMQPReceiverCloser(ctrl)

// 	rcvr.EXPECT().Receive(gomock.Any()).DoAndReturn(func(ctx context.Context) (*amqp.Message, error) {
// 		select {
// 		case <-ctx.Done():
// 			return nil, ctx.Err()
// 		case msg := <-ch:
// 			// TODO: later let's do some credit tracking as well.
// 			log.Writef("Mock", "Received %#v from %s\n", msg, source)
// 			return msg, nil
// 		}
// 	}).AnyTimes()

// 	return rcvr, nil
// }

// func NewMockAMQPSenderCloser(ctrl *gomock.Controller, ch chan<- *amqp.Message, ctx context.Context, target string, opts *amqp.SenderOptions) (*mock.MockAMQPSenderCloser, error) {
// 	sender := mock.NewMockAMQPSenderCloser(ctrl)

// 	sender.EXPECT().Send(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, msg *amqp.Message) error {
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		default:
// 			log.Writef("Mock", "Sent %#v to %s\n", msg, target)
// 			ch <- msg
// 			return nil
// 		}
// 	}).AnyTimes()

// 	return sender, nil
// }

type Queue struct {
	creditsCh chan int
	src       chan *amqp.Message
	dest      chan *amqp.Message
	pumpFn    sync.Once
}

func NewQueue() *Queue {
	return &Queue{
		creditsCh: make(chan int, 1000),
		src:       make(chan *amqp.Message, 1000),
		dest:      make(chan *amqp.Message, 1000),
	}
}

func (q *Queue) Send(ctx context.Context, msg *amqp.Message) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case q.src <- msg:
		return nil
	}
}

func (q *Queue) IssueCredit(credit int) {
	q.creditsCh <- credit
}

func (q *Queue) Receive(ctx context.Context) (*amqp.Message, error) {
	q.pumpFn.Do(q.pumpMessages)

	select {
	case msg := <-q.dest:
		return msg, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (q *Queue) pumpMessages() {
	go func() {
		for {
			credit := <-q.creditsCh

			if credit == 0 {
				break
			}

			for i := 0; i < credit; i++ {
				msg := <-q.src

				if msg == nil {
					break
				}

				q.dest <- msg
			}
		}
	}()
}

func (q *Queue) Close() {
	close(q.creditsCh)
	close(q.src)
	close(q.dest)
}
