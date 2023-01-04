package main

import (
	"context"
	"testing"

	"google.golang.org/grpc/credentials/insecure"
)

func TestGreet(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	

}
