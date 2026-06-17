package mpsc

type Queue struct {
	C chan []byte
}

func NewMPSC(
	size int,
) *Queue {

	return &Queue{
		C: make(
			chan []byte,
			size,
		),
	}
}
