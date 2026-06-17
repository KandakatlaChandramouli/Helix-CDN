package session

import (
	"sync/atomic"
	"time"
)

type Session struct {
	ID        uint64
	Connected time.Time
	LastSeen  atomic.Int64
}

func New(
	id uint64,
) *Session {
	s := &Session{
		ID:        id,
		Connected: time.Now(),
	}

	s.Touch()

	return s
}

func (s *Session) Touch() {
	s.LastSeen.Store(
		time.Now().UnixNano(),
	)
}
