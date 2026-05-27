package http

import (
	"encoding/json"
	nethttp "net/http"

	"helixcdn/internal/runtime/metrics"
)

type Stats struct {
	Connections uint64 `json:"connections"`
	Messages    uint64 `json:"messages"`
	Broadcasts  uint64 `json:"broadcasts"`
}

func Start() error {

	nethttp.HandleFunc(
		"/stats",
		func(
			w nethttp.ResponseWriter,
			r *nethttp.Request,
		) {

			resp := Stats{
				Connections: metrics.Connections.Load(),
				Messages:    metrics.Messages.Load(),
				Broadcasts:  metrics.Broadcasts.Load(),
			}

			_ = json.NewEncoder(w).Encode(resp)
		},
	)

	return nethttp.ListenAndServe(
		":18080",
		nil,
	)
}
