package checksum

func Sum(
	buf []byte,
) uint32 {

	var out uint32

	for _, b := range buf {
		out += uint32(b)
	}

	return out
}
