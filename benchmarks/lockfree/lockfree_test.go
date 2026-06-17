package lockfree

import (
	"testing"

	"helixcdn/internal/runtime/reactorlockfree"
)

func BenchmarkLockFree(
	b *testing.B,
) {

	var c reactorlockfree.Counter

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Inc()
	}
}
