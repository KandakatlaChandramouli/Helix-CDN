package benchmarks

import (
	"testing"

	"helixcdn/internal/runtime/queue"
)

func BenchmarkMPSCQueue(
	b *testing.B,
) {
	q := queue.New()

	payload := make([]byte, 1024)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		q.Push(payload)

		q.Pop()
	}
}
