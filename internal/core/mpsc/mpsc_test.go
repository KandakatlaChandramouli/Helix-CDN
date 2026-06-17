package mpsc

import (
	"sync"
	"testing"
)

func TestQueueBasic(
	t *testing.T,
) {
	q := New[uint64]()

	q.Enqueue(10)

	v, ok := q.Dequeue()

	if !ok {
		t.Fatal("expected value")
	}

	if v != 10 {
		t.Fatalf("unexpected value: %d", v)
	}
}

func TestQueueEmpty(
	t *testing.T,
) {
	q := New[uint64]()

	_, ok := q.Dequeue()

	if ok {
		t.Fatal("expected empty queue")
	}
}

func TestQueueConcurrent(
	t *testing.T,
) {
	q := New[uint64]()

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 10000; j++ {
				q.Enqueue(uint64(j))
			}
		}()
	}

	wg.Wait()
}
