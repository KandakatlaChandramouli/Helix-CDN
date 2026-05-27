package reuse

import "sync"

var Pool = sync.Pool{
	New: func() any {
		buf := make([]byte, 4096)
		return &buf
	},
}

func Get() *[]byte {
	return Pool.Get().(*[]byte)
}

func Put(
	buf *[]byte,
) {
	Pool.Put(buf)
}
