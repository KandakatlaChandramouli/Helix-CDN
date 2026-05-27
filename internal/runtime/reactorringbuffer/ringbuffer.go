package reactorringbuffer

type Ring struct {
	buf  [][]byte
	head uint64
	tail uint64
	size uint64
}

func New(size uint64) *Ring {

	return &Ring{
		buf:  make([][]byte, size),
		size: size,
	}
}

func (r *Ring) Push(v []byte) {

	r.buf[r.tail%r.size] = v

	r.tail++
}

func (r *Ring) Get(idx uint64) []byte {

	return r.buf[idx%r.size]
}
