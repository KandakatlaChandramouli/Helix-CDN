package fanout

import (
	"helixcdn/internal/runtime/connection"
	"helixcdn/internal/runtime/metrics"
)

type Broadcaster struct {
	manager *connection.Manager
}

func NewBroadcaster(
	manager *connection.Manager,
) *Broadcaster {

	return &Broadcaster{
		manager: manager,
	}
}

func (b *Broadcaster) Broadcast(
	payload []byte,
) {

	b.manager.Range(
		func(
			client *connection.Client,
		) {

			if client.Conn != nil {

				_, _ = client.Conn.Write(
					payload,
				)
			}
		},
	)

	metrics.Broadcasts.Add(1)
}
