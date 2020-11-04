package perf

import (
	log "code.google.com/p/log4go"
	"net"
	"net/http"
	"net/http/pprof"
	"time"
)

const (
	httpReadTimeout = 15
)

var (
	closed bool

	httpListener []net.Listener
)

// StartPprof start http pprof.
func Init(pprofBind []string) error {
	pprofServeMux := http.NewServeMux()
	pprofServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	server := &http.Server{Handler: pprofServeMux, ReadTimeout: httpReadTimeout * time.Second}
	for _, addr := range pprofBind {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("net.Listen(\"tcp\", \"%s\") error(%v)", addr, err)
			return err
		}
		httpListener = append(httpListener, l)
		go func() {
			if err := server.Serve(l); err != nil {
				if !closed {
					log.Error("server.Serve(\"%s\") error(%v)", addr, err)
					panic(err)
				}
			}
		}()
	}
	return nil
}

// Close close the resource.
func Close() {
	closed = true
	for _, l := range httpListener {
		if err := l.Close(); err != nil {
			log.Error("l.Close() error(%v)", err)
		}
	}
	httpListener = []net.Listener{}
}
