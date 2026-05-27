package dispatcher

import (
	"runtime"

	"helixcdn/internal/runtime/queue"
)

type Dispatcher struct {
	queues []*queue.MPSC
}

func New() *Dispatcher {
	n := runtime.NumCPU()

	d := &Dispatcher{
		queues: make([]*queue.MPSC, n),
	}

	for i := 0; i < n; i++ {
		d.queues[i] = queue.New()
	}

	return d
}

func (d *Dispatcher) Queue(
	core int,
) *queue.MPSC {
	return d.queues[
		core%len(d.queues),
	]
}
