package app

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	user_v12 "github.com/neracastle/auth/api/user_v1"
	grpc_server "github.com/neracastle/auth/internal/grpc-server"
	"github.com/neracastle/auth/internal/grpc-server/interceptors"
	"github.com/neracastle/auth/internal/kafka"
	"github.com/neracastle/auth/pkg/user_v1"
)

// App приложение
type App struct {
	grpc        *grpc.Server
	srvProvider *serviceProvider
}

// NewApp новый экземпляр приложения
func NewApp(ctx context.Context) *App {
	app := &App{srvProvider: newServiceProvider()}
	app.init(ctx)
	return app
}

func (a *App) init(ctx context.Context) {
	lg := logger.SetupLogger(a.srvProvider.Config().Env)
	a.grpc = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptors.RequestIDInterceptor,
			interceptors.NewLoggerInterceptor(lg)),
	)

	reflection.Register(a.grpc)
	user_v1.RegisterUserV1Server(a.grpc, grpc_server.NewServer(a.srvProvider.UsersService(ctx)))
}

// Start запускает сервис на прием запросов
func (a *App) Start() error {
	conn, err := net.Listen("tcp", a.srvProvider.Config().GRPC.Address())
	if err != nil {
		return err
	}

	log.Printf("UserAPI service started on %s\n", a.srvProvider.Config().GRPC.Address())

	if err = a.grpc.Serve(conn); err != nil {
		return err
	}

	return nil
}

// StartHTTP запускает http сервис на прием запросов
func (a *App) StartHTTP() error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	_ = user_v1.RegisterUserV1HandlerFromEndpoint(context.Background(), mux, a.srvProvider.Config().GRPC.Address(), opts)

	httpServer := &http.Server{
		Addr:              a.srvProvider.Config().HTTP.Address(),
		Handler:           NewCORSMux(mux),
		ReadHeaderTimeout: 5 * time.Second, //защита от Slowloris Attack
	}

	log.Printf("UserAPI HTTP started on %s\n", a.srvProvider.Config().HTTP.Address())

	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// RunTopicLogger запускает прослушку сообщений кафки и просто их логгирует
func (a *App) RunTopicLogger(ctx context.Context) {
	topic := a.srvProvider.Config().NewUsersTopic
	lg := logger.SetupLogger(a.srvProvider.Config().Env)
	lg = lg.With(slog.String("topic", topic))

	cons := a.srvProvider.KafkaConsumer()
	cons.GroupHandler().SetMessageHandler(func(ctx context.Context, msg *sarama.ConsumerMessage) error {
		lg.Info("received message", slog.String("value", string(msg.Value)))

		return nil
	})

	for {
		err := cons.RunConsume(ctx, topic)
		if err != nil {
			if errors.Is(kafka.ErrHandlerError, err) {
				lg.Info("consumer handler error", slog.String("error", err.Error()))
				continue
			}

			lg.Info("consumer error", slog.String("error", err.Error()))
			break
		}
	}

}

// StartSwaggerServer запускает сервер со swagger-документацией
func (a *App) StartSwaggerServer() error {
	mux := http.NewServeMux()
	mux.Handle("/", user_v12.NewSwaggerFS(a.srvProvider.Config().HTTP.Port))

	httpServer := &http.Server{
		Addr:              a.srvProvider.Config().Swagger.Address(),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second, //защита от Slowloris Attack
	}

	log.Printf("Swagger server started on %s\n", a.srvProvider.Config().Swagger.Address())
	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// Shutdown мягко закрывает все соединения и службы
func (a *App) Shutdown(ctx context.Context) {

	allClosed := make(chan struct{})
	go func() {
		_ = a.srvProvider.consumer.Close()
		_ = a.srvProvider.producer.Close()
		_ = a.srvProvider.dbc.Close()
		a.grpc.GracefulStop()
		close(allClosed)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-allClosed:
			return
		}
	}
}
