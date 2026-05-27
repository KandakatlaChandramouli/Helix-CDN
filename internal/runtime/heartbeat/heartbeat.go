package heartbeat

import "time"

func Interval() time.Duration {
	return 30 * time.Second
}
