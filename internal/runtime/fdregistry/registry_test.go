package fdregistry

import "testing"

func TestRegistry(
	t *testing.T,
) {
	r := New()

	r.Add(10, "socket")

	v, ok := r.Get(10)

	if !ok {
		t.Fatal("missing fd")
	}

	if v != "socket" {
		t.Fatal("invalid kind")
	}
}
