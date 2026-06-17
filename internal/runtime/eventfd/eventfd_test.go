package eventfd

import "testing"

func TestEventFDCreate(
	t *testing.T,
) {
	ev, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if ev.FD <= 0 {
		t.Fatal("invalid eventfd")
	}
}
