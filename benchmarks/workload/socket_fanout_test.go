package workload

import (
	"sync/atomic"
	"testing"

	"helixcdn/internal/runtime/workqueue"
)

var fanoutSink atomic.Uint64

func BenchmarkSocketFanout(
	b *testing.B,
) {
	q := workqueue.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fanoutSink.Add(1)

		_ = q
	}
}
