package benchmarks

import (
	"testing"

	"helixcdn/internal/edge/frame"
)

func BenchmarkFrameDecode(
	b *testing.B,
) {
	buf := make([]byte, 1024)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		frame.Decode(buf)
	}
}
