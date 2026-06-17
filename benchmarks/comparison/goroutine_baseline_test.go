package comparison

import (
	"sync"
	"sync/atomic"
	"testing"
)

var goroutineSink atomic.Uint64

func BenchmarkGoroutineBaseline(
	b *testing.B,
) {
	var wg sync.WaitGroup

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		go func() {
			goroutineSink.Add(1)
			wg.Done()
		}()
	}

	wg.Wait()
}
