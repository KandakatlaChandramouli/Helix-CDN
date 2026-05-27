package socketruntime

import (
	"net"
)

type RuntimeSocket struct {
	Listener net.Listener
}

func Listen(
	addr string,
) (*RuntimeSocket, error) {
	ln, err := net.Listen(
		"tcp",
		addr,
	)

	if err != nil {
		return nil, err
	}

	return &RuntimeSocket{
		Listener: ln,
	}, nil
}

func (r *RuntimeSocket) Close() error {
	return r.Listener.Close()
}
