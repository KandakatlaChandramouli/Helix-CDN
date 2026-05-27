package p99latency

import (
	"sort"
)

type Tracker struct {
	values []int64
}

func New() *Tracker {
	return &Tracker{
		values: make([]int64, 0),
	}
}

func (t *Tracker) Observe(
	v int64,
) {
	t.values = append(t.values, v)
}

func (t *Tracker) P99() int64 {
	if len(t.values) == 0 {
		return 0
	}

	sort.Slice(
		t.values,
		func(i, j int) bool {
			return t.values[i] < t.values[j]
		},
	)

	idx := int(float64(len(t.values)) * 0.99)

	if idx >= len(t.values) {
		idx = len(t.values) - 1
	}

	return t.values[idx]
}
