package queue

import "sync"

var nodePool = sync.Pool{
	New: func() any {
		return &node{}
	},
}

func acquireNode() *node {
	return nodePool.Get().(*node)
}

func releaseNode(
	n *node,
) {
	n.value = nil
	n.next.Store(nil)

	nodePool.Put(n)
}
