package workload

import (
	"sync/atomic"
	"testing"

	"helixcdn/internal/runtime/adaptivescheduler"
	"helixcdn/internal/runtime/tasks"
)

var burstSink atomic.Uint64

func BenchmarkBurstWorkload(
	b *testing.B,
) {
	scheduler := adaptivescheduler.New(8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 128; j++ {
			scheduler.Submit(
				tasks.New(
					uint64(j),
					func() {
						burstSink.Add(1)
					},
				),
			)
		}

		scheduler.RunOnce()
	}
}
