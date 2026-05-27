package reactorrouter

func Route(
	topic string,
	shards int,
) int {

	if shards == 0 {
		return 0
	}

	sum := 0

	for _, c := range topic {
		sum += int(c)
	}

	return sum % shards
}
