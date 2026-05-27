package reactormetrics

import "sync/atomic"

var Connections atomic.Uint64
var Messages atomic.Uint64
var Broadcasts atomic.Uint64
var BytesIn atomic.Uint64
var BytesOut atomic.Uint64
