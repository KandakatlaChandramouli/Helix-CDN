package eventdispatcher

import "testing"

func TestDispatcher(
	t *testing.T,
) {
	d := New(8)

	called := false

	d.Submit(func() {
		called = true
	})

	d.RunOnce()

	if !called {
		t.Fatal("dispatcher failed")
	}
}
