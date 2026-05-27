package benchmarks

import (
	"testing"

	"helixcdn/internal/runtime/packet"
)

func BenchmarkPacketPool(
	b *testing.B,
) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p := packet.Get()

		packet.Put(p)
	}
}
