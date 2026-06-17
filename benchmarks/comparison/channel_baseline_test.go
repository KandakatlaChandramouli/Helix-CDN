package comparison

import (
	"sync/atomic"
	"testing"
)

var channelSink atomic.Uint64

func BenchmarkChannelBaseline(
	b *testing.B,
) {
	ch := make(chan int, 1024)

	done := make(chan struct{})

	go func() {
		for range ch {
			channelSink.Add(1)
		}

		close(done)
	}()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ch <- i
	}

	close(ch)

	<-done
}
