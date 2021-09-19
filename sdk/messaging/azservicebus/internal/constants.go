package internal

// ReceiveMode represents the lock style to use for a reciever - either
// `PeekLock` or `ReceiveAndDelete`
type ReceiveMode int

const (
	// PeekLock will lock messages as they are received and can be settled
	// using the Receiver or Processor's (Complete|Abandon|DeadLetter|Defer)Message
	// functions.
	PeekLock ReceiveMode = 0
	// ReceiveAndDelete will delete messages as they are received.
	ReceiveAndDelete ReceiveMode = 1
)
