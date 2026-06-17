package serialize

import (
	"testing"

	"helixcdn/internal/protocol/packet"
	"helixcdn/internal/protocol/serializer"
)

func BenchmarkSerialize(
	b *testing.B,
) {

	p := packet.Packet{
		Version: 1,
		Opcode:  2,
		Payload: []byte("helix"),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = serializer.Serialize(
			p,
		)
	}
}
