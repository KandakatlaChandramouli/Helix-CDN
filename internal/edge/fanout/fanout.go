package fanout

import (
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
	e.registry.Range(func(
		client *session.Client,
	) {
		_ = client
	})
}
