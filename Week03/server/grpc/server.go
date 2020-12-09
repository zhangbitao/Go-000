package grpc

import (
	"context"
	"fmt"
	"net"

	"Week03/server"

	"google.golang.org/grpc"
)

var _ server.Server = new(grpcServer)

type grpcServer struct {
	*grpc.Server
	port int16
}

func New(port int16) server.Server {
	srv := grpc.NewServer()
	return &grpcServer{Server: srv, port: port}
}

func (s *grpcServer) Start(_ context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

func (s *grpcServer) Stop(_ context.Context) error {
	s.GracefulStop()
	return nil
}
