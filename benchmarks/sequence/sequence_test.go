package sequence

import (
	"testing"

	"helixcdn/internal/runtime/reactorsequence"
)

func BenchmarkSequence(
	b *testing.B,
) {

	var seq reactorsequence.Sequence

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = seq.Next()
	}
}
