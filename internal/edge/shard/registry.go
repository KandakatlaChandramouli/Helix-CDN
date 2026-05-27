package shard

import (
	"sync"

	"helixcdn/internal/edge/session"
)

const (
	ShardCount = 256
)

type bucket struct {
	sync.RWMutex
	clients map[uint64]*session.Session
}

type Registry struct {
	shards [ShardCount]bucket
}

func New() *Registry {
	r := &Registry{}

	for i := 0; i < ShardCount; i++ {
		r.shards[i].clients = make(
			map[uint64]*session.Session,
		)
	}

	return r
}

func (r *Registry) shard(
	id uint64,
) *bucket {
	return &r.shards[id%ShardCount]
}

func (r *Registry) Add(
	s *session.Session,
) {
	b := r.shard(s.ID)

	b.Lock()
	b.clients[s.ID] = s
	b.Unlock()
}

func (r *Registry) Remove(
	id uint64,
) {
	b := r.shard(id)

	b.Lock()
	delete(b.clients, id)
	b.Unlock()
}

func (r *Registry) Range(
	fn func(*session.Session),
) {
	for i := 0; i < ShardCount; i++ {
		b := &r.shards[i]

		b.RLock()

		for _, s := range b.clients {
			fn(s)
		}

		b.RUnlock()
	}
}
