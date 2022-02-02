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
	// Message format code.
	//
	// The upper three octets of a message format code identify a particular message
	// format. The lowest octet indicates the version of said message format. Any
	// given version of a format is forwards compatible with all higher versions.
	Format uint32

	// The DeliveryTag can be up to 32 octets of binary data.
	// Note that when mode one is enabled there will be no delivery tag.
	DeliveryTag []byte

	// The header section carries standard delivery details about the transfer
	// of a message through the AMQP network.
	Header *AMQPMessageHeader
	// If the header section is omitted the receiver MUST assume the appropriate
	// default values (or the meaning implied by no value being set) for the
	// fields within the header unless other target or node specific defaults
	// have otherwise been set.

	// The delivery-annotations section is used for delivery-specific non-standard
	// properties at the head of the message. Delivery annotations convey information
	// from the sending peer to the receiving peer.
	DeliveryAnnotations AMQPAnnotations
	// If the recipient does not understand the annotation it cannot be acted upon
	// and its effects (such as any implied propagation) cannot be acted upon.
	// Annotations might be specific to one implementation, or common to multiple
	// implementations. The capabilities negotiated on link attach and on the source
	// and target SHOULD be used to establish which annotations a peer supports. A
	// registry of defined annotations and their meanings is maintained [AMQPDELANN].
	// The symbolic key "rejected" is reserved for the use of communicating error
	// information regarding rejected messages. Any values associated with the
	// "rejected" key MUST be of type error.
	//
	// If the delivery-annotations section is omitted, it is equivalent to a
	// delivery-annotations section containing an empty map of annotations.

	// The message-annotations section is used for properties of the message which
	// are aimed at the infrastructure.
	Annotations AMQPAnnotations
	// The message-annotations section is used for properties of the message which
	// are aimed at the infrastructure and SHOULD be propagated across every
	// delivery step. Message annotations convey information about the message.
	// Intermediaries MUST propagate the annotations unless the annotations are
	// explicitly augmented or modified (e.g., by the use of the modified outcome).
	//
	// The capabilities negotiated on link attach and on the source and target can
	// be used to establish which annotations a peer understands; however, in a
	// network of AMQP intermediaries it might not be possible to know if every
	// intermediary will understand the annotation. Note that for some annotations
	// it might not be necessary for the intermediary to understand their purpose,
	// i.e., they could be used purely as an attribute which can be filtered on.
	//
	// A registry of defined annotations and their meanings is maintained [AMQPMESSANN].
	//
	// If the message-annotations section is omitted, it is equivalent to a
	// message-annotations section containing an empty map of annotations.

	// The properties section is used for a defined set of standard properties of
	// the message.
	Properties *AMQPMessageProperties
	// The properties section is part of the bare message; therefore,
	// if retransmitted by an intermediary, it MUST remain unaltered.

	// The application-properties section is a part of the bare message used for
	// structured application data. Intermediaries can use the data within this
	// structure for the purposes of filtering or routing.
	ApplicationProperties map[string]interface{}
	// The keys of this map are restricted to be of type string (which excludes
	// the possibility of a null key) and the values are restricted to be of
	// simple types only, that is, excluding map, list, and array types.

	// Data payloads.
	Data [][]byte
	// A data section contains opaque binary data.

	// Value is the payload of the value-section for an AMQP message.
	// An amqp-value section contains a single AMQP value.
	Value interface{}

	// The footer section is used for details about the message or delivery which
	// can only be calculated or evaluated once the whole bare message has been
	// constructed or seen (for example message hashes, HMACs, signatures and
	// encryption details).
	Footer AMQPAnnotations

	// m is the original amqp.Message, as given to us by go-amqp. This is only needed
	// when we're passing the message, opaquely, between our API and go-amqp.
	m *goamqp.Message
}

type AMQPAnnotations map[interface{}]interface{}

// AMQPMessageProperties is the defined set of properties for AMQP messages.
type AMQPMessageProperties struct {
	// MessageID, if set, uniquely identifies a message within the message system.
	// The message producer is usually responsible for setting the message-id in
	// such a way that it is assured to be globally unique. A broker MAY discard a
	// message as a duplicate if the value of the message-id matches that of a
	// previously received message sent to the same node.
	MessageID interface{} // uint64, UUID, []byte, or string

	// UserID is the identity of the user responsible for producing the message.
	// The client sets this value, and it MAY be authenticated by intermediaries.
	UserID []byte

	// To identifies the node that is the intended destination of the message.
	// On any given transfer this might not be the node at the receiving end of the link.
	To *string

	// Subject is a common field for summary information about the message content and purpose.
	Subject *string

	// ReplyTo is the address of the node to send replies to.
	ReplyTo *string

	// CorrelationID is the client-specific id that can be used to mark or identify messages
	// between clients.
	CorrelationID interface{} // uint64, UUID, []byte, or string

	// ContentType is the RFC-2046 [RFC2046] MIME type for the message's application-data section
	// (body). As per RFC-2046 [RFC2046] this can contain a charset parameter defining
	// the character encoding used: e.g., 'text/plain; charset="utf-8"'.
	//
	// For clarity, as per section 7.2.1 of RFC-2616 [RFC2616], where the content type
	// is unknown the content-type SHOULD NOT be set. This allows the recipient the
	// opportunity to determine the actual type. Where the section is known to be truly
	// opaque binary data, the content-type SHOULD be set to application/octet-stream.
	//
	// When using an application-data section with a section code other than data,
	// content-type SHOULD NOT be set.
	ContentType *string

	// ContentEncoding is the content-encoding property is used as a modifier to the
	// content-type. When present, its value indicates what additional content encodings have been
	// applied to the application-data, and thus what decoding mechanisms need to be
	// applied in order to obtain the media-type referenced by the content-type header
	// field.
	//
	// Content-encoding is primarily used to allow a document to be compressed without
	// losing the identity of its underlying content type.
	//
	// Content-encodings are to be interpreted as per section 3.5 of RFC 2616 [RFC2616].
	// Valid content-encodings are registered at IANA [IANAHTTPPARAMS].
	//
	// The content-encoding MUST NOT be set when the application-data section is other
	// than data. The binary representation of all other application-data section types
	// is defined completely in terms of the AMQP type system.
	//
	// Implementations MUST NOT use the identity encoding. Instead, implementations
	// SHOULD NOT set this property. Implementations SHOULD NOT use the compress encoding,
	// except as to remain compatible with messages originally sent with other protocols,
	// e.g. HTTP or SMTP.
	//
	// Implementations SHOULD NOT specify multiple content-encoding values except as to
	// be compatible with messages originally sent with other protocols, e.g. HTTP or SMTP.
	ContentEncoding *string

	// AbsoluteExpiryTime is an absolute time when this message is considered to be expired.
	AbsoluteExpiryTime *time.Time

	// CreationTime is an absolute time when this message was created.
	CreationTime *time.Time

	// GroupID identifies the group the message belongs to.
	GroupID *string

	// GroupSequence identifies the relative position of this message within its group.
	GroupSequence *uint32 // RFC-1982 sequence number

	// ReplyToGroupID is the client-specific id that is used so that client can send replies to this
	// message to a specific group.
	ReplyToGroupID *string
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
	return &AMQPMessage{
		Format:                m.Format,
		DeliveryTag:           m.DeliveryTag,
		Header:                (*AMQPMessageHeader)(m.Header),
		DeliveryAnnotations:   AMQPAnnotations(m.DeliveryAnnotations),
		Annotations:           AMQPAnnotations(m.Annotations),
		Properties:            (*AMQPMessageProperties)(m.Properties),
		ApplicationProperties: m.ApplicationProperties,
		Data:                  m.Data,
		Value:                 m.Value,
		Footer:                AMQPAnnotations(m.Footer),
		m:                     m,
	}
}

func (m *AMQPMessage) toGoAMQPMessage() *goamqp.Message {
	return &goamqp.Message{
		Format:                m.Format,
		DeliveryTag:           m.DeliveryTag,
		Header:                (*goamqp.MessageHeader)(m.Header),
		Properties:            (*goamqp.MessageProperties)(m.Properties),
		DeliveryAnnotations:   goamqp.Annotations(m.DeliveryAnnotations),
		Annotations:           goamqp.Annotations(m.Annotations),
		ApplicationProperties: m.ApplicationProperties,
		Data:                  m.Data,
		Value:                 m.Value,
		Footer:                goamqp.Annotations(m.Footer),
	}
}
