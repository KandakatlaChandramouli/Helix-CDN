package stats

import (
	"runtime"
	"sync/atomic"
)

var (
	Connections atomic.Uint64
	Messages    atomic.Uint64
	Broadcasts  atomic.Uint64
)

type Snapshot struct {
	Connections uint64
	Messages    uint64
	Broadcasts  uint64
	Goroutines  int
}

func Read() Snapshot {
	return Snapshot{
		Connections: Connections.Load(),
		Messages:    Messages.Load(),
		Broadcasts:  Broadcasts.Load(),
		Goroutines:  runtime.NumGoroutine(),
	}
}
