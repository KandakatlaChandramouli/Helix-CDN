package reactor

import (
	"testing"

	"helixcdn/internal/runtime/reactorrouter"
)

func BenchmarkRouter(
	b *testing.B,
) {

	for i := 0; i < b.N; i++ {

		_ = reactorrouter.Route(
			"ipl-live-score",
			128,
		)
	}
}
