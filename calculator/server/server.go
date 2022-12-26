package main

import pb "github.com/BFDavidGamboa/grpc-go-course/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
