package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	productpb "product/gen/go/product"
	"product/internal/api/grpcapi"
	"product/internal/service"
)

// Server represents the gRPC server for the Product Service.
type Server struct {
	port       int
	gRPCServer *grpc.Server
}

func NewServerRPC(port int, productService *service.Service) *Server {
	gRPC := grpc.NewServer()

	productpb.RegisterProductServiceServer(gRPC, grpcapi.NewServerAPI(productService))

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

	log.Println(op, " Stopping gRPC server")

	s.gRPCServer.GracefulStop()
}
