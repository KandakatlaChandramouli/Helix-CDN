//go:build darwin

package epollcore

type Epoll struct {
	fd int
}

func New() (*Epoll, error) {
	return &Epoll{fd: 1}, nil
}

func (e *Epoll) Add(fd int, events uint32) error {
	return nil
}

func (e *Epoll) Wait(events interface{}, timeout int) (int, error) {
	return 0, nil
}

func (e *Epoll) Close() error {
	return nil
}
