package opcode

type Opcode uint8

const (
	Heartbeat Opcode = 1
	Broadcast Opcode = 2
	Subscribe Opcode = 3
	Unsubscribe Opcode = 4
	Ack Opcode = 5
)
