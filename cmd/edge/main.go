package main

import (
	"log"
	"net"
	"sync/atomic"

	runtimehttp "helixcdn/internal/runtime/http"

	"helixcdn/internal/runtime/connection"
	"helixcdn/internal/runtime/fanout"
	"helixcdn/internal/runtime/metrics"
	"helixcdn/internal/runtime/node"
)

var nextID atomic.Uint64

var manager = connection.New()

var broadcaster = fanout.NewBroadcaster(
	manager,
)

func handleConnection(
	c net.Conn,
) {
	defer c.Close()

	id := nextID.Add(1)

	manager.Add(
		&connection.Client{
			ID:   id,
			Conn: c,
		},
	)

	metrics.Connections.Store(
		uint64(
			manager.Count(),
		),
	)

	buffer := make([]byte, 4096)

	for {

		n, err := c.Read(buffer)

		if err != nil {
			break
		}

		if n > 0 {

			metrics.Messages.Add(1)

			broadcaster.Broadcast(
				buffer[:n],
			)
		}
	}

	manager.Remove(id)

	metrics.Connections.Store(
		uint64(
			manager.Count(),
		),
	)
}

func main() {

	n := node.New()

	go func() {
		if err := runtimehttp.Start(); err != nil {
			log.Println(err)
		}
	}()

	listener, err := net.Listen(
		"tcp",
		":9000",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"helix edge booted",
		n.ID,
	)

	for {

		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}
