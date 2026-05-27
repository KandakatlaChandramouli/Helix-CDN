package reactorpool

import "sync"

var Buffers = sync.Pool{
	New: func() any {
		return make(
			[]byte,
			4096,
		)
	},
}
