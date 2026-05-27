package session

import (
	"sync"
	"sync/atomic"
)

type Client struct {
	ID uint64
}

type Registry struct {
	clients sync.Map
	count   atomic.Uint64
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Add(c *Client) {
	r.clients.Store(c.ID, c)
	r.count.Add(1)
}

func (r *Registry) Remove(id uint64) {
	r.clients.Delete(id)
	r.count.Add(^uint64(0))
}

func (r *Registry) Count() uint64 {
	return r.count.Load()
}

func (r *Registry) Range(fn func(*Client)) {
	r.clients.Range(func(_, value any) bool {
		client := value.(*Client)

		fn(client)

		return true
	})
}
