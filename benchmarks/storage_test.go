package benchmarks

import (
	"os"
	"testing"

	"helixcdn/internal/storage/mmap"
)

func BenchmarkMmapAppend(b *testing.B) {
	file := "bench.db"

	engine, err := mmap.Open(
		file,
		64*1024*1024,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer os.Remove(file)

	defer engine.Close()

	payload := make([]byte, 1024)

	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := engine.Append(payload)

		if err != nil {
			b.Fatal(err)
		}
	}
}
