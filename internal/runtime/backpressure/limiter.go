package backpressure

import "sync/atomic"

type Limiter struct {
	dropped atomic.Uint64
}

func New() *Limiter {
	return &Limiter{}
}

func (l *Limiter) Drop() {
	l.dropped.Add(1)
}

func (l *Limiter) Dropped() uint64 {
	return l.dropped.Load()
}
