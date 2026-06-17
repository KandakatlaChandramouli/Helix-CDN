package ringbuffer

import (
	"testing"

	"helixcdn/internal/core/ringbuffer"
)

func BenchmarkRingBufferPush(
	b *testing.B,
) {
	rb := ringbuffer.New(1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rb.Push(uint64(i))
	}
}

func BenchmarkRingBufferPop(
	b *testing.B,
) {
	rb := ringbuffer.New(1024)

	for i := 0; i < 1024; i++ {
		rb.Push(uint64(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rb.Pop()
	}
}
