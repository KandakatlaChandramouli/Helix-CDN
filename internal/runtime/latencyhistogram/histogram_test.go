package latencyhistogram

import "testing"

func TestHistogram(
	t *testing.T,
) {
	h := New()

	h.Observe(64)

	total := uint64(0)

	for _, v := range h.Buckets {
		total += v
	}

	if total != 1 {
		t.Fatal("histogram observation missing")
	}
}
