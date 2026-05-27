package epollcore

import (
	"golang.org/x/sys/unix"
)

type Epoll struct {
	fd int
}

func New() (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)

	if err != nil {
		return nil, err
	}

	return &Epoll{
		fd: fd,
	}, nil
}

func (e *Epoll) Add(
	fd int,
	events uint32,
) error {
	ev := &unix.EpollEvent{
		Events: events,
		Fd:     int32(fd),
	}

	return unix.EpollCtl(
		e.fd,
		unix.EPOLL_CTL_ADD,
		fd,
		ev,
	)
}

func (e *Epoll) Wait(
	events []unix.EpollEvent,
	timeout int,
) (int, error) {
	return unix.EpollWait(
		e.fd,
		events,
		timeout,
	)
}
