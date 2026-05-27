package batchercore

type Batcher struct {
	size int
	data []func()
}

func New(
	size int,
) *Batcher {
	return &Batcher{
		size: size,
		data: make([]func(), 0, size),
	}
}

func (b *Batcher) Add(
	fn func(),
) bool {
	b.data = append(b.data, fn)

	return len(b.data) >= b.size
}

func (b *Batcher) Flush() {
	for _, fn := range b.data {
		fn()
	}

	b.data = b.data[:0]
}
