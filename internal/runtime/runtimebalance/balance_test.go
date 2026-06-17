package runtimebalance

import "testing"

func TestBalancer(
	t *testing.T,
) {
	b := New(4)

	v := b.Next()

	if v >= 4 {
		t.Fatal("invalid worker selection")
	}
}
