package protocol

import "encoding/binary"

const (
	HeaderSize = 16
)

type Message struct {
	Opcode uint16
	Topic  uint32
	Size   uint32
	ID     uint64
	Data   []byte
}

func Encode(
	msg *Message,
) []byte {
	size := HeaderSize + len(msg.Data)

	buf := make([]byte, size)

	binary.LittleEndian.PutUint16(
		buf[0:2],
		msg.Opcode,
	)

	binary.LittleEndian.PutUint32(
		buf[2:6],
		msg.Topic,
	)

	binary.LittleEndian.PutUint32(
		buf[6:10],
		msg.Size,
	)

	binary.LittleEndian.PutUint64(
		buf[10:18],
		msg.ID,
	)

	copy(buf[HeaderSize:], msg.Data)

	return buf
}
