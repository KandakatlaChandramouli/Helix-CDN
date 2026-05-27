package serializer

import (
	"helixcdn/internal/protocol/encoder"
	"helixcdn/internal/protocol/packet"
)

func Serialize(
	p packet.Packet,
) []byte {

	return encoder.Encode(
		p,
	)
}
