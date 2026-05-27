package packet

import "sync"

var pool = sync.Pool{
	New: func() any {
		return &Packet{
			Payload: make([]byte, 0, 4096),
		}
	},
}

func Get() *Packet {
	return pool.Get().(*Packet)
}

func Put(p *Packet) {

	p.Opcode = 0
	p.Payload = p.Payload[:0]

	pool.Put(p)
}
