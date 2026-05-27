package queue

import "sync/atomic"

type node struct {
	value []byte
	next  atomic.Pointer[node]
}

type MPSC struct {
	head atomic.Pointer[node]
	tail *node
}

func New() *MPSC {
	stub := &node{}

	q := &MPSC{
		tail: stub,
	}

	q.head.Store(stub)

	return q
}

func (q *MPSC) Push(
	payload []byte,
) {
	n := &node{
		value: payload,
	}

	prev := q.head.Swap(n)

	prev.next.Store(n)
}

func (q *MPSC) Pop() ([]byte, bool) {
	next := q.tail.next.Load()

	if next == nil {
		return nil, false
	}

	value := next.value

	q.tail = next

	return value, true
}
