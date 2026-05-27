package eventfd

import "golang.org/x/sys/unix"

type EventFD struct {
	FD int
}

func New() (*EventFD, error) {
	fd, err := unix.Eventfd(
		0,
		unix.EFD_NONBLOCK,
	)

	if err != nil {
		return nil, err
	}

	return &EventFD{
		FD: fd,
	}, nil
}
