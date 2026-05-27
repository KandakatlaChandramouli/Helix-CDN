package quorum

import (
	"testing"

	"helixcdn/internal/distributed/quorum"
)

func BenchmarkMajority(
	b *testing.B,
) {

	for i := 0; i < b.N; i++ {

		_ = quorum.Majority(
			5,
		)
	}
}
