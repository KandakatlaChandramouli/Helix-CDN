package http

import (
	"encoding/json"
	nethttp "net/http"

	"helixcdn/internal/runtime/stats"
)

func Start() error {
	nethttp.HandleFunc(
		"/stats",
		func(
			w nethttp.ResponseWriter,
			r *nethttp.Request,
		) {
			_ = json.NewEncoder(w).Encode(
				stats.Read(),
			)
		},
	)

	return nethttp.ListenAndServe(
		":18080",
		nil,
	)
}
