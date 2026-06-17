package runtimebalance

import (
	"sync/atomic"
)

type Balancer struct {
	next atomic.Uint64
	size uint64
}

func New(
	size uint64,
) *Balancer {
	return &Balancer{
		size: size,
	}
}

func (b *Balancer) Next() uint64 {
	v := b.next.Add(1)

	return v % b.size
}
