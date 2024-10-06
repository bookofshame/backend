package sms

type Payload struct {
	Number  string
	Message string
}

type Sms interface {
	// Send The body should be a valid HTML string
	Send(payload Payload) error
}
