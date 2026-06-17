//go:build darwin

package eventfd

type EventFD struct {
	FD int
}

func New() (*EventFD, error) {
	return &EventFD{FD: 1}, nil
}

func (e *EventFD) Signal() error {
	return nil
}

func (e *EventFD) Close() error {
	return nil
}
