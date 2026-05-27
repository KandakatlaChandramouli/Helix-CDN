package fanoutpool

import (
	"github.com/panjf2000/ants/v2"
)

func NewPool() (*ants.Pool, error) {

	return ants.NewPool(
		1024,
	)
}
