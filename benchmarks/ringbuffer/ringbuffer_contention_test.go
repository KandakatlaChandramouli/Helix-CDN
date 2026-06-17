package ringbuffer

import (
	"runtime"
	"sync"
	"testing"

	core "helixcdn/internal/core/ringbuffer"
)

var sink uint64

func BenchmarkRingBufferSPSC(
	b *testing.B,
) {
	rb := core.New(1 << 16)

	var wg sync.WaitGroup

	wg.Add(2)

	b.ResetTimer()

	go func() {
		defer wg.Done()

		for i := 0; i < b.N; i++ {
			for !rb.Push(uint64(i)) {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < b.N; i++ {
			for {
				v, ok := rb.Pop()

				if ok {
					sink = v
					break
				}

				runtime.Gosched()
			}
		}
	}()

	wg.Wait()
}
