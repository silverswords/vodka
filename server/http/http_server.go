package http

import (
	"fmt"
	"net"
	"net/http"

	"github.com/yusank/vodka/server"
)

type httpsServer struct {
	opts    *server.Options
	handler http.Handler
	exit    chan error
}

func (h *httpsServer) Start() error {
	ln, err := net.Listen("tcp", h.opts.Address)
	if err != nil {
		return err
	}

	go http.Serve(ln, h.handler)

	go func() {
	loop:
		for {
			select {
			case <-h.exit:
				break loop
			}
		}

		_ = ln.Close()
	}()

	return nil
}

func (h *httpsServer) Stop() {
	h.exit <- fmt.Errorf("stop")
}

func (h *httpsServer) Handle(i interface{}) error {
	hdl, ok := i.(http.Handler)
	if !ok {
		return fmt.Errorf("unkonw handler")
	}

	h.handler = hdl
	return nil
}

func (h *httpsServer) Name() string {
	return "http"
}

func NewServer(opts ...server.Option) server.Server {
	defOpts := server.NewOptions()
	for _, o := range opts {
		o(defOpts)
	}

	hsrv := &httpsServer{opts: defOpts}
	return hsrv
}
