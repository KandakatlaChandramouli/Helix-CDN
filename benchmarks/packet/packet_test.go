package packet

import (
	"testing"

	ppool "helixcdn/internal/protocol/pool"
)

func BenchmarkPacketPool(
	b *testing.B,
) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		p := ppool.Packets.Get()

		ppool.Packets.Put(p)
	}
}
