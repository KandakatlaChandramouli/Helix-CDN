package mpsc

import (
	"runtime"
	"sync/atomic"
)

type node[T any] struct {
	value T
	next  atomic.Pointer[node[T]]
}

type Queue[T any] struct {
	head atomic.Pointer[node[T]]
	tail atomic.Pointer[node[T]]
}

func New[T any]() *Queue[T] {
	n := &node[T]{}

	q := &Queue[T]{}

	q.head.Store(n)
	q.tail.Store(n)

	return q
}

func (q *Queue[T]) Enqueue(v T) {
	n := &node[T]{
		value: v,
	}

	prev := q.head.Swap(n)

	prev.next.Store(n)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	tail := q.tail.Load()

	next := tail.next.Load()

	if next == nil {
		var zero T
		return zero, false
	}

	q.tail.Store(next)

	runtime.KeepAlive(tail)

	return next.value, true
}
