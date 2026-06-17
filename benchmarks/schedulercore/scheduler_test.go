package schedulercore

import (
	"testing"

	"helixcdn/internal/runtime/tasks"
	"helixcdn/internal/runtime/workscheduler"
)

var schedulerSink interface{}

func BenchmarkSchedulerRun(
	b *testing.B,
) {
	s := workscheduler.New(4)

	for i := 0; i < b.N; i++ {
		s.Submit(
			i%4,
			tasks.New(
				uint64(i),
				func() {
					schedulerSink = i
				},
			),
		)

		s.RunOnce()
	}
}
