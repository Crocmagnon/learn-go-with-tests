package main

import (
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/grpcserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &grpcserver.GreetServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
