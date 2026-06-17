package metrics

import "sync/atomic"

var (
	Connections atomic.Uint64
	Broadcasts  atomic.Uint64
	Dropped     atomic.Uint64
	BytesIn     atomic.Uint64
	BytesOut    atomic.Uint64
)
