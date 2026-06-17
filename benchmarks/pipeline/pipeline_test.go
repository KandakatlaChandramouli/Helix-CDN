package pipeline

import "testing"

func BenchmarkPipeline(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
