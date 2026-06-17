package adaptivescheduler

import (
	"helixcdn/internal/runtime/runtimebalance"
	"helixcdn/internal/runtime/tasks"
	"helixcdn/internal/runtime/workscheduler"
)

type Scheduler struct {
	balance *runtimebalance.Balancer
	runtime *workscheduler.Scheduler
}

func New(
	workers uint64,
) *Scheduler {
	return &Scheduler{
		balance: runtimebalance.New(workers),
		runtime: workscheduler.New(int(workers)),
	}
}

func (s *Scheduler) Submit(
	task tasks.Task,
) {
	worker := s.balance.Next()

	s.runtime.Submit(
		int(worker),
		task,
	)
}

func (s *Scheduler) RunOnce() {
	s.runtime.RunOnce()
}
