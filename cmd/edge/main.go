package main

import (
	"log"

	"helixcdn/internal/edge/server"
	"helixcdn/internal/logging"
	"helixcdn/internal/runtime"
)

func main() {
	runtime.Tune()

	if err := logging.Init(); err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
