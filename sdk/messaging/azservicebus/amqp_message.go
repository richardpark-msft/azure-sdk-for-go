package azservicebus

import (
	"time"

	goamqp "github.com/Azure/go-amqp"
)

// AMQPMessage is the "raw" AMQP message which gives full control and access to
// all parts of an AMQP message.
// An AMQPMessage will be included in the ReceivedMessage which is retrieved
// from
type AMQPMessage struct {
	// The application-properties section is a part of the bare message used for
	// structured application data.
	ApplicationProperties map[string]interface{}

	// Data corresponds to the 'data' section of the AMQP payload.
	Data [][]byte

	// The delivery-annotations section is used for delivery-specific non-standard
	// properties at the head of the message.
	DeliveryAnnotations AMQPAnnotations

	// The footer section is used for details about the message or delivery.
	Footer AMQPAnnotations

	// The header section carries standard delivery details about the transfer
	// of a message through the AMQP network.
	Header *AMQPMessageHeader

	// The message-annotations section is used for properties of the message which
	// are aimed at the infrastructure.
	MessageAnnotations AMQPAnnotations

	// The properties section is used for a defined set of standard properties of
	// the message.
	Properties *AMQPMessageProperties

	// Value is the payload of the value-section for an AMQP message.
	// An amqp-value section contains a single AMQP value.
	Value interface{}

	// m is the original amqp.Message, as given to us by go-amqp. This is only needed
	// when we're passing the message, opaquely, between our API and go-amqp.
	m *goamqp.Message
}

type AMQPAnnotations map[interface{}]interface{}

// AMQPMessageProperties is the defined set of properties for AMQP messages.
type AMQPMessageProperties struct {
	// AbsoluteExpiryTime is an absolute time when this message is considered to be expired.
	AbsoluteExpiryTime *time.Time

	// ContentEncoding is used as a modifier to the ContentType.
	ContentEncoding *string

	// ContentType is the RFC-2046 [RFC2046] MIME type for the message's application-data section
	// (body). As per RFC-2046 [RFC2046] this can contain a charset parameter defining
	// the character encoding used: e.g., 'text/plain; charset="utf-8"'.
	ContentType *string

	// CorrelationID is the client-specific id that can be used to mark or identify messages
	// between clients.
	CorrelationID interface{} // uint64, UUID, []byte, or string

	// CreationTime is an absolute time when this message was created.
	CreationTime *time.Time

	// GroupID identifies the group the message belongs to.
	GroupID *string

	// GroupSequence identifies the relative position of this message within its group.
	GroupSequence *uint32 // RFC-1982 sequence number

	// MessageID, if set, uniquely identifies a message within the message system.
	MessageID interface{} // uint64, UUID, []byte, or string

	// ReplyTo is the address of the node to send replies to.
	ReplyTo *string

	// ReplyToGroupID is the client-specific id that is used so that client can send replies to this
	// message to a specific group.
	ReplyToGroupID *string

	// Subject is a common field for summary information about the message content and purpose.
	Subject *string

	// To identifies the node that is the intended destination of the message.
	// On any given transfer this might not be the node at the receiving end of the link.
	To *string
}

// AMQPMessageHeader carries standard delivery details about the transfer
// of a message.
type AMQPMessageHeader struct {
	Durable       bool
	Priority      uint8
	TTL           time.Duration // from milliseconds
	FirstAcquirer bool
	DeliveryCount uint32
}

func newAMQPMessage(m *goamqp.Message) *AMQPMessage {
	var amqpProps *AMQPMessageProperties

	if m.Properties != nil {
		amqpProps = &AMQPMessageProperties{
			AbsoluteExpiryTime: m.Properties.AbsoluteExpiryTime,
			ContentEncoding:    m.Properties.ContentEncoding,
			ContentType:        m.Properties.ContentType,
			CorrelationID:      m.Properties.CorrelationID,
			CreationTime:       m.Properties.CreationTime,
			GroupID:            m.Properties.GroupID,
			GroupSequence:      m.Properties.GroupSequence,
			MessageID:          m.Properties.MessageID,
			ReplyTo:            m.Properties.ReplyTo,
			ReplyToGroupID:     m.Properties.ReplyToGroupID,
			Subject:            m.Properties.Subject,
			To:                 m.Properties.To,
		}
	}

	return &AMQPMessage{
		Header:                (*AMQPMessageHeader)(m.Header),
		DeliveryAnnotations:   AMQPAnnotations(m.DeliveryAnnotations),
		MessageAnnotations:    AMQPAnnotations(m.Annotations),
		Properties:            amqpProps,
		ApplicationProperties: m.ApplicationProperties,
		Data:                  m.Data,
		Value:                 m.Value,
		Footer:                AMQPAnnotations(m.Footer),
		m:                     m,
	}
}

func (m *AMQPMessage) toGoAMQPMessage() *goamqp.Message {
	var goprops *goamqp.MessageProperties

	if m.Properties != nil {
		goprops = &goamqp.MessageProperties{
			AbsoluteExpiryTime: m.Properties.AbsoluteExpiryTime,
			ContentEncoding:    m.Properties.ContentEncoding,
			ContentType:        m.Properties.ContentType,
			CorrelationID:      m.Properties.CorrelationID,
			CreationTime:       m.Properties.CreationTime,
			GroupID:            m.Properties.GroupID,
			GroupSequence:      m.Properties.GroupSequence,
			MessageID:          m.Properties.MessageID,
			ReplyTo:            m.Properties.ReplyTo,
			ReplyToGroupID:     m.Properties.ReplyToGroupID,
			Subject:            m.Properties.Subject,
			To:                 m.Properties.To,
		}
	}

	return &goamqp.Message{
		Header:                (*goamqp.MessageHeader)(m.Header),
		Properties:            goprops,
		DeliveryAnnotations:   goamqp.Annotations(m.DeliveryAnnotations),
		Annotations:           goamqp.Annotations(m.MessageAnnotations),
		ApplicationProperties: m.ApplicationProperties,
		Data:                  m.Data,
		Value:                 m.Value,
		Footer:                goamqp.Annotations(m.Footer),
	}
}
