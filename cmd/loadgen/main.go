package main

import (
	"net"
	"sync"
)

const (
	Connections = 10000
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < Connections; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			conn, err := net.Dial(
				"tcp",
				"127.0.0.1:9000",
			)

			if err != nil {
				return
			}

			defer conn.Close()

			select {}
		}()
	}

	wg.Wait()
}
