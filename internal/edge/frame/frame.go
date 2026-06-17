package frame

import (
	"encoding/binary"
	"errors"
)

const (
	HeaderSize = 12
)

type Frame struct {
	Opcode uint16
	Topic  uint32
	Size   uint32
	Data   []byte
}

func Decode(
	buf []byte,
) (*Frame, error) {
	if len(buf) < HeaderSize {
		return nil, errors.New(
			"short frame",
		)
	}

	size := binary.LittleEndian.Uint32(
		buf[6:10],
	)

	if len(buf) < int(HeaderSize+size) {
		return nil, errors.New(
			"incomplete frame",
		)
	}

	return &Frame{
		Opcode: binary.LittleEndian.Uint16(
			buf[0:2],
		),
		Topic: binary.LittleEndian.Uint32(
			buf[2:6],
		),
		Size: size,
		Data: buf[HeaderSize : HeaderSize+size],
	}, nil
}
