package ringbuffer

import (
	"testing"

	"helixcdn/internal/runtime/reactorringbuffer"
)

func BenchmarkRingBuffer(
	b *testing.B,
) {

	r := reactorringbuffer.New(
		4096,
	)

	payload := []byte(
		"helix",
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Push(payload)
	}
}
