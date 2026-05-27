package reactorregistry

import (
	"sync"

	"github.com/panjf2000/gnet/v2"
)

type Registry struct {
	mu sync.RWMutex

	connections map[string]gnet.Conn
}

func New() *Registry {

	return &Registry{
		connections: make(
			map[string]gnet.Conn,
		),
	}
}

func (r *Registry) Add(
	id string,
	conn gnet.Conn,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.connections[id] = conn
}

func (r *Registry) Remove(
	id string,
) {

	r.mu.Lock()
	defer r.mu.Unlock()

	delete(
		r.connections,
		id,
	)
}

func (r *Registry) Range(
	fn func(string, gnet.Conn),
) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	for id, conn := range r.connections {
		fn(id, conn)
	}
}
