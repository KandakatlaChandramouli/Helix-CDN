package chaos

import "testing"

func BenchmarkChaos(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
