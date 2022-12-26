package main

import (
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stram pb.CalculatorService_PrimesServer) error {
	log.Printf("primes function was invoked with %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stram.Send(&pb.PrimesResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
