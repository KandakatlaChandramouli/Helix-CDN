package queue

type Queue struct {
	C chan []byte
}

func NewQueue(size int) *Queue {
	return &Queue{
		C: make(chan []byte, size),
	}
}
