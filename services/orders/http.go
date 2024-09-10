package main

import (
	"log"
	"net/http"

	handler "github.com/rickalon/gRPCFastFoodOrderService/services/orders/handler/orders"
	"github.com/rickalon/gRPCFastFoodOrderService/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServe(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)
	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
