package worker

import "github.com/panjf2000/ants/v2"

type Pool struct {
	pool *ants.Pool
}

func New(size int) (*Pool, error) {
	p, err := ants.NewPool(size)

	if err != nil {
		return nil, err
	}

	return &Pool{
		pool: p,
	}, nil
}

func (p *Pool) Submit(
	task func(),
) error {
	return p.pool.Submit(task)
}
