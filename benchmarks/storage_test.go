package benchmarks

import "testing"

func BenchmarkBufferAccess(b *testing.B) {
	buf := make([]byte, 4096)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = buf[0]
	}
}
