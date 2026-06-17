package main

import (
	"log"

	"helixcdn/internal/runtime/gnetserver"

	"github.com/panjf2000/gnet/v2"
)

func main() {

	server := gnetserver.New()

	log.Fatal(
		gnet.Run(
			server,
			"tcp://:9100",
			gnet.WithMulticore(true),
			gnet.WithReusePort(true),
		),
	)
}
