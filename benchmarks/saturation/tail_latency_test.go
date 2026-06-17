package saturation

import (
	"sync/atomic"
	"testing"
	"time"

	"helixcdn/internal/runtime/p99latency"
)

var latencySink atomic.Uint64

func BenchmarkTailLatencyTracking(
	b *testing.B,
) {
	tracker := p99latency.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()

		latencySink.Add(1)

		tracker.Observe(
			time.Since(start).Nanoseconds(),
		)
	}

	_ = tracker.P99()
}
