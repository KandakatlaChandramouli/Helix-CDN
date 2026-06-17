package corequeue

import (
	"sync"
	"testing"

	"helixcdn/internal/core/mpsc"
)

func BenchmarkMPSCQueue(
	b *testing.B,
) {
	q := mpsc.New[uint64]()

	b.ResetTimer()

	b.RunParallel(func(
		pb *testing.PB,
	) {
		for pb.Next() {
			q.Enqueue(1)
		}
	})
}

func BenchmarkMPSCContention(
	b *testing.B,
) {
	q := mpsc.New[uint64]()

	var wg sync.WaitGroup

	b.ResetTimer()

	for i := 0; i < 8; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < b.N; j++ {
				q.Enqueue(uint64(j))
			}
		}()
	}

	wg.Wait()
}
