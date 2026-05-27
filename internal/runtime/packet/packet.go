package packet

import "sync"

type Packet struct {
	Buf []byte
}

var Pool = sync.Pool{
	New: func() any {
		return &Packet{
			Buf: make([]byte, 4096),
		}
	},
}

func Get() *Packet {
	return Pool.Get().(*Packet)
}

func Put(
	p *Packet,
) {
	Pool.Put(p)
}
