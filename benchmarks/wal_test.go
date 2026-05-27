package benchmarks

import (
	"os"
	"testing"

	"helixcdn/internal/storage/wal"
)

func BenchmarkWalAppend(b *testing.B) {
	file := "bench.wal"

	writer, err := wal.Open(file)

	if err != nil {
		b.Fatal(err)
	}

	defer os.Remove(file)

	defer writer.Close()

	payload := make([]byte, 1024)

	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := writer.Append(payload); err != nil {
			b.Fatal(err)
		}
	}
}
