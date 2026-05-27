package broadcast

import (
	"helixcdn/internal/edge/pool"
	"helixcdn/internal/edge/session"
)

type Engine struct {
	registry *session.Registry
}

func New(
	registry *session.Registry,
) *Engine {
	return &Engine{
		registry: registry,
	}
}

func (e *Engine) Broadcast(
	payload []byte,
) {
	buf := pool.Get()

	copy(*buf, payload)

	e.registry.Range(func(
		client *session.Client,
	) {
		_ = client
	})

	pool.Put(buf)
}
