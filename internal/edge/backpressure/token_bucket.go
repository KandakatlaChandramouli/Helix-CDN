package backpressure

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu       sync.Mutex
	tokens   int64
	capacity int64
	refill   int64
	last     time.Time
}

func New(
	capacity int64,
	refill int64,
) *TokenBucket {
	return &TokenBucket{
		tokens:   capacity,
		capacity: capacity,
		refill:   refill,
		last:     time.Now(),
	}
}

func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()

	elapsed := now.Sub(b.last).Seconds()

	b.tokens += int64(
		elapsed * float64(b.refill),
	)

	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}

	b.last = now

	if b.tokens <= 0 {
		return false
	}

	b.tokens--

	return true
}
