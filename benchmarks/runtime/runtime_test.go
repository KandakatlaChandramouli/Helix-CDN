package runtime

import "testing"

func BenchmarkRuntime(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
