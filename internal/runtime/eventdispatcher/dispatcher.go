package eventdispatcher

type Dispatcher struct {
	queue chan func()
}

func New(
	size int,
) *Dispatcher {
	return &Dispatcher{
		queue: make(chan func(), size),
	}
}

func (d *Dispatcher) Submit(
	fn func(),
) {
	d.queue <- fn
}

func (d *Dispatcher) RunOnce() {
	select {
	case fn := <-d.queue:
		fn()
	default:
	}
}
