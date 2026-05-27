package reactorsubscription

import "sync"

type Registry struct {
	mu sync.RWMutex

	topics map[string]map[string]struct{}
}

func New() *Registry {

	return &Registry{
		topics: make(
			map[string]map[string]struct{},
		),
	}
}

func (r *Registry) Subscribe(
	client string,
	topic string,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.topics[topic]; !ok {

		r.topics[topic] = make(
			map[string]struct{},
		)
	}

	r.topics[topic][client] = struct{}{}
}
