package registerpprof

import (
	"net/http"
	"net/http/pprof"
)

func RegisterPprof(mux *http.ServeMux) {
	// Copy-paste from https://golang.org/src/net/http/pprof/pprof.go .
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
