package frame

type Frame struct {
	Opcode byte
	Topic  string
	Data   []byte
}
