package main

import (
	"context"
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum was invoked with %v\n", in)
	response := &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}
	return response, nil
}
