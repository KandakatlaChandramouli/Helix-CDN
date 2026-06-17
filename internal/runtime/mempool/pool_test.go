package mempool

import "testing"

func TestPool(
	t *testing.T,
) {
	p := New(1024)

	b := p.Get()

	if len(b) != 1024 {
		t.Fatal("invalid buffer size")
	}

	p.Put(b)
}
