package main

import (
	"context"
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("error while opening stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number%d\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response: %v\n", err)
	}

	log.Printf("Avg %f\n", res.Result)

}
