package main

import (
	"context"
	"log"
	"net"

	pb "github.com/BFDavidGamboa/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Error received when trying to establish new Client to Mongo %v\n", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("Error received when trying to connect with Mongo %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening on %s: %v", addr, err)
	}

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	tsl := true // change that to false if needed

	if tsl {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &Server{})

	defer s.Stop()
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
