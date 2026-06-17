package epollruntime

import (
	"testing"

	"helixcdn/internal/runtime/epollcore"
)

var epollSink interface{}

func BenchmarkEpollCreate(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
		ep, err := epollcore.New()

		if err != nil {
			b.Fatal(err)
		}

		epollSink = ep

		_ = ep.Close()
	}
}
