package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	authpb "user/gen/go/auth"
	"user/internal/api/grpcapi"
)

type Server struct {
	port       int
	gRPCServer *grpc.Server
}

func NewServerRPC(port int) *Server {
	gRPC := grpc.NewServer()

	authpb.RegisterAuthServiceServer(gRPC, grpcapi.NewServerAPI())

	return &Server{
		port:       port,
		gRPCServer: gRPC,
	}
}

func (s *Server) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}

func (s *Server) Stop() {
	const op = "grpcapp.Stop"
	log.Println("stopping grpc Server")

	s.gRPCServer.GracefulStop()
}
