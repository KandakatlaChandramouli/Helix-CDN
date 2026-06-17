package loadshedding

import "sync/atomic"

type Shedder struct {
	current atomic.Int64
	limit   int64
}

func New(limit int64) *Shedder {
	return &Shedder{
		limit: limit,
	}
}

func (s *Shedder) Allow() bool {
	current := s.current.Add(1)

	defer s.current.Add(-1)

	return current <= s.limit
}
