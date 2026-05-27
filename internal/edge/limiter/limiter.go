package limiter

import "sync/atomic"

type Limiter struct {
	current atomic.Int64
	max     int64
}

func New(max int64) *Limiter {
	return &Limiter{
		max: max,
	}
}

func (l *Limiter) Acquire() bool {
	current := l.current.Add(1)

	if current > l.max {
		l.current.Add(-1)
		return false
	}

	return true
}

func (l *Limiter) Release() {
	l.current.Add(-1)
}
