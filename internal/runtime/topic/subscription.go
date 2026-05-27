package topic

import "sync"

type Registry struct {
	mu sync.RWMutex

	topics map[string]map[uint64]struct{}
}

func New() *Registry {
	return &Registry{
		topics: make(
			map[string]map[uint64]struct{},
		),
	}
}
