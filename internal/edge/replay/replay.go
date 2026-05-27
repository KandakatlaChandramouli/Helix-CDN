package replay

import (
	"sync"

	"helixcdn/internal/edge/ring"
)

type Replay struct {
	mu     sync.RWMutex
	topics map[uint32]*ring.Ring
}

func New() *Replay {
	return &Replay{
		topics: make(
			map[uint32]*ring.Ring,
		),
	}
}

func (r *Replay) Append(
	topic uint32,
	payload []byte,
) {
	r.mu.Lock()

	buffer, ok := r.topics[topic]

	if !ok {
		buffer = ring.New(1024)

		r.topics[topic] = buffer
	}

	r.mu.Unlock()

	buffer.Push(payload)
}
