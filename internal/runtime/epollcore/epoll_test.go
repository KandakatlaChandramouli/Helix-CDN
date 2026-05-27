package epollcore

import "testing"

func TestCreateEpoll(
	t *testing.T,
) {
	ep, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if ep.fd <= 0 {
		t.Fatal("invalid epoll fd")
	}
}
