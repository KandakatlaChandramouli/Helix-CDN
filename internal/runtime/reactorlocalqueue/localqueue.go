package reactorlocalqueue

type Queue struct {
	C chan []byte
}

func New(
	size int,
) *Queue {

	return &Queue{
		C: make(
			chan []byte,
			size,
		),
	}
}
