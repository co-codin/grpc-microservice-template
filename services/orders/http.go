package main

import (
	"grpc-microservice/services/orders/service"
	"log"
	"net/http"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
