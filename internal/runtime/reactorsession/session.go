package reactorsession

import "time"

type Session struct {
	ID string

	ConnectedAt time.Time
}
