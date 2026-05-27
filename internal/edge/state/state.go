package state

import "sync/atomic"

type Runtime struct {
	Connections atomic.Uint64
	Messages    atomic.Uint64
	Broadcasts  atomic.Uint64
	Dropped     atomic.Uint64
}
