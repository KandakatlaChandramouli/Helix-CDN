package main

import (
	"log"
	"net"

	"helixcdn/internal/runtime/node"
)

func handleConnection(c net.Conn) {
	defer c.Close()

	buffer := make([]byte, 4096)

	for {
		_, err := c.Read(buffer)

		if err != nil {
			return
		}
	}
}

func main() {

	n := node.New()

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
