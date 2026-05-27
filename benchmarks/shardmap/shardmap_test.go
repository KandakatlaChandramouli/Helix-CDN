package shardmap

import (
	"testing"

	"helixcdn/internal/runtime/reactorshardmap"
)

func BenchmarkShardMap(
	b *testing.B,
) {

	sm := reactorshardmap.New()

	payload := []byte("ipl-event")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		sm.Append(
			uint64(i % 128),
			payload,
		)
	}
}
