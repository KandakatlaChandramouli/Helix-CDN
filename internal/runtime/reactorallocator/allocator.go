package reactorallocator

type Allocator struct {
	PageSize int
}

func New() *Allocator {

	return &Allocator{
		PageSize: 4096,
	}
}
