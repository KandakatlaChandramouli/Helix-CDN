package pool

import "sync"

const (
	BufferSize = 4096
)

var BufferPool = sync.Pool{
	New: func() any {
		buf := make([]byte, BufferSize)
		return &buf
	},
}

func Get() *[]byte {
	return BufferPool.Get().(*[]byte)
}

func Put(buf *[]byte) {
	BufferPool.Put(buf)
}
