package reactorscheduler

import "time"

type Scheduler struct {
	Tick time.Duration
}

func New() *Scheduler {

	return &Scheduler{
		Tick: 10 * time.Millisecond,
	}
}
