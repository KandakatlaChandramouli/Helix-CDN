package errors

type ProtocolError struct {
	Message string
}

func (e ProtocolError) Error() string {
	return e.Message
}
