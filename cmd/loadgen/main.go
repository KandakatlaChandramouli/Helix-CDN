package main

import (
	"log"
	"net"
	"sync"
	"time"
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

			for {
				_, err := conn.Write(
					[]byte("ping"),
				)

				if err != nil {
					return
				}

				time.Sleep(
					time.Second,
				)
			}
		}()
	}

	log.Println(
		"loadgen started",
	)

	wg.Wait()
}
