package main

import (
	"context"
	"log"

	pb "github.com/BFDavidGamboa/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not David",
		Title:    "A new tittle",
		Content:  "Content of the first blog, with some awsome additions",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("error happened while updating blog: %v\n", err)
	}

	log.Println("Bloga was updated")

}
