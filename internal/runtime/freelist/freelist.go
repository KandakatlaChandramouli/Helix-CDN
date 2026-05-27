package freelist

import "sync"

type FreeList struct {
	pool sync.Pool
}

func New() *FreeList {

	return &FreeList{
		pool: sync.Pool{
			New: func() any {
				return make(
					[]byte,
					4096,
				)
			},
		},
	}
}
