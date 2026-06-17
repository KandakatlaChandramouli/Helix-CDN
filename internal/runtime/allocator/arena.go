package allocator

type Arena struct {
	Size int
}

func New(
	size int,
) *Arena {
	return &Arena{
		Size: size,
	}
}
