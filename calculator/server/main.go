package main

import (
	"log"
	"net"

	pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening on %s: %v", addr, err)
	}

	log.Printf("Listening at %s\n", addr)

	// opts := []grpc.ServerOption{}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}