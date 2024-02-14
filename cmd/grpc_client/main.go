package main

import (
	"context"

	desc "github.com/sarastee/auth/pkg/user_api_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		if err = conn.Close(); err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)

	c := desc.NewUserAPIV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Create(ctx, &desc.CreateRequest{
		User: &desc.UserInfo{
			Name:  "Ilya Lyakhov",
			Email: "ilja.sarasti@mail.ru",
			Role:  2,
		},
		Password:        "qwerty123",
		PasswordConfirm: "qwerty123",
	})
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	testUser, err := c.Get(ctx, &desc.GetRequest{
		Id: r.Id,
	})
	if err != nil {
		log.Fatalf("User ID: %v not found", r.Id)
	}

	log.Println(testUser.User.String())

	_, err = c.Update(ctx, &desc.UpdateRequest{
		Id: 0,
		Update: &desc.UpdateUserInfo{
			Name:  wrapperspb.String("Ilya Lyakhov"),
			Email: wrapperspb.String("ilja.sarasti@mail.ru"),
			Role:  2,
		},
	})
	if err != nil {
		log.Fatalf("failed to update ser ID: %v", r.Id)
	}

	testUser, err = c.Get(ctx, &desc.GetRequest{
		Id: r.Id,
	})
	if err != nil {
		log.Fatalf("User ID: %v not found", r.Id)
	}

	log.Println(testUser.User.String())

	_, err = c.Delete(ctx, &desc.DeleteRequest{
		Id: 0,
	})
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}

}
