package reconnect

import (
	"math/rand"
	"time"
)

func Jitter(
	base time.Duration,
) time.Duration {
	offset := rand.Int63n(
		int64(base),
	)

	return base + time.Duration(offset)
}
