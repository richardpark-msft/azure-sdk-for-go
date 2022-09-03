// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs

import (
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/go-amqp"
)

// EventData is an event that can be sent, using the ProducerClient, to an Event Hub.
type EventData struct {
	// ApplicationProperties can be used to store custom metadata for a message.
	Properties map[string]any

	// Body is the payload for a message.
	Body []byte

	// ContentType describes the payload of the message, with a descriptor following
	// the format of Content-Type, specified by RFC2045 (ex: "application/json").
	ContentType *string

	// CorrelationID is a client-specific id that can be used to mark or identify messages
	// between clients.
	// The type of CorrelationID can be a uint64, UUID, []byte, or string
	CorrelationID any

	// MessageID is an application-defined value that uniquely identifies
	// the message and its payload. The identifier is a free-form string.
	//
	// If enabled, the duplicate detection feature identifies and removes further submissions
	// of messages with the same MessageId.
	MessageID *string
}

// ReceivedEventData is an event that has been received using the ConsumerClient.
type ReceivedEventData struct {
	EventData

	// EnqueuedTime is the UTC time when the message was accepted and stored by Event Hubs.
	EnqueuedTime *time.Time

	// Offset is the offset of the event.
	Offset *int64

	// PartitionKey is used with a partitioned entity and enables assigning related messages
	// to the same internal partition. This ensures that the submission sequence order is correctly
	// recorded. The partition is chosen by a hash function in Event Hubs and cannot be chosen
	// directly.
	PartitionKey *string

	// SequenceNumber is a unique number assigned to a message by Event Hubs.
	SequenceNumber int64

	// Properties set by the Event Hubs service.
	SystemProperties map[string]any

	// RawAMQPMessage is the AMQP message, as received by the client. This can be useful to get access
	// to properties that are not exposed by ReceivedMessage such as payloads encoded into the
	// Value or Sequence section, payloads sent as multiple Data sections, as well as Footer
	// and Header fields.
	RawAMQPMessage *AMQPAnnotatedMessage
}

// Event Hubs custom properties
const (
	// Annotation properties
	partitionKeyAnnotation   = "x-opt-partition-key"
	partitionIDAnnotation    = "x-opt-partition-id"
	sequenceNumberAnnotation = "x-opt-sequence-number"
	offsetNumberAnnotation   = "x-opt-offset"
	enqueuedTimeAnnotation   = "x-opt-enqueued-time"
)

func (e *EventData) toAMQPMessage() *amqp.Message {
	amqpMsg := amqp.NewMessage(e.Body)

	var messageID any

	if e.MessageID != nil {
		messageID = *e.MessageID
	}

	amqpMsg.Properties = &amqp.MessageProperties{
		MessageID: messageID,
	}

	amqpMsg.Properties.ContentType = e.ContentType
	amqpMsg.Properties.CorrelationID = e.CorrelationID

	if len(e.Properties) > 0 {
		amqpMsg.ApplicationProperties = make(map[string]any)
		for key, value := range e.Properties {
			amqpMsg.ApplicationProperties[key] = value
		}
	}

	return amqpMsg
}

// newReceivedEventData creates a received message from an AMQP message.
// NOTE: this converter assumes that the Body of this message will be the first
// serialized byte array in the Data section of the messsage.
func newReceivedEventData(tmp *amqp.Message) *ReceivedEventData {
	amqpMsg := newAMQPAnnotatedMessage(tmp)
	re := &ReceivedEventData{
		RawAMQPMessage: amqpMsg,
	}

	if len(amqpMsg.Body.Data) == 1 {
		re.Body = amqpMsg.Body.Data[0]
	}

	if amqpMsg.Properties != nil {
		if id, ok := amqpMsg.Properties.MessageID.(string); ok {
			re.MessageID = &id
		}

		re.ContentType = amqpMsg.Properties.ContentType
		re.CorrelationID = amqpMsg.Properties.CorrelationID
	}

	if amqpMsg.ApplicationProperties != nil {
		re.Properties = make(map[string]any, len(amqpMsg.ApplicationProperties))
		for key, value := range amqpMsg.ApplicationProperties {
			re.Properties[key] = value
		}
	}

	updateFromAMQPAnnotations(amqpMsg, re)
	return re
}

// the "SystemProperties" in an EventData are any annotations that are
// NOT available at the top level as normal fields. So excluing  sequence
// number, offset, enqueued time, and  partition key.
func updateFromAMQPAnnotations(src *AMQPAnnotatedMessage, dest *ReceivedEventData) {
	if src.MessageAnnotations == nil {
		return
	}

	for kAny, v := range src.MessageAnnotations {
		keyStr, keyIsString := kAny.(string)

		if !keyIsString {
			continue
		}

		switch keyStr {
		case sequenceNumberAnnotation:
			if asInt64, ok := v.(int64); ok {
				dest.SequenceNumber = asInt64
			}
		case partitionKeyAnnotation:
			if asString, ok := v.(string); ok {
				dest.PartitionKey = to.Ptr(asString)
			}
		case enqueuedTimeAnnotation:
			if asTime, ok := v.(time.Time); ok {
				dest.EnqueuedTime = &asTime
			}
		case offsetNumberAnnotation:
			if offsetStr, ok := v.(string); ok {
				if offset, err := strconv.ParseInt(offsetStr, 10, 64); err == nil {
					dest.Offset = &offset
				}
			}
		default:

			if dest.SystemProperties == nil {
				dest.SystemProperties = map[string]any{}
			}

			dest.SystemProperties[keyStr] = v
		}
	}
}
