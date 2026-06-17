package connection

import (
	"net"
	"sync"
)

type Client struct {
	ID   uint64
	Conn net.Conn
}

type Manager struct {
	mu sync.RWMutex

	clients map[uint64]*Client
}

func New() *Manager {
	return &Manager{
		clients: make(
			map[uint64]*Client,
		),
	}
}

func (m *Manager) Add(
	c *Client,
) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[c.ID] = c
}

func (m *Manager) Remove(
	id uint64,
) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(
		m.clients,
		id,
	)
}

func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.clients)
}

func (m *Manager) Range(
	fn func(*Client),
) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, client := range m.clients {
		fn(client)
	}
}
