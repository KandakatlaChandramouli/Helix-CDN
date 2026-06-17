package checkpoint

import "testing"

func BenchmarkCheckpoint(
	b *testing.B,
) {

	var lsn uint64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lsn++
	}
}
