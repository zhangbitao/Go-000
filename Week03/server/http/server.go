package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"Week03/server"
)

var _ server.Server = new(httpServer)

type httpServer struct {
	*http.Server
	port int16
}

func New(port int16) server.Server {
	srv := &http.Server{}
	return &httpServer{Server: srv, port: port}
}

func (s *httpServer) Start(_ context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

func (s *httpServer) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
