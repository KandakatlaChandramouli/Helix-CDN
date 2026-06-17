package presence

import (
	"sync"
	"time"
)

type Tracker struct {
	mu sync.RWMutex

	clients map[uint64]time.Time
}

func New() *Tracker {
	return &Tracker{
		clients: make(
			map[uint64]time.Time,
		),
	}
}

func (t *Tracker) Heartbeat(
	id uint64,
) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.clients[id] = time.Now()
}

func (t *Tracker) Count() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return len(t.clients)
}
