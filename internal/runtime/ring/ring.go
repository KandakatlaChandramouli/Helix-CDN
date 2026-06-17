package ring

type Ring struct {
	Size int
}

func New(
	size int,
) *Ring {
	return &Ring{
		Size: size,
	}
}
