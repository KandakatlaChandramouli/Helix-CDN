package deserializer

import (
	"helixcdn/internal/protocol/decoder"
	"helixcdn/internal/protocol/packet"
)

func Deserialize(
	buf []byte,
) (
	packet.Packet,
	error,
) {

	return decoder.Decode(
		buf,
	)
}
