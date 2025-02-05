package eh

import (
	"errors"

	"github.com/Azure/go-amqp"
)

// ErrorConditionGeoDR occurs when using an old integer offset against a hub that has
// georeplication enabled, which requires the new stroffset format.
const ErrorConditionGeoDR = amqp.ErrCond("com.microsoft:georeplication:invalid-offset")

func IsInvalidGeoDROffsetError(err error) bool {
	if amqpErr := (*amqp.Error)(nil); errors.As(err, &amqpErr) {
		if amqpErr.Condition == ErrorConditionGeoDR {
			return true
		}
	}
	return false
}
