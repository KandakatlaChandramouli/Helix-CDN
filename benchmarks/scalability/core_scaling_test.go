package scalability

import (
	"fmt"
	"sync/atomic"
	"testing"

	"helixcdn/internal/runtime/adaptivescheduler"
	"helixcdn/internal/runtime/tasks"
)

var scalingSink atomic.Uint64

func benchmarkScaling(
	b *testing.B,
	workers uint64,
) {
	scheduler := adaptivescheduler.New(workers)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		scheduler.Submit(
			tasks.New(
				uint64(i),
				func() {
					scalingSink.Add(1)
				},
			),
		)

		if i%32 == 0 {
			scheduler.RunOnce()
		}
	}

	scheduler.RunOnce()
}

func BenchmarkScaling1(
	b *testing.B,
) {
	benchmarkScaling(b, 1)
}

func BenchmarkScaling2(
	b *testing.B,
) {
	benchmarkScaling(b, 2)
}

func BenchmarkScaling4(
	b *testing.B,
) {
	benchmarkScaling(b, 4)
}

func BenchmarkScaling8(
	b *testing.B,
) {
	benchmarkScaling(b, 8)
}

func BenchmarkScaling16(
	b *testing.B,
) {
	benchmarkScaling(b, 16)
}

func ExampleScaling() {
	fmt.Println("runtime scaling benchmark")
}
