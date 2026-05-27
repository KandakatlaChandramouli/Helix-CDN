package pool

import (
	"testing"

	"helixcdn/internal/runtime/reactorpool"
)

func BenchmarkPool(
	b *testing.B,
) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		buf := reactorpool.Buffers.Get().([]byte)

		reactorpool.Buffers.Put(buf)
	}
}
