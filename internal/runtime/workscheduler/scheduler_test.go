package workscheduler

import (
	"testing"

	"helixcdn/internal/runtime/tasks"
)

func TestScheduler(
	t *testing.T,
) {
	called := false

	s := New(2)

	s.Submit(
		0,
		tasks.New(
			1,
			func() {
				called = true
			},
		),
	)

	s.RunOnce()

	if !called {
		t.Fatal("scheduler failed")
	}
}
