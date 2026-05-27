package batcher

import (
	"time"
)

type Batch struct {
	Items [][]byte
}

type Batcher struct {
	maxSize  int
	interval time.Duration
	queue    chan []byte
}

func New(
	maxSize int,
	interval time.Duration,
) *Batcher {
	return &Batcher{
		maxSize:  maxSize,
		interval: interval,
		queue: make(chan []byte, 1<<20),
	}
}

func (b *Batcher) Push(
	payload []byte,
) {
	select {
	case b.queue <- payload:
	default:
	}
}

func (b *Batcher) Run(
	fn func(*Batch),
) {
	ticker := time.NewTicker(
		b.interval,
	)

	batch := &Batch{}

	for {
		select {
		case msg := <-b.queue:
			batch.Items = append(
				batch.Items,
				msg,
			)

			if len(batch.Items) >= b.maxSize {
				fn(batch)

				batch = &Batch{}
			}

		case <-ticker.C:
			if len(batch.Items) > 0 {
				fn(batch)

				batch = &Batch{}
			}
		}
	}
}
