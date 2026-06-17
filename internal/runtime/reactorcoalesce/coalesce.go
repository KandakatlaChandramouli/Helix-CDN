package reactorcoalesce

import "sync"

type Coalescer struct {
	mu sync.Mutex

	events [][]byte
}

func New() *Coalescer {

	return &Coalescer{
		events: make(
			[][]byte,
			0,
			4096,
		),
	}
}

func (c *Coalescer) Add(
	event []byte,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.events = append(
		c.events,
		event,
	)
}

func (c *Coalescer) Flush() [][]byte {

	c.mu.Lock()
	defer c.mu.Unlock()

	out := c.events

	c.events = make(
		[][]byte,
		0,
		4096,
	)

	return out
}
