package replay

import "sync"

type Buffer struct {
	mu sync.RWMutex

	frames [][]byte
}

func New() *Buffer {
	return &Buffer{
		frames: make(
			[][]byte,
			0,
			1024,
		),
	}
}

func (b *Buffer) Add(
	frame []byte,
) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.frames = append(
		b.frames,
		frame,
	)
}

func (b *Buffer) Frames() [][]byte {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.frames
}
