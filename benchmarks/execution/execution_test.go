package execution

import "testing"

func BenchmarkExecution(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
