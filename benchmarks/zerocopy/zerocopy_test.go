package zerocopy

import (
	"testing"

	"helixcdn/internal/runtime/reactorzerocopy"
)

func BenchmarkZeroCopy(
	b *testing.B,
) {

	payload := make(
		[]byte,
		4096,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = reactorzerocopy.Slice(
			payload,
		)
	}
}
