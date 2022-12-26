package main

import (
	"context"
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	r, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 10, SecondNumber: 3})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", r.Result)
}
