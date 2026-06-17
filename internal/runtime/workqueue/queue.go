package workqueue

import (
	"sync"

	"helixcdn/internal/runtime/tasks"
)

type Queue struct {
	mu    sync.Mutex
	tasks []tasks.Task
}

func New() *Queue {
	return &Queue{
		tasks: make([]tasks.Task, 0),
	}
}

func (q *Queue) Push(
	t tasks.Task,
) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.tasks = append(q.tasks, t)
}

func (q *Queue) Pop() (tasks.Task, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 {
		return tasks.Task{}, false
	}

	t := q.tasks[0]

	q.tasks = q.tasks[1:]

	return t, true
}
