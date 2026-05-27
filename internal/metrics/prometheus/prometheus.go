package prometheus

import (
	"fmt"
	"net/http"
)

func Handler() http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		fmt.Fprintf(
			w,
			"helix_connections 0\n",
		)
	}
}
