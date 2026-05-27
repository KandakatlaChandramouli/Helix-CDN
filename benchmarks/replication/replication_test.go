package replication

import "testing"

func BenchmarkReplication(
	b *testing.B,
) {

	payload := []byte(
		"replicated-packet",
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = payload
	}
}
