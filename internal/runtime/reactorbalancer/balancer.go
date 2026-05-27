package reactorbalancer

type Balancer struct {
	Shards int
}

func New(
	shards int,
) *Balancer {

	return &Balancer{
		Shards: shards,
	}
}
