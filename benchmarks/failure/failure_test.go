package failure

import "testing"

func BenchmarkFailure(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
