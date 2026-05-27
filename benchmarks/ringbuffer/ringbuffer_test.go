package ringbuffer

import (
	"testing"

	core "helixcdn/internal/core/ringbuffer"
)

var sink interface{}

func BenchmarkRingBufferPush(
	b *testing.B,
) {
	rb := core.New(1024)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sink = rb.Push(uint64(i))
	}
}

func BenchmarkRingBufferPop(
	b *testing.B,
) {
	rb := core.New(1024)

	for i := 0; i < 1024; i++ {
		rb.Push(uint64(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v, _ := rb.Pop()
		sink = v
	}
}
