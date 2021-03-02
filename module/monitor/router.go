package monitor

import (
	"net/http"

	_ "net/http/pprof"

	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

// init request path to handler
func init() {
	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":10108", nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":10109", nil)
}
