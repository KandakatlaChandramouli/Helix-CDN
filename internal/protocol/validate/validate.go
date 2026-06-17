package validate

func Payload(
	buf []byte,
) bool {

	return len(buf) > 0
}
