package opcode

const (
	OpPing uint16 = iota + 1
	OpPong
	OpSubscribe
	OpUnsubscribe
	OpBroadcast
	OpReplay
)
