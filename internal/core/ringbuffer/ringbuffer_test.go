package ringbuffer

import "testing"

func TestPushPop(
	t *testing.T,
) {
	rb := New(1024)

	ok := rb.Push(42)

	if !ok {
		t.Fatal("push failed")
	}

	v, ok := rb.Pop()

	if !ok {
		t.Fatal("pop failed")
	}

	if v != 42 {
		t.Fatalf("unexpected value: %d", v)
	}
}

func TestEmpty(
	t *testing.T,
) {
	rb := New(1024)

	_, ok := rb.Pop()

	if ok {
		t.Fatal("expected empty")
	}
}
