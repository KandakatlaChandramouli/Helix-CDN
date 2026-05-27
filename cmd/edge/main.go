package main

import (
	"log"

	runtimehttp "helixcdn/internal/runtime/http"
	"helixcdn/internal/runtime/node"
	"helixcdn/internal/runtime/shutdown"
)

func main() {
	n := node.New()

	log.Println(
		"helix edge booted",
		n.ID,
	)

	go func() {
		if err := runtimehttp.Start(); err != nil {
			log.Println(err)
		}
	}()

	shutdown.Wait(func() {
		log.Println("helix edge shutdown")
	})
}
