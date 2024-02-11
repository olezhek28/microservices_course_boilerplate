package main

import (
	"context"
	"fmt"
	"log"
	"time"

	userDesc "github.com/a1exCross/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

	createRes, err := cl.Create(ctx, &userDesc.CreateRequest{
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

	user, err := cl.Get(ctx, &userDesc.GetRequest{
		Id: createRes.Id,
	})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	_, err = cl.Update(ctx, &userDesc.UpdateRequest{
		Id:    user.User.Id,
		Name:  wrapperspb.String("new name"),
		Email: wrapperspb.String("new email"),
		Role:  userDesc.UserRole_user,
	})
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	_, err = cl.Delete(ctx, &userDesc.DeleteRequest{
		Id: createRes.Id,
	})
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}
}
