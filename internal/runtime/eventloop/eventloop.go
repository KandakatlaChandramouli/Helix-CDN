package eventloop

type Loop struct {
	ID int
	Load uint64
}

func NewLoop(
	id int,
) *Loop {

	return &Loop{
		ID: id,
	}
}
