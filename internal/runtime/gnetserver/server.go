package gnetserver

import (
	"log"
	"strconv"

	"helixcdn/internal/runtime/reactorbroadcast"
	"helixcdn/internal/runtime/reactormetrics"
	"helixcdn/internal/runtime/reactorregistry"

	"github.com/panjf2000/gnet/v2"
)

type Server struct {
	gnet.BuiltinEventEngine

	registry *reactorregistry.Registry
}

func New() *Server {

	return &Server{
		registry: reactorregistry.New(),
	}
}

func (s *Server) OnBoot(
	eng gnet.Engine,
) gnet.Action {

	log.Println(
		"helix gnet reactor booted",
	)

	return gnet.None
}

func (s *Server) OnOpen(
	c gnet.Conn,
) (
	[]byte,
	gnet.Action,
) {

	id := strconv.Itoa(
		int(c.Fd()),
	)

	s.registry.Add(
		id,
		c,
	)

	reactormetrics.Connections.Add(1)

	return nil, gnet.None
}

func (s *Server) OnClose(
	c gnet.Conn,
	err error,
) gnet.Action {

	id := strconv.Itoa(
		int(c.Fd()),
	)

	s.registry.Remove(id)

	reactormetrics.Connections.Add(^uint64(0))

	return gnet.None
}

func (s *Server) OnTraffic(
	c gnet.Conn,
) gnet.Action {

	buf, _ := c.Next(-1)

	if len(buf) > 0 {

		reactormetrics.Messages.Add(1)

		reactormetrics.BytesIn.Add(
			uint64(len(buf)),
		)

		s.registry.Range(
			func(
				id string,
				conn gnet.Conn,
			) {

				reactorbroadcast.Broadcast(
					conn,
					buf,
				)
			},
		)
	}

	return gnet.None
}
