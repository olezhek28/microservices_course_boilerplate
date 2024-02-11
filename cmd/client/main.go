package main

import (
	"context"
	"log"
	"time"

	desc "github.com/mchekalov/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server %v", err)
	}
	defer conn.Close()

	c := desc.NewUserV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: 333})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	log.Printf("User:\n %v", r.GetUser())
}
