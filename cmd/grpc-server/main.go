package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/neracastle/auth/internal/app"
)

func main() {

	ctx := context.Background()
	ap := app.NewApp(ctx)

	go func() {
		err := ap.Start()
		if err != nil {
			log.Fatalf("failed to start app: %v", err)
		}
	}()

	go func() {
		err := ap.StartHTTP()
		if err != nil {
			log.Fatalf("failed to start http: %v", err)
		}
	}()

	go func() {
		err := ap.StartSwaggerServer()
		if err != nil {
			log.Printf("failed to start swagger: %v\n", err)
		}
	}()

	go ap.RunTopicLogger(ctx)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("received signal, graceful shutdown", sig)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ap.Shutdown(ctx)
}
