package benchmarks

import (
	"testing"

	"helixcdn/internal/edge/backpressure"
)

func BenchmarkTokenBucket(
	b *testing.B,
) {
	tb := backpressure.New(
		1000000,
		1000000,
	)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		tb.Allow()
	}
}
