package node

import (
	"crypto/rand"
	"encoding/hex"
)

type Node struct {
	ID string
}

func New() *Node {

	buf := make([]byte, 16)

	_, _ = rand.Read(buf)

	return &Node{
		ID: hex.EncodeToString(buf),
	}
}
