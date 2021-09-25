package internal

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/Azure/go-amqp"
)

const (
	LockedUntilAnnotation = "x-opt-locked-until"
)

// GetLockExpirationTime gets the Service Bus annotation
// with the lock expiration time (`x-opt-locked-until`).
// Returns the lock expiration time or the zero time instant.
func GetLockExpirationTime(m *amqp.Message) time.Time {
	if lockedUntil, ok := m.Annotations[LockedUntilAnnotation]; ok {
		asTime := lockedUntil.(time.Time)
		return asTime
	}

	return time.Time{}
}

func GetLockToken(msg *amqp.Message) (*amqp.UUID, error) {
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

func GetMessageID(msg *amqp.Message) string {
	if msg == nil || msg.Properties == nil {
		return ""
	}

	switch asType := msg.Properties.MessageID.(type) {
	case amqp.UUID:
		return asType.String()
	case []byte:
		return hex.EncodeToString(asType)
	case string:
		return asType
	case uint64:
		return strconv.FormatUint(asType, 10)
	}

	// unknown type, treat as "undecodeable"
	return ""
}
