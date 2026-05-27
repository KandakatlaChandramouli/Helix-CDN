package reactor

type Reactor struct{}

func New() *Reactor {
	return &Reactor{}
}

func (r *Reactor) Run() {}
