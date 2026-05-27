package telemetry

import "testing"

func BenchmarkTelemetry(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
