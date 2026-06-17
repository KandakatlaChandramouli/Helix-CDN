package workqueue

import (
	"testing"

	"helixcdn/internal/runtime/tasks"
)

func TestQueue(
	t *testing.T,
) {
	q := New()

	q.Push(tasks.New(1, nil))

	_, ok := q.Pop()

	if !ok {
		t.Fatal("pop failed")
	}
}
