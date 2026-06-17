package reactor

import (
	"runtime"
	"sync"
	"testing"
)

var reactorSink int

func BenchmarkReactorContention(
	b *testing.B,
) {
	var wg sync.WaitGroup

	counter := 0

	for i := 0; i < 8; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < b.N; j++ {
				counter++
				reactorSink = counter
				runtime.Gosched()
			}
		}()
	}

	wg.Wait()
}
