package workscheduler

import (
	"runtime"

	"helixcdn/internal/runtime/tasks"
	"helixcdn/internal/runtime/workqueue"
	"helixcdn/internal/runtime/workstealing"
)

type Scheduler struct {
	workers []*workqueue.Queue
}

func New(
	n int,
) *Scheduler {
	s := &Scheduler{
		workers: make([]*workqueue.Queue, n),
	}

	for i := 0; i < n; i++ {
		s.workers[i] = workqueue.New()
	}

	return s
}

func (s *Scheduler) Submit(
	worker int,
	task tasks.Task,
) {
	s.workers[worker].Push(task)
}

func (s *Scheduler) RunOnce() {
	for _, q := range s.workers {
		workstealing.Execute(q)
	}

	runtime.Gosched()
}
