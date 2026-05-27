package pool

import (
	"sync"

	"helixcdn/internal/protocol/packet"
)

var Packets = sync.Pool{
	New: func() any {

		return &packet.Packet{
			Payload: make(
				[]byte,
				0,
				4096,
			),
		}
	},
}
