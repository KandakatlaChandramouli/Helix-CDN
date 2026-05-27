package router

import (
	"sync"

	"helixcdn/internal/edge/worker"
)

type Router struct {
	mu      sync.RWMutex
	workers []*worker.Pool
}

func New() *Router {
	return &Router{}
}

func (r *Router) AddWorker(
	pool *worker.Pool,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.workers = append(
		r.workers,
		pool,
	)
}
