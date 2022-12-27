package main

import pb "github.com/BFDavidGamboa/grpc-go-course/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
