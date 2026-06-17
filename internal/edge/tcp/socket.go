package tcp

import "net"

type Socket struct {
	Conn net.Conn
	ID   uint64
}
