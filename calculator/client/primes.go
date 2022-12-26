package main

import (
	"context"
	"io"
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimesRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling primes: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}
}
