package runtime

import (
	"testing"

	"helixcdn/internal/runtime/mempool"
)

var memSink []byte

func BenchmarkMemoryPool(
	b *testing.B,
) {
	pool := mempool.New(4096)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v := pool.Get()
		memSink = v
		pool.Put(v)
	}
}
