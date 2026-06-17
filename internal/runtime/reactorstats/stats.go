package reactorstats

type RuntimeStats struct {
	Connections uint64
	Messages    uint64
	Broadcasts  uint64
	BytesIn     uint64
	BytesOut    uint64
}
