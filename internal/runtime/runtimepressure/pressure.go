package runtimepressure

import "sync/atomic"

type Pressure struct {
	queued atomic.Uint64
}

func (p *Pressure) Inc() {
	p.queued.Add(1)
}

func (p *Pressure) Dec() {
	p.queued.Add(^uint64(0))
}

func (p *Pressure) Load() uint64 {
	return p.queued.Load()
}
