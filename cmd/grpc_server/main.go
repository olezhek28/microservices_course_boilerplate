package main

import (
	"context"
	"flag"
	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/dmtrybogdanov/auth/internal/config"
	"github.com/dmtrybogdanov/auth/internal/config/env"
	desc "github.com/dmtrybogdanov/auth/pkg/auth_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedAuthV1Server
	pool *pgxpool.Pool
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("create name: %s", req.Name)

	builderSelectOne := sq.Select("id").
		From("auth").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"name": req.Name}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var id int64
	var name string

	err = s.pool.QueryRow(ctx, query, args...).Scan(&id, &name)
	if err != nil {
		log.Fatalf("failed to select notes: %v", err)
	}

	log.Printf("id: %d, usernames: %s: %v\n", id, name)

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("get id: %d", req.GetId())

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("update id: %d", req.GetId())

	return nil, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("delete id: %d", req.GetId())

	return nil, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
