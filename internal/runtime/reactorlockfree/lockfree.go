package reactorlockfree

import "sync/atomic"

type Counter struct {
	v atomic.Uint64
}

func (c *Counter) Inc() {
	c.v.Add(1)
}
