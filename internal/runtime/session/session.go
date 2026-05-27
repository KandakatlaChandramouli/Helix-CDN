package session

import "time"

type Session struct {
	ID        uint64
	CreatedAt time.Time
}
