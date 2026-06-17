package fanout

import (
	"runtime"
)

type Sharded struct {
	workers []*Worker
}

func NewSharded() *Sharded {
	n := runtime.NumCPU()

	s := &Sharded{
		workers: make(
			[]*Worker,
			n,
		),
	}

	for i := 0; i < n; i++ {
		s.workers[i] = NewWorker()

		go s.workers[i].Run()
	}

	return s
}

func (s *Sharded) Dispatch(
	key uint64,
	payload []byte,
) {
	worker := s.workers[key%uint64(len(s.workers))]

	worker.Push(payload)
}
