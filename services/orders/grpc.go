package main

import (
	"log"
	"net"

	handler "github.com/rickalon/gRPCFastFoodOrderService/services/orders/handler/orders"
	"github.com/rickalon/gRPCFastFoodOrderService/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to create listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGRPCServer(grpcServer, orderService)

	log.Println("Starting gRPC server...")
	return grpcServer.Serve(lis)
}
