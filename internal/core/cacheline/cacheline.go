package cacheline

const Size = 64

type Pad struct {
	_ [Size]byte
}
