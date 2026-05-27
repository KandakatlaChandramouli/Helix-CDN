package speculative

import "testing"

func BenchmarkSpeculative(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
