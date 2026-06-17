package latency

import (
	"testing"
	"time"
)

func BenchmarkLatency(
	b *testing.B,
) {

	start := time.Now()

	for i := 0; i < b.N; i++ {
		_ = i * i
	}

	_ = time.Since(start)
}
