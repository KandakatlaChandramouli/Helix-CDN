package reactorsequence

import "sync/atomic"

type Sequence struct {
	v atomic.Uint64
}

func (s *Sequence) Next() uint64 {
	return s.v.Add(1)
}
