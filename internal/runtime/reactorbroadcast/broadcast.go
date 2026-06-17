package reactorbroadcast

import (
	"helixcdn/internal/runtime/reactormetrics"

	"github.com/panjf2000/gnet/v2"
)

func Broadcast(
	c gnet.Conn,
	payload []byte,
) {

	reactormetrics.Broadcasts.Add(1)

	reactormetrics.BytesOut.Add(
		uint64(len(payload)),
	)

	_ = c.AsyncWrite(
		payload,
		nil,
	)
}
