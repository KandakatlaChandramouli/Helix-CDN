package benchmarks

import (
	"testing"

	"helixcdn/internal/edge/session"
	"helixcdn/internal/edge/shard"
)

func BenchmarkShardRegistry(
	b *testing.B,
) {
	registry := shard.New()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		registry.Add(
			session.New(
				uint64(i),
			),
		)
	}
}
