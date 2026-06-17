package epoll

type Poller struct {
	MaxEvents int
}

func New() *Poller {
	return &Poller{
		MaxEvents: 4096,
	}
}
