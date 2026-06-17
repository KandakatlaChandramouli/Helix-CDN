package reactor

type Reactor struct {
	ID int
}

func New(
	id int,
) *Reactor {
	return &Reactor{
		ID: id,
	}
}
