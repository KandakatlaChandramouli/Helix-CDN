package memory

import "testing"

func BenchmarkAllocation(
	b *testing.B,
) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		_ = make(
			[]byte,
			4096,
		)
	}
}
