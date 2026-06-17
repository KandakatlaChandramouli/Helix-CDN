package coalesce

import (
	"testing"

	"helixcdn/internal/runtime/reactorcoalesce"
)

func BenchmarkCoalesce(
	b *testing.B,
) {

	c := reactorcoalesce.New()

	payload := []byte(
		"live-event",
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Add(payload)
	}
}
