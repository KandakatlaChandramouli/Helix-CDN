package batchercore

import "testing"

func TestBatcher(
	t *testing.T,
) {
	called := false

	b := New(2)

	b.Add(func() {
		called = true
	})

	b.Flush()

	if !called {
		t.Fatal("batch flush failed")
	}
}
