package amqpsb

// func isRetryableError(err error) bool {
// 	// are these also returned if we've been temporarily disconnected?
// 	// or is it _only_ if the link/client has been explicitly closed?
// 	if errors.Is(err, amqp.ErrConnClosed) ||
// 		errors.Is(err, amqp.ErrLinkClosed) ||
// 		errors.Is(err, amqp.ErrSessionClosed) {
// 		return false
// 	}

// 	if errors.Is(err, amqp.ErrLinkDetached) {

// 	}
// }
