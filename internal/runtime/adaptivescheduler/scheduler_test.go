package adaptivescheduler

import (
	"testing"

	"helixcdn/internal/runtime/tasks"
)

func TestAdaptiveScheduler(
	t *testing.T,
) {
	called := false

	s := New(4)

	s.Submit(
		tasks.New(
			1,
			func() {
				called = true
			},
		),
	)

	s.RunOnce()

	if !called {
		t.Fatal("adaptive scheduler failed")
	}
}
