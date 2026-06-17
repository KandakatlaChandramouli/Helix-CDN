package workstealing

import (
	"testing"

	"helixcdn/internal/runtime/tasks"
	"helixcdn/internal/runtime/workqueue"
)

func TestSteal(
	t *testing.T,
) {
	q1 := workqueue.New()
	q2 := workqueue.New()

	q1.Push(tasks.New(1, nil))

	ok := Steal(q1, q2)

	if !ok {
		t.Fatal("steal failed")
	}
}
