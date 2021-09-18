// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	common "github.com/Azure/azure-amqp-common-go/v3"
	"github.com/Azure/azure-amqp-common-go/v3/rpc"
	"github.com/Azure/azure-amqp-common-go/v3/uuid"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/spans"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

const (
	amqpRetryDefaultTimes    int           = 3
	amqpRetryDefaultDelay    time.Duration = time.Second
	amqpRetryBusyServerDelay time.Duration = 10 * time.Second
)

// Error Conditions
const (
	// Service Bus Errors
	errorServerBusy         amqp.ErrorCondition = "com.microsoft:server-busy"
	errorTimeout            amqp.ErrorCondition = "com.microsoft:timeout"
	errorOperationCancelled amqp.ErrorCondition = "com.microsoft:operation-cancelled"
	errorContainerClose     amqp.ErrorCondition = "com.microsoft:container-close"
)

type (
	mgmtClient struct {
		ns             *internal.Namespace
		managementPath string

		clientMu sync.RWMutex
		link     *rpc.Link

		sessionID          *string
		isSessionFilterSet bool
	}
)

// tracing operation names
type rpcSpanName string

const (
	spanNameRenewLock              rpcSpanName = "sb.mgmt.RenewLock"
	spanNameDefer                  rpcSpanName = "sb.mgmt.Defer"
	spanNameSettle                 rpcSpanName = "sb.mgmt.SendDisposition"
	spanNameScheduleMessage        rpcSpanName = "sb.mgmt.Schedule"
	spanNameCancelScheduledMessage rpcSpanName = "sb.mgmt.CancelScheduled"
	spanNameRecover                rpcSpanName = "sb.mgmt.Recover"
	spanNameTryRecover             rpcSpanName = "sb.mgmt.TryRecover"
)

func newRPCClient(ctx context.Context, entityPath string, ns *internal.Namespace) (*mgmtClient, error) {
	r := &mgmtClient{
		ns:             ns,
		managementPath: fmt.Sprintf("%s/$management", entityPath),
	}

	return r, nil
}

// Recover will attempt to close the current session and link, then rebuild them
func (mc *mgmtClient) recover(ctx context.Context) error {
	mc.clientMu.Lock()
	defer mc.clientMu.Unlock()

	ctx, span := mc.startSpanFromContext(ctx, string(spanNameRecover))
	defer span.End()

	if mc.link != nil {
		if err := mc.link.Close(ctx); err != nil {
			tab.For(ctx).Debug(fmt.Sprintf("Error while closing old link in recovery: %s", err.Error()))
		}
		mc.link = nil
	}

	if _, err := mc.getLinkWithoutLock(ctx); err != nil {
		return err
	}

	return nil
}

// getLinkWithoutLock returns the currently cached link (or creates a new one)
func (mc *mgmtClient) getLinkWithoutLock(ctx context.Context) (*rpc.Link, error) {
	if mc.link != nil {
		return mc.link, nil
	}

	client, err := mc.ns.GetAMQPClient(ctx)

	if err != nil {
		return nil, err
	}

	mc.link, err = rpc.NewLink(client, mc.managementPath)

	if err != nil {
		return nil, err
	}

	return mc.link, nil
}

// closeWithoutLock closes the currently held link and nil's it.
func (mc *mgmtClient) closeWithoutLock(ctx context.Context) error {
	if mc.link == nil {
		return nil
	}

	var l *rpc.Link
	l, mc.link = mc.link, nil

	if err := l.Close(ctx); err != nil {
		tab.For(ctx).Debug(fmt.Sprintf("Error while closing old link in recovery: %s", err.Error()))
		return err
	}

	return nil
}

// Close will close the AMQP connection
func (mc *mgmtClient) Close(ctx context.Context) error {
	mc.clientMu.Lock()
	defer mc.clientMu.Unlock()
	err := mc.link.Close(ctx)
	mc.link = nil
	return err
}

// creates a new link and sends the RPC request, recovering and retrying on certain AMQP errors
func (mc *mgmtClient) doRPCWithRetry(ctx context.Context, msg *amqp.Message, times int, delay time.Duration, opts ...rpc.LinkOption) (*rpc.Response, error) {
	// track the number of times we attempt to perform the RPC call.
	// this is to avoid a potential infinite loop if the returned error
	// is always transient and Recover() doesn't fail.
	sendCount := 0

	for {
		mc.clientMu.RLock()
		rpcLink, err := mc.getLinkWithoutLock(ctx)
		mc.clientMu.RUnlock()

		var rsp *rpc.Response

		if err == nil {
			rsp, err = rpcLink.RetryableRPC(ctx, times, delay, msg)

			if err == nil {
				return rsp, err
			}
		}

		if sendCount >= amqpRetryDefaultTimes || !isAMQPTransientError(ctx, err) {
			return nil, err
		}
		sendCount++
		// if we get here, recover and try again
		tab.For(ctx).Debug("recovering RPC connection")

		_, retryErr := common.Retry(amqpRetryDefaultTimes, amqpRetryDefaultDelay, func() (interface{}, error) {
			ctx, sp := mc.startProducerSpanFromContext(ctx, string(spanNameTryRecover))
			defer sp.End()

			if err := mc.recover(ctx); err == nil {
				tab.For(ctx).Debug("recovered RPC connection")
				return nil, nil
			}
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
				return nil, common.Retryable(err.Error())
			}
		})

		if retryErr != nil {
			tab.For(ctx).Debug("RPC recovering retried, but error was unrecoverable")
			return nil, retryErr
		}
	}
}

// returns true if the AMQP error is considered transient
func isAMQPTransientError(ctx context.Context, err error) bool {
	// always retry on a detach error
	var amqpDetach *amqp.DetachError
	if errors.As(err, &amqpDetach) {
		return true
	}
	// for an AMQP error, only retry depending on the condition
	var amqpErr *amqp.Error
	if errors.As(err, &amqpErr) {
		switch amqpErr.Condition {
		case errorServerBusy, errorTimeout, errorOperationCancelled, errorContainerClose:
			return true
		default:
			tab.For(ctx).Debug(fmt.Sprintf("isAMQPTransientError: condition %s is not transient", amqpErr.Condition))
			return false
		}
	}
	tab.For(ctx).Debug(fmt.Sprintf("isAMQPTransientError: %T is not transient", err))
	return false
}

func (mc *mgmtClient) ReceiveDeferred(ctx context.Context, mode ReceiveMode, sequenceNumbers ...int64) ([]*amqp.Message, error) {
	ctx, span := spans.ForReceiver(ctx, string(spanNameDefer), mc.managementPath, mc.ns.GetHostname())
	defer span.End()

	const messagesField, messageField = "messages", "message"

	backwardsMode := uint32(0)
	if mode == PeekLock {
		backwardsMode = 1
	}

	values := map[string]interface{}{
		"sequence-numbers":     sequenceNumbers,
		"receiver-settle-mode": uint32(backwardsMode), // pick up messages with peek lock
	}

	var opts []rpc.LinkOption
	if mc.isSessionFilterSet {
		opts = append(opts, rpc.LinkWithSessionFilter(mc.sessionID))
		values["session-id"] = mc.sessionID
	}

	msg := &amqp.Message{
		ApplicationProperties: map[string]interface{}{
			"operation": "com.microsoft:receive-by-sequence-number",
		},
		Value: values,
	}

	rsp, err := mc.doRPCWithRetry(ctx, msg, 5, 5*time.Second, opts...)
	if err != nil {
		tab.For(ctx).Error(err)
		return nil, err
	}

	if rsp.Code == 204 {
		return nil, internal.ErrNoMessages{}
	}

	// Deferred messages come back in a relatively convoluted manner:
	// a map (always with one key: "messages")
	// 	of arrays
	// 		of maps (always with one key: "message")
	// 			of an array with raw encoded Service Bus messages
	val, ok := rsp.Message.Value.(map[string]interface{})
	if !ok {
		return nil, internal.NewErrIncorrectType(messageField, map[string]interface{}{}, rsp.Message.Value)
	}

	rawMessages, ok := val[messagesField]
	if !ok {
		return nil, internal.ErrMissingField(messagesField)
	}

	messages, ok := rawMessages.([]interface{})
	if !ok {
		return nil, internal.NewErrIncorrectType(messagesField, []interface{}{}, rawMessages)
	}

	transformedMessages := make([]*amqp.Message, len(messages))
	for i := range messages {
		rawEntry, ok := messages[i].(map[string]interface{})
		if !ok {
			return nil, internal.NewErrIncorrectType(messageField, map[string]interface{}{}, messages[i])
		}

		rawMessage, ok := rawEntry[messageField]
		if !ok {
			return nil, internal.ErrMissingField(messageField)
		}

		marshaled, ok := rawMessage.([]byte)
		if !ok {
			return nil, new(internal.ErrMalformedMessage)
		}

		var rehydrated amqp.Message
		err = rehydrated.UnmarshalBinary(marshaled)
		if err != nil {
			return nil, err
		}

		transformedMessages[i] = &rehydrated
	}

	return transformedMessages, nil
}

func (mc *mgmtClient) GetNextPage(ctx context.Context, fromSequenceNumber int64, messageCount int32) ([]*amqp.Message, error) {
	ctx, span := spans.StartConsumerSpanFromContext(ctx, "sb.rpcClient.GetNextPage")
	defer span.End()

	const messagesField, messageField = "messages", "message"

	msg := &amqp.Message{
		ApplicationProperties: map[string]interface{}{
			"operation": "com.microsoft:peek-message",
		},
		Value: map[string]interface{}{
			"from-sequence-number": fromSequenceNumber,
			"message-count":        messageCount,
		},
	}

	if deadline, ok := ctx.Deadline(); ok {
		msg.ApplicationProperties["server-timeout"] = uint(time.Until(deadline) / time.Millisecond)
	}

	rsp, err := mc.doRPCWithRetry(ctx, msg, 5, 5*time.Second)
	if err != nil {
		tab.For(ctx).Error(err)
		return nil, err
	}

	if rsp.Code == 204 {
		return nil, internal.ErrNoMessages{}
	}

	// Peeked messages come back in a relatively convoluted manner:
	// a map (always with one key: "messages")
	// 	of arrays
	// 		of maps (always with one key: "message")
	// 			of an array with raw encoded Service Bus messages
	val, ok := rsp.Message.Value.(map[string]interface{})
	if !ok {
		err = internal.NewErrIncorrectType(messageField, map[string]interface{}{}, rsp.Message.Value)
		tab.For(ctx).Error(err)
		return nil, err
	}

	rawMessages, ok := val[messagesField]
	if !ok {
		err = internal.ErrMissingField(messagesField)
		tab.For(ctx).Error(err)
		return nil, err
	}

	messages, ok := rawMessages.([]interface{})
	if !ok {
		err = internal.NewErrIncorrectType(messagesField, []interface{}{}, rawMessages)
		tab.For(ctx).Error(err)
		return nil, err
	}

	transformedMessages := make([]*amqp.Message, len(messages))
	for i := range messages {
		rawEntry, ok := messages[i].(map[string]interface{})
		if !ok {
			err = internal.NewErrIncorrectType(messageField, map[string]interface{}{}, messages[i])
			tab.For(ctx).Error(err)
			return nil, err
		}

		rawMessage, ok := rawEntry[messageField]
		if !ok {
			err = internal.ErrMissingField(messageField)
			tab.For(ctx).Error(err)
			return nil, err
		}

		marshaled, ok := rawMessage.([]byte)
		if !ok {
			err = new(internal.ErrMalformedMessage)
			tab.For(ctx).Error(err)
			return nil, err
		}

		var rehydrated amqp.Message
		err = rehydrated.UnmarshalBinary(marshaled)
		if err != nil {
			tab.For(ctx).Error(err)
			return nil, err
		}

		transformedMessages[i] = &rehydrated

		// transformedMessages[i], err = MessageFromAMQPMessage(&rehydrated)
		// if err != nil {
		// 	tab.For(ctx).Error(err)
		// 	return nil, err
		// }

		// transformedMessages[i].useSession = r.isSessionFilterSet
		// transformedMessages[i].sessionID = r.sessionID
	}

	// This sort is done to ensure that folks wanting to peek messages in sequence order may do so.
	// sort.Slice(transformedMessages, func(i, j int) bool {
	// 	iSeq := *transformedMessages[i].SystemProperties.SequenceNumber
	// 	jSeq := *transformedMessages[j].SystemProperties.SequenceNumber
	// 	return iSeq < jSeq
	// })

	return transformedMessages, nil
}

// RenewLocks renews the locks in a single 'com.microsoft:renew-lock' operation.
// NOTE: this function assumes all the messages received on the same link.
func (mc *mgmtClient) RenewLocks(ctx context.Context, messages ...*ReceivedMessage) (err error) {
	ctx, span := spans.StartConsumerSpanFromContext(ctx, "sb.RenewLocks")
	defer span.End()

	if len(messages) == 0 {
		return nil
	}

	var linkName string = messages[0].linkName

	lockTokens := make([]amqp.UUID, 0, len(messages))
	for _, m := range messages {
		if m.LockToken == nil {
			return fmt.Errorf("failed: message %s has nil lock token, cannot renew lock", m.ID)
		}

		amqpLockToken := amqp.UUID(*m.LockToken)
		lockTokens = append(lockTokens, amqpLockToken)
	}

	renewRequestMsg := &amqp.Message{
		ApplicationProperties: map[string]interface{}{
			"operation": "com.microsoft:renew-lock",
		},
		Value: map[string]interface{}{
			"lock-tokens": lockTokens,
		},
	}

	if linkName != "" {
		renewRequestMsg.ApplicationProperties["associated-link-name"] = linkName
	}

	response, err := mc.doRPCWithRetry(ctx, renewRequestMsg, 3, 1*time.Second)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	if response.Code != 200 {
		err := fmt.Errorf("error renewing locks: %v", response.Description)
		tab.For(ctx).Error(err)
		return err
	}

	return nil
}

func (mc *mgmtClient) SendDisposition(ctx context.Context, lockToken *amqp.UUID, state disposition) error {
	ctx, span := spans.StartConsumerSpanFromContext(ctx, "sb.rpcClient.SendDisposition")
	defer span.End()

	if lockToken == nil {
		err := errors.New("lock token on the message is not set, thus cannot send disposition")
		tab.For(ctx).Error(err)
		return err
	}

	var opts []rpc.LinkOption
	value := map[string]interface{}{
		"disposition-status": string(state.Status),
		"lock-tokens":        []amqp.UUID{*lockToken},
	}

	if state.DeadLetterReason != nil {
		value["deadletter-reason"] = state.DeadLetterReason
	}

	if state.DeadLetterDescription != nil {
		value["deadletter-description"] = state.DeadLetterDescription
	}

	msg := &amqp.Message{
		ApplicationProperties: map[string]interface{}{
			"operation": "com.microsoft:update-disposition",
		},
		Value: value,
	}

	// no error, then it was successful
	_, err := mc.doRPCWithRetry(ctx, msg, 5, 5*time.Second, opts...)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	return nil
}

// ScheduleAt will send a batch of messages to a Queue, schedule them to be enqueued, and return the sequence numbers
// that can be used to cancel each message.
func (mc *mgmtClient) ScheduleAt(ctx context.Context, enqueueTime time.Time, messages ...*Message) ([]int64, error) {
	ctx, span := spans.StartConsumerSpanFromContext(ctx, string(spanNameScheduleMessage))
	defer span.End()

	if len(messages) <= 0 {
		return nil, errors.New("expected one or more messages")
	}

	transformed := make([]interface{}, 0, len(messages))
	for i := range messages {
		enqueueTimeAsUTC := enqueueTime.UTC()
		messages[i].ScheduledEnqueueTime = &enqueueTimeAsUTC

		if messages[i].ID == "" {
			id, err := uuid.NewV4()
			if err != nil {
				return nil, err
			}
			messages[i].ID = id.String()
		}

		rawAmqp, err := messages[i].toAMQPMessage()
		if err != nil {
			return nil, err
		}
		encoded, err := rawAmqp.MarshalBinary()
		if err != nil {
			return nil, err
		}

		individualMessage := map[string]interface{}{
			"message-id": messages[i].ID,
			"message":    encoded,
		}
		if messages[i].SessionID != nil {
			individualMessage["session-id"] = *messages[i].SessionID
		}
		if partitionKey := messages[i].PartitionKey; partitionKey != nil {
			individualMessage["partition-key"] = *partitionKey
		}
		if viaPartitionKey := messages[i].TransactionPartitionKey; viaPartitionKey != nil {
			individualMessage["via-partition-key"] = *viaPartitionKey
		}

		transformed = append(transformed, individualMessage)
	}

	msg := &amqp.Message{
		ApplicationProperties: map[string]interface{}{
			"operation": "com.microsoft:schedule-message",
		},
		Value: map[string]interface{}{
			"messages": transformed,
		},
	}

	if deadline, ok := ctx.Deadline(); ok {
		msg.ApplicationProperties["com.microsoft:server-timeout"] = uint(time.Until(deadline) / time.Millisecond)
	}

	resp, err := mc.doRPCWithRetry(ctx, msg, 5, 5*time.Second)
	if err != nil {
		tab.For(ctx).Error(err)
		return nil, err
	}

	if resp.Code != 200 {
		return nil, internal.ErrAMQP(*resp)
	}

	retval := make([]int64, 0, len(messages))
	if rawVal, ok := resp.Message.Value.(map[string]interface{}); ok {
		const sequenceFieldName = "sequence-numbers"
		if rawArr, ok := rawVal[sequenceFieldName]; ok {
			if arr, ok := rawArr.([]int64); ok {
				for i := range arr {
					retval = append(retval, arr[i])
				}
				return retval, nil
			}
			return nil, internal.NewErrIncorrectType(sequenceFieldName, []int64{}, rawArr)
		}
		return nil, internal.ErrMissingField(sequenceFieldName)
	}
	return nil, internal.NewErrIncorrectType("value", map[string]interface{}{}, resp.Message.Value)
}

// CancelScheduled allows for removal of messages that have been handed to the Service Bus broker for later delivery,
// but have not yet ben enqueued.
func (mc *mgmtClient) CancelScheduled(ctx context.Context, seq ...int64) error {
	ctx, span := spans.StartConsumerSpanFromContext(ctx, string(spanNameCancelScheduledMessage))
	defer span.End()

	msg := &amqp.Message{
		ApplicationProperties: map[string]interface{}{
			"operation": "com.microsoft:cancel-scheduled-message",
		},
		Value: map[string]interface{}{
			"sequence-numbers": seq,
		},
	}

	if deadline, ok := ctx.Deadline(); ok {
		msg.ApplicationProperties["com.microsoft:server-timeout"] = uint(time.Until(deadline) / time.Millisecond)
	}

	resp, err := mc.doRPCWithRetry(ctx, msg, 5, 5*time.Second)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	if resp.Code != 200 {
		return internal.ErrAMQP(*resp)
	}

	return nil
}

func (mc *mgmtClient) startSpan(ctx context.Context, operationName rpcSpanName) (context.Context, tab.Spanner) {
	ctx, span := spans.StartConsumerSpanFromContext(ctx, string(operationName))
	span.AddAttributes(tab.StringAttribute("message_bus.destination", mc.managementPath))
	return ctx, span
}

func (mc *mgmtClient) startSpanFromContext(ctx context.Context, operationName string) (context.Context, tab.Spanner) {
	ctx, span := spans.StartConsumerSpanFromContext(ctx, operationName)
	span.AddAttributes(tab.StringAttribute("message_bus.destination", mc.managementPath))
	return ctx, span
}

func (mc *mgmtClient) startProducerSpanFromContext(ctx context.Context, operationName string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	spans.ApplyComponentInfo(span)
	span.AddAttributes(
		tab.StringAttribute("span.kind", "producer"),
		tab.StringAttribute("message_bus.destination", mc.managementPath),
	)
	return ctx, span
}
