package executor

import "testing"

func BenchmarkExecutor(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
