package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userDesc "github.com/a1exCross/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort = 50051
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect grpc: %v", err)
	}

	cl := userDesc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := cl.Create(ctx, &userDesc.CreateRequest{
		Info: &userDesc.UserInfo{
			Name:  "test connection",
			Email: "test connection",
			Role:  userDesc.UserRole_admin,
		},
		Pass: &userDesc.UserPassword{
			Password:        "test connection",
			PasswordConfirm: "test connection",
		},
	})
	if err != nil {
		log.Fatalf("failed to craete user: %v", err)
	}

	_ = user
}
