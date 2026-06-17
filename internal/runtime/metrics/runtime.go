package metrics

import "sync/atomic"

var Connections atomic.Uint64
var Messages atomic.Uint64
var Broadcasts atomic.Uint64
