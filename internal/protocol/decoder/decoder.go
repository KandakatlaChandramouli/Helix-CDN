package decoder

import (
	"encoding/binary"
	"errors"

	"helixcdn/internal/protocol/packet"
)

func Decode(
	buf []byte,
) (
	packet.Packet,
	error,
) {

	if len(buf) < 6 {

		return packet.Packet{},
			errors.New(
				"invalid packet",
			)
	}

	length := binary.BigEndian.Uint32(
		buf[2:6],
	)

	if len(buf) < int(6+length) {

		return packet.Packet{},
			errors.New(
				"incomplete packet",
			)
	}

	return packet.Packet{
		Version: buf[0],
		Opcode: buf[1],
		Payload: buf[6 : 6+length],
	}, nil
}
