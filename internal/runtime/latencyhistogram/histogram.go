package latencyhistogram

type Histogram struct {
	Buckets [64]uint64
}

func New() *Histogram {
	return &Histogram{}
}

func (h *Histogram) Observe(
	v int64,
) {
	idx := 0

	for v > 1 && idx < 63 {
		v >>= 1
		idx++
	}

	h.Buckets[idx]++
}
