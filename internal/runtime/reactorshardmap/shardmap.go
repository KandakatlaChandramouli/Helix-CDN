package reactorshardmap

import "sync"

type ShardMap struct {
	mu sync.RWMutex

	shards map[uint64][][]byte
}

func New() *ShardMap {

	return &ShardMap{
		shards: make(
			map[uint64][][]byte,
		),
	}
}

func (s *ShardMap) Append(
	shard uint64,
	payload []byte,
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.shards[shard] = append(
		s.shards[shard],
		payload,
	)
}
