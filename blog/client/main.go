package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/BFDavidGamboa/grpc-go-course/blog/proto"
)

var addr string = "localhost:50051"

func main() {

	tls := true //Change to false if
	// Unexpected gRPC error rpc error:
	// code = Unavailable desc = connection error: desc = "error reading server preface: EOF"

	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("error while loading CA trust certifications %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	// readBlog(c, id)
	// readBlog(c, "aNonExistingID")
	// updateBlog(c, id)
	// listBlog(c)
	deleteBlog(c, id)
}
