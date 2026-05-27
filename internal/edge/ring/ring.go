package ring

import "sync/atomic"

type Ring struct {
	mask uint64
	head atomic.Uint64
	data [][]byte
}

func New(size uint64) *Ring {
	return &Ring{
		mask: size - 1,
		data: make([][]byte, size),
	}
}

func (r *Ring) Push(buf []byte) {
	idx := r.head.Add(1)

	r.data[idx&r.mask] = buf
}

func (r *Ring) Get(idx uint64) []byte {
	return r.data[idx&r.mask]
}
