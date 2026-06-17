package ringbuffer

import "sync/atomic"

type RingBuffer struct {
	buffer []uint64
	mask   uint64

	head atomic.Uint64
	tail atomic.Uint64
}

func New(size uint64) *RingBuffer {
	return &RingBuffer{
		buffer: make([]uint64, size),
		mask:   size - 1,
	}
}

func (r *RingBuffer) Push(v uint64) bool {
	head := r.head.Load()
	tail := r.tail.Load()

	if head-tail >= uint64(len(r.buffer)) {
		return false
	}

	r.buffer[head&r.mask] = v

	r.head.Store(head + 1)

	return true
}

func (r *RingBuffer) Pop() (uint64, bool) {
	tail := r.tail.Load()
	head := r.head.Load()

	if tail >= head {
		return 0, false
	}

	v := r.buffer[tail&r.mask]

	r.tail.Store(tail + 1)

	return v, true
}
