package tasks

type Task struct {
	ID uint64
	Fn func()
}

func New(
	id uint64,
	fn func(),
) Task {
	return Task{
		ID: id,
		Fn: fn,
	}
}

func (t *Task) Run() {
	if t.Fn != nil {
		t.Fn()
	}
}
