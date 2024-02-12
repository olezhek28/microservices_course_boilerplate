package main

import (
	"context"
	"github.com/dmtrybogdanov/auth/pkg/user_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	userID  = 10
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed connect to server: %v", err)
	}

	defer conn.Close()

	c := user_v1.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &user_v1.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("Failed to get user by id: %v", err)
	}
	log.Printf(color.RedString("User info: \n"), color.GreenString("%+v", r.Name))
}
