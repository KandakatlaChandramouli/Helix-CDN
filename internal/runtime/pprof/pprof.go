package pprof

import (
	"net/http"
	_ "net/http/pprof"
)

func Start() error {
	return http.ListenAndServe(
		":6060",
		nil,
	)
}
