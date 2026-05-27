package stealing

import "testing"

func BenchmarkStealing(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
