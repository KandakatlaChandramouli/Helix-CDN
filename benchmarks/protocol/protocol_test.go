package protocol

import (
	"testing"

	"helixcdn/internal/protocol/encoder"
	"helixcdn/internal/protocol/packet"
)

func BenchmarkProtocolEncode(
	b *testing.B,
) {

	p := packet.Packet{
		Version: 1,
		Opcode: 2,
		Payload: []byte("ipl-live-score"),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = encoder.Encode(p)
	}
}
