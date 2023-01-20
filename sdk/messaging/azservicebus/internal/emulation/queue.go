// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package emulation

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
)

type Operation struct {
	Op       string
	Entity   string
	LinkName string

	Credits uint32
	M       *amqp.Message
}

type Queue struct {
	name      string
	creditsCh chan int
	src       chan *amqp.Message
	dest      chan *amqp.Message
	pumpFn    sync.Once
	activity  chan Operation
}

func NewQueue(name string, activity chan Operation) *Queue {
	return &Queue{
		creditsCh: make(chan int, 1000),
		dest:      make(chan *amqp.Message, 1000),
		name:      name,
		src:       make(chan *amqp.Message, 1000),
		activity:  activity,
	}
}

func (q *Queue) Send(ctx context.Context, msg *amqp.Message, linkName string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case q.src <- msg:
		q.activity <- Operation{
			Op:       "sent",
			LinkName: linkName,
			Entity:   q.name,

			M: msg,
		}
		return nil
	}
}

func (q *Queue) IssueCredit(credit uint32, linkName string) error {
	q.creditsCh <- int(credit)
	q.activity <- Operation{
		Op:       "credits",
		LinkName: linkName,
		Entity:   q.name,

		Credits: credit,
	}
	return nil
}

func (q *Queue) Receive(ctx context.Context, linkName string) (*amqp.Message, error) {
	q.pumpFn.Do(q.pumpMessages)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg := <-q.dest:
		q.activity <- Operation{
			Op:       "received",
			LinkName: linkName,
			Entity:   q.name,

			M: msg,
		}
		return msg, nil
	}
}

func (q *Queue) AcceptMessage(ctx context.Context, msg *amqp.Message, linkName string) error {
	q.activity <- Operation{
		Op:       "accept",
		LinkName: linkName,
		Entity:   q.name,

		M: msg,
	}
	return nil
}

func (q *Queue) RejectMessage(ctx context.Context, msg *amqp.Message, e *amqp.Error, linkName string) error {
	q.activity <- Operation{
		Op:       "reject",
		LinkName: linkName,
		Entity:   q.name,

		M: msg,
	}
	return nil
}

func (q *Queue) ReleaseMessage(ctx context.Context, msg *amqp.Message, linkName string) error {
	q.activity <- Operation{
		Op:       "release",
		LinkName: linkName,
		Entity:   q.name,

		M: msg,
	}
	return nil
}

func (q *Queue) ModifyMessage(ctx context.Context, msg *amqp.Message, options *amqp.ModifyMessageOptions, linkName string) error {
	q.activity <- Operation{
		Op:       "modify",
		LinkName: linkName,
		Entity:   q.name,

		M: msg,
	}
	return nil
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
