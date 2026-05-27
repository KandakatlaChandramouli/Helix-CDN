package reactorheartbeat

import "time"

func Interval() time.Duration {
	return 30 * time.Second
}
