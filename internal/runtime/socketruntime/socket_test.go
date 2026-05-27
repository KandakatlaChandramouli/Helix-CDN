package socketruntime

import "testing"

func TestSocketRuntime(
	t *testing.T,
) {
	s, err := Listen("127.0.0.1:0")

	if err != nil {
		t.Fatal(err)
	}

	defer s.Close()
}
