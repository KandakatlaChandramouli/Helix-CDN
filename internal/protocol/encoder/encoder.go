package encoder

import (
	"encoding/binary"

	"helixcdn/internal/protocol/packet"
)

func Encode(
	p packet.Packet,
) []byte {

	size := 6 + len(p.Payload)

	buf := make(
		[]byte,
		size,
	)

	buf[0] = p.Version
	buf[1] = p.Opcode

	binary.BigEndian.PutUint32(
		buf[2:6],
		uint32(len(p.Payload)),
	)

	copy(
		buf[6:],
		p.Payload,
	)

	return buf
}
