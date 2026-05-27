package benchmarks

import (
	"testing"

	"helixcdn/internal/edge/session"
)

func BenchmarkRegistryAdd(
	b *testing.B,
) {
	registry := session.NewRegistry()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		registry.Add(&session.Client{
			ID: uint64(i),
		})
	}
}
