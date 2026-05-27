package tasks

import "testing"

func TestTaskRun(
	t *testing.T,
) {
	called := false

	task := New(
		1,
		func() {
			called = true
		},
	)

	task.Run()

	if !called {
		t.Fatal("task failed")
	}
}
