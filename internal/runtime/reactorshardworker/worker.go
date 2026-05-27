package reactorshardworker

type Worker struct {
	ID uint64

	Processed uint64
}

func New(
	id uint64,
) *Worker {

	return &Worker{
		ID: id,
	}
}
