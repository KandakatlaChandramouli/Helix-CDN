package parser

import (
	"testing"

	"helixcdn/internal/protocol/encoder"
	"helixcdn/internal/protocol/packet"
	"helixcdn/internal/protocol/parser"
)

func BenchmarkParser(
	b *testing.B,
) {

	p := packet.Packet{
		Version: 1,
		Opcode:  2,
		Payload: []byte("helix"),
	}

	buf := encoder.Encode(p)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, _ = parser.Parse(
			buf,
		)
	}
}
