package fdregistry

import "sync"

type Registry struct {
	mu sync.RWMutex

	fds map[int]string
}

func New() *Registry {
	return &Registry{
		fds: make(map[int]string),
	}
}

func (r *Registry) Add(
	fd int,
	kind string,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.fds[fd] = kind
}

func (r *Registry) Get(
	fd int,
) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	v, ok := r.fds[fd]

	return v, ok
}
