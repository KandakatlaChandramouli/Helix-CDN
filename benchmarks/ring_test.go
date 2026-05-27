package benchmarks

import (
	"testing"

	"helixcdn/internal/edge/ring"
)

func BenchmarkRingPush(
	b *testing.B,
) {
	r := ring.New(1024)

	payload := make([]byte, 1024)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r.Push(payload)
	}
}
