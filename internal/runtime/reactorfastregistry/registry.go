package reactorfastregistry

import "sync"

type Registry struct {
	mu sync.RWMutex

	clients map[uint64]struct{}
}

func New() *Registry {

	return &Registry{
		clients: make(
			map[uint64]struct{},
		),
	}
}

func (r *Registry) Add(
	id uint64,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.clients[id] = struct{}{}
}
