// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package models

import (
	"fmt"
	"time"

	"github.com/Azure/go-amqp"
)

type (
	// ReceivedMessage is a received message from a Client.NewReceiver() or Client.NewProcessor().
	ReceivedMessage struct {
		Message

		LockToken              *amqp.UUID
		DeliveryCount          uint32
		LockedUntil            *time.Time // `mapstructure:"x-opt-locked-until"`
		SequenceNumber         *int64     // :"x-opt-sequence-number"`
		EnqueuedSequenceNumber *int64     // :"x-opt-enqueue-sequence-number"`
		EnqueuedTime           *time.Time // :"x-opt-enqueued-time"`
		DeadLetterSource       *string    // :"x-opt-deadletter-source"`

		// available in the raw AMQP message, but not exported by default
		// GroupSequence  *uint32

		// these are internal attributes and will no longer be exported
		// after we merge jhendrix's refactor to move the settlement
		// methods from the message to the receiver.
		RawAMQPMessage *amqp.Message

		// the actual AMQP link ID this message was received on.
		LinkName string

		// the 'revision' of the link we received on (ie, if it was recovered it won't match)
		LinkRevision uint64
	}

	// Message is a SendableMessage which can be sent using a Client.NewSender().
	Message struct {
		ID string

		ContentType   string
		CorrelationID string
		// Body corresponds to the first []byte array in the Data section of an AMQP message.
		Body             []byte
		SessionID        *string
		Subject          string // used to be Label
		ReplyTo          string
		ReplyToSessionID string // used to be ReplyToGroupID
		To               string
		TimeToLive       *time.Duration

		PartitionKey            *string
		TransactionPartitionKey *string
		ScheduledEnqueueTime    *time.Time // `mapstructure:"x-opt-scheduled-enqueue-time"`

		ApplicationProperties map[string]interface{} // was UserProperties
		Format                uint32
	}

	// SystemProperties are used to store properties that are set by the system.
	SystemProperties struct {

		// PartitionID            *int16                 `mapstructure:"x-opt-partition-id"`

		Annotations map[string]interface{} `mapstructure:"-"`
	}
)

const (
	lockTokenName = "x-opt-lock-token"
)

// Set implements tab.Carrier
func (m *Message) Set(key string, value interface{}) {
	if m.ApplicationProperties == nil {
		m.ApplicationProperties = make(map[string]interface{})
	}
	m.ApplicationProperties[key] = value
}

// GetKeyValues implements tab.Carrier
func (m *Message) GetKeyValues() map[string]interface{} {
	return m.ApplicationProperties
}

func (m *Message) MessageType() string {
	return "Message"
}

func (m *Message) ToAMQPMessage() (*amqp.Message, error) {
	amqpMsg := amqp.NewMessage(m.Body)

	if m.TimeToLive != nil {
		if amqpMsg.Header == nil {
			amqpMsg.Header = new(amqp.MessageHeader)
		}
		amqpMsg.Header.TTL = *m.TimeToLive
	}

	amqpMsg.Properties = &amqp.MessageProperties{
		MessageID: m.ID,
	}

	if m.SessionID != nil {
		amqpMsg.Properties.GroupID = *m.SessionID
	}

	// if m.GroupSequence != nil {
	// 	amqpMsg.Properties.GroupSequence = *m.GroupSequence
	// }

	amqpMsg.Properties.CorrelationID = m.CorrelationID
	amqpMsg.Properties.ContentType = m.ContentType
	amqpMsg.Properties.Subject = m.Subject
	amqpMsg.Properties.To = m.To
	amqpMsg.Properties.ReplyTo = m.ReplyTo
	amqpMsg.Properties.ReplyToGroupID = m.ReplyToSessionID

	if len(m.ApplicationProperties) > 0 {
		amqpMsg.ApplicationProperties = make(map[string]interface{})
		for key, value := range m.ApplicationProperties {
			amqpMsg.ApplicationProperties[key] = value
		}
	}

	// These are 'received' message properties
	// if m.SystemProperties != nil {
	// 	// Set the raw annotations first (they may be nil) and add the explicit
	// 	// system properties second to ensure they're set properly.
	// 	amqpMsg.Annotations = addMapToAnnotations(amqpMsg.Annotations, m.SystemProperties.Annotations)

	// 	sysPropMap, err := encodeStructureToMap(m.SystemProperties)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	amqpMsg.Annotations = addMapToAnnotations(amqpMsg.Annotations, sysPropMap)
	// }

	// if m.LockToken != nil {
	// 	if amqpMsg.DeliveryAnnotations == nil {
	// 		amqpMsg.DeliveryAnnotations = make(amqp.Annotations)
	// 	}
	// 	amqpMsg.DeliveryAnnotations[lockTokenName] = *m.LockToken
	// }

	return amqpMsg, nil
}

func addMapToAnnotations(a amqp.Annotations, m map[string]interface{}) amqp.Annotations {
	if a == nil && len(m) > 0 {
		a = make(amqp.Annotations)
	}

	for key, val := range m {
		a[key] = val
	}
	return a
}

func MessageFromAMQPMessage(msg *amqp.Message) (*ReceivedMessage, error) {
	return newReceivedMessage(msg.GetData(), msg)
}

func newReceivedMessage(body []byte, amqpMsg *amqp.Message) (*ReceivedMessage, error) {
	msg := &ReceivedMessage{
		Message: Message{
			Body: body,
		},
		RawAMQPMessage: amqpMsg,
	}

	if amqpMsg.Properties != nil {
		if id, ok := amqpMsg.Properties.MessageID.(string); ok {
			msg.ID = id
		}
		msg.SessionID = &amqpMsg.Properties.GroupID
		//msg.GroupSequence = &amqpMsg.Properties.GroupSequence

		if id, ok := amqpMsg.Properties.CorrelationID.(string); ok {
			msg.CorrelationID = id
		}
		msg.ContentType = amqpMsg.Properties.ContentType
		msg.Subject = amqpMsg.Properties.Subject
		msg.To = amqpMsg.Properties.To
		msg.ReplyTo = amqpMsg.Properties.ReplyTo
		msg.ReplyToSessionID = amqpMsg.Properties.ReplyToGroupID
		if amqpMsg.Header != nil {
			msg.DeliveryCount = amqpMsg.Header.DeliveryCount + 1
			msg.TimeToLive = &amqpMsg.Header.TTL
		}
	}

	if amqpMsg.ApplicationProperties != nil {
		msg.ApplicationProperties = make(map[string]interface{}, len(amqpMsg.ApplicationProperties))
		for key, value := range amqpMsg.ApplicationProperties {
			msg.ApplicationProperties[key] = value
		}
	}

	if amqpMsg.Annotations != nil {
		// grab the standard annotations

		/*
			LockedUntil            *time.Time             `mapstructure:"x-opt-locked-until"`
			SequenceNumber         *int64                 `mapstructure:"x-opt-sequence-number"`
			PartitionID            *int16                 `mapstructure:"x-opt-partition-id"`
			PartitionKey           *string                `mapstructure:"x-opt-partition-key"`
			EnqueuedTime           *time.Time             `mapstructure:"x-opt-enqueued-time"`
			DeadLetterSource       *string                `mapstructure:"x-opt-deadletter-source"`
			ScheduledEnqueueTime   *time.Time             `mapstructure:"x-opt-scheduled-enqueue-time"`
			EnqueuedSequenceNumber *int64                 `mapstructure:"x-opt-enqueue-sequence-number"`
			ViaPartitionKey        *string                `mapstructure:"x-opt-via-partition-key"`
		*/

		if lockedUntil, ok := amqpMsg.Annotations["x-opt-locked-until"]; ok {
			msg.LockedUntil = lockedUntil.(*time.Time)
		}

		if sequenceNumber, ok := amqpMsg.Annotations["x-opt-sequence-number"]; ok {
			msg.SequenceNumber = sequenceNumber.(*int64)
		}

		if partitionKey, ok := amqpMsg.Annotations["x-opt-partition-key"]; ok {
			msg.PartitionKey = partitionKey.(*string)
		}

		if enqueuedTime, ok := amqpMsg.Annotations["x-opt-enqueued-time"]; ok {
			msg.EnqueuedTime = enqueuedTime.(*time.Time)
		}

		if deadLetterSource, ok := amqpMsg.Annotations["x-opt-deadletter-source"]; ok {
			msg.DeadLetterSource = deadLetterSource.(*string)
		}

		if scheduledEnqueueTime, ok := amqpMsg.Annotations["x-opt-scheduled-enqueue-time"]; ok {
			msg.ScheduledEnqueueTime = scheduledEnqueueTime.(*time.Time)
		}

		if enqueuedSequenceNumber, ok := amqpMsg.Annotations["x-opt-enqueue-sequence-number"]; ok {
			msg.EnqueuedSequenceNumber = enqueuedSequenceNumber.(*int64)
		}

		if viaPartitionKey, ok := amqpMsg.Annotations["x-opt-via-partition-key"]; ok {
			msg.TransactionPartitionKey = viaPartitionKey.(*string)
		}

		// TODO: annotation propagation is a thing.

		// If we didn't populate any system properties, set up the struct so we
		// can put the annotations in it
		// if msg.SystemProperties == nil {
		// 	msg.SystemProperties = new(SystemProperties)
		// }

		// Take all string-keyed annotations because the protocol reserves all
		// numeric keys for itself and there are no numeric keys defined in the
		// protocol today:
		//
		//	http://www.amqp.org/sites/amqp.org/files/amqp.pdf (section 3.2.10)
		//
		// This approach is also consistent with the behavior of .NET:
		//
		//	https://docs.microsoft.com/en-us/dotnet/api/azure.messaging.eventhubs.eventdata.systemproperties?view=azure-dotnet#Azure_Messaging_EventHubs_EventData_SystemProperties
		// msg.SystemProperties.Annotations = make(map[string]interface{})
		// for key, val := range amqpMsg.Annotations {
		// 	if s, ok := key.(string); ok {
		// 		msg.SystemProperties.Annotations[s] = val
		// 	}
		// }
	}

	if amqpMsg.DeliveryTag != nil && len(amqpMsg.DeliveryTag) > 0 {
		lockToken, err := lockTokenFromMessageTag(amqpMsg)
		if err != nil {
			return msg, err
		}
		msg.LockToken = lockToken
	}

	if token, ok := amqpMsg.DeliveryAnnotations[lockTokenName]; ok {
		if id, ok := token.(amqp.UUID); ok {
			sid := amqp.UUID([16]byte(id))
			msg.LockToken = &sid
		}
	}

	msg.Format = amqpMsg.Format
	return msg, nil
}

func lockTokenFromMessageTag(msg *amqp.Message) (*amqp.UUID, error) {
	return uuidFromLockTokenBytes(msg.DeliveryTag)
}

func uuidFromLockTokenBytes(bytes []byte) (*amqp.UUID, error) {
	if len(bytes) != 16 {
		return nil, fmt.Errorf("invalid lock token, token was not 16 bytes long")
	}

	var swapIndex = func(indexOne, indexTwo int, array *[16]byte) {
		v1 := array[indexOne]
		array[indexOne] = array[indexTwo]
		array[indexTwo] = v1
	}

	// Get lock token from the deliveryTag
	var lockTokenBytes [16]byte
	copy(lockTokenBytes[:], bytes[:16])
	// translate from .net guid byte serialisation format to amqp rfc standard
	swapIndex(0, 3, &lockTokenBytes)
	swapIndex(1, 2, &lockTokenBytes)
	swapIndex(4, 5, &lockTokenBytes)
	swapIndex(6, 7, &lockTokenBytes)
	amqpUUID := amqp.UUID(lockTokenBytes)

	return &amqpUUID, nil
}

// func encodeStructureToMap(structPointer interface{}) (map[string]interface{}, error) {
// 	valueOfStruct := reflect.ValueOf(structPointer)
// 	s := valueOfStruct.Elem()
// 	if s.Kind() != reflect.Struct {
// 		return nil, fmt.Errorf("must provide a struct")
// 	}

// 	encoded := make(map[string]interface{})
// 	for i := 0; i < s.NumField(); i++ {
// 		f := s.Field(i)
// 		if f.IsValid() && f.CanSet() {
// 			tf := s.Type().Field(i)
// 			tag, err := parseMapStructureTag(tf.Tag)
// 			if err != nil {
// 				return nil, err
// 			}

// 			// Skip any entries with an exclude tag
// 			if tag.Name == "-" {
// 				continue
// 			}

// 			if tag != nil {
// 				switch f.Kind() {
// 				case reflect.Ptr:
// 					if !f.IsNil() || tag.PersistEmpty {
// 						if f.IsNil() {
// 							encoded[tag.Name] = nil
// 						} else {
// 							encoded[tag.Name] = f.Elem().Interface()
// 						}
// 					}
// 				default:
// 					if f.Interface() != reflect.Zero(f.Type()).Interface() || tag.PersistEmpty {
// 						encoded[tag.Name] = f.Interface()
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return encoded, nil
// }

// func parseMapStructureTag(tag reflect.StructTag) (*mapStructureTag, error) {
// 	str, ok := tag.Lookup("mapstructure")
// 	if !ok {
// 		return nil, nil
// 	}

// 	mapTag := new(mapStructureTag)
// 	split := strings.Split(str, ",")
// 	mapTag.Name = strings.TrimSpace(split[0])

// 	if len(split) > 1 {
// 		for _, tagKey := range split[1:] {
// 			switch tagKey {
// 			case "persistempty":
// 				mapTag.PersistEmpty = true
// 			default:
// 				return nil, fmt.Errorf("key %q is not understood", tagKey)
// 			}
// 		}
// 	}
// 	return mapTag, nil
// }
