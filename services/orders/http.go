package main

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {

}
