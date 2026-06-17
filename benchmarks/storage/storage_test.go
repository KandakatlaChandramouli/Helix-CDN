package storage

import "testing"

func BenchmarkStorage(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
