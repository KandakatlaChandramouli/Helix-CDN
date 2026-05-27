package p99latency

import "testing"

func TestP99(
	t *testing.T,
) {
	p := New()

	for i := 0; i < 100; i++ {
		p.Observe(int64(i))
	}

	if p.P99() <= 0 {
		t.Fatal("invalid p99")
	}
}
