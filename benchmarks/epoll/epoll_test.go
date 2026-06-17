package epoll

import "testing"

func BenchmarkEpoll(
	b *testing.B,
) {
	for i := 0; i < b.N; i++ {
	}
}
