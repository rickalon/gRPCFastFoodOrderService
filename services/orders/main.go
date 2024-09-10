package main

func main() {
	httpServer := NewHttpServe(":8000")
	go httpServer.Run()
	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()
}
