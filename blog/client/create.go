package main

import (
	"context"
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----createBlog was invoked----")

	blog := &pb.Blog{
		AuthorId: "Clement",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error creating blog: %v", err)
	}

	log.Printf("Blog has been created successfully %s\n", res.Id)
	return res.Id
}
