package reactordispatch

func Dispatch(
	shards uint64,
	key uint64,
) uint64 {

	if shards == 0 {
		return 0
	}

	return key % shards
}
