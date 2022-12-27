package main

import (
	"io"
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("primes function was invoked with")
	var maximun int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}

		if number := req.Number; number > maximun {
			maximun = number
			err = stream.Send(&pb.MaxResponse{
				Result: maximun,
			})

			if err != nil {
				log.Fatalf("error while sending data to client stream: %v", err)
			}
		}
	}
}
