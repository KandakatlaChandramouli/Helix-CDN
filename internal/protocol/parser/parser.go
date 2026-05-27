package parser

import (
	"helixcdn/internal/protocol/decoder"
	"helixcdn/internal/protocol/packet"
)

func Parse(
	buf []byte,
) (
	packet.Packet,
	error,
) {

	return decoder.Decode(
		buf,
	)
}
