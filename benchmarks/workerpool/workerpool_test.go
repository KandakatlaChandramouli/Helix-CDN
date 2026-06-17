package workerpool

import "testing"

func BenchmarkWorkerPool(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
