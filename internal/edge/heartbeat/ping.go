package heartbeat

import "time"

const (
	PingInterval = 15 * time.Second
	Timeout      = 45 * time.Second
)
