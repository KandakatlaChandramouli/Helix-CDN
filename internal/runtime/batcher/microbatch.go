package batcher

type MicroBatch struct {
	Items [][]byte
}

func NewMicroBatch() *MicroBatch {
	return &MicroBatch{
		Items: make([][]byte, 0, 1024),
	}
}
