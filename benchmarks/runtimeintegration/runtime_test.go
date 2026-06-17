package runtimeintegration

import (
	"sync/atomic"
	"testing"

	"helixcdn/internal/runtime/adaptivescheduler"
	"helixcdn/internal/runtime/tasks"
)

var runtimeSink atomic.Uint64

func BenchmarkRuntimeIntegration(
	b *testing.B,
) {
	scheduler := adaptivescheduler.New(8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		scheduler.Submit(
			tasks.New(
				uint64(i),
				func() {
					runtimeSink.Add(1)
				},
			),
		)

		scheduler.RunOnce()
	}
}
