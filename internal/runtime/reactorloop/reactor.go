package reactorloop

type Loop struct {
	ID int
}

func New(
	id int,
) *Loop {
	return &Loop{
		ID: id,
	}
}
