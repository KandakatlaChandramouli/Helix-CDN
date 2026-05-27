package saturation

import (
	"sync/atomic"
	"testing"

	"helixcdn/internal/runtime/runtimepressure"
)

var pressureSink atomic.Uint64

func BenchmarkQueuePressure(
	b *testing.B,
) {
	var p runtimepressure.Pressure

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p.Inc()

		pressureSink.Store(
			p.Load(),
		)

		p.Dec()
	}
}
