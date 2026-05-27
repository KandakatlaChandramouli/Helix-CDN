package packet

type Packet struct {
	Version uint8
	Opcode uint8
	Payload []byte
}
