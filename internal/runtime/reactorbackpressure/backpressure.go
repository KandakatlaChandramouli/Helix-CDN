package reactorbackpressure

import "sync/atomic"

type Limiter struct {
	Dropped atomic.Uint64
}

func New() *Limiter {
	return &Limiter{}
}

func (l *Limiter) Drop() {
	l.Dropped.Add(1)
}
