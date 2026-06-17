package microbatch

import (
	"sync"
	"time"
)

type Batch struct {
	Topic uint32
	Count uint64
}

type Aggregator struct {
	mu      sync.Mutex
	batches map[uint32]*Batch
	ticker  *time.Ticker
}

func New(
	interval time.Duration,
) *Aggregator {
	return &Aggregator{
		batches: make(
			map[uint32]*Batch,
		),
		ticker: time.NewTicker(interval),
	}
}

func (a *Aggregator) Add(
	topic uint32,
) {
	a.mu.Lock()
	defer a.mu.Unlock()

	batch, ok := a.batches[topic]

	if !ok {
		batch = &Batch{
			Topic: topic,
		}

		a.batches[topic] = batch
	}

	batch.Count++
}

func (a *Aggregator) Flush(
	fn func([]*Batch),
) {
	for range a.ticker.C {
		a.mu.Lock()

		var batches []*Batch

		for _, batch := range a.batches {
			batches = append(
				batches,
				batch,
			)
		}

		a.batches = make(
			map[uint32]*Batch,
		)

		a.mu.Unlock()

		fn(batches)
	}
}
