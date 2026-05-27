package dispatch

import (
	"testing"

	"helixcdn/internal/runtime/reactordispatch"
)

func BenchmarkDispatch(
	b *testing.B,
) {

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = reactordispatch.Dispatch(
			128,
			uint64(i),
		)
	}
}
