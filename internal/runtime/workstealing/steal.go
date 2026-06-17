package workstealing

import (
	"helixcdn/internal/runtime/tasks"
	"helixcdn/internal/runtime/workqueue"
)

func Steal(
	from *workqueue.Queue,
	to *workqueue.Queue,
) bool {
	task, ok := from.Pop()

	if !ok {
		return false
	}

	to.Push(task)

	return true
}

func Execute(
	q *workqueue.Queue,
) bool {
	task, ok := q.Pop()

	if !ok {
		return false
	}

	task.Run()

	return true
}

func Submit(
	q *workqueue.Queue,
	t tasks.Task,
) {
	q.Push(t)
}
