package server

import (
	"log"

	"github.com/panjf2000/gnet/v2"
)

type Engine struct {
	gnet.BuiltinEventEngine
}

func (e *Engine) OnBoot(
	eng gnet.Engine,
) gnet.Action {
	log.Println("helix edge booted")

	return gnet.None
}

func Start() error {
	return gnet.Run(
		&Engine{},
		"tcp://:9000",
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
	)
}
