package fanout

import (
	"sync/atomic"

	"helixcdn/internal/runtime/queue"
)

type Worker struct {
	queue   *queue.MPSC
	handled atomic.Uint64
}

func NewWorker() *Worker {
	return &Worker{
		queue: queue.New(),
	}
}

func (w *Worker) Push(
	payload []byte,
) {
	w.queue.Push(payload)
}

func (w *Worker) Run() {
	for {
		_, ok := w.queue.Pop()

		if !ok {
			continue
		}

		w.handled.Add(1)
	}
}
