package saturation

import (
	"sync/atomic"
	"testing"

	"helixcdn/internal/runtime/adaptivescheduler"
	"helixcdn/internal/runtime/tasks"
)

var saturationSink atomic.Uint64

func BenchmarkRuntimeSaturation(
	b *testing.B,
) {
	scheduler := adaptivescheduler.New(16)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		scheduler.Submit(
			tasks.New(
				uint64(i),
				func() {
					saturationSink.Add(1)
				},
			),
		)

		if i%64 == 0 {
			scheduler.RunOnce()
		}
	}

	for i := 0; i < 128; i++ {
		scheduler.RunOnce()
	}
}
