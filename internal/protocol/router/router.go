package router

func Route(
	opcode uint8,
) string {

	switch opcode {

	case 1:
		return "heartbeat"

	case 2:
		return "broadcast"

	default:
		return "unknown"
	}
}
