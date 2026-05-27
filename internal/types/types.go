package types

type NodeID string

type Topic string

type SessionID string

type Message struct {
	ID      uint64
	Topic   Topic
	Payload []byte
}
