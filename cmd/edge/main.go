package main

import (
	"log"

	"helixcdn/internal/logging"
	"helixcdn/internal/runtime"
)

func main() {
	runtime.Tune()

	if err := logging.Init(); err != nil {
		log.Fatal(err)
	}

	log.Println("helixcdn edge runtime booted")
}
