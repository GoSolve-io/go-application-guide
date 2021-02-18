package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/nglogic/go-example-project/internal/adapter/database"
	"github.com/nglogic/go-example-project/internal/app/bikes"
	"github.com/nglogic/go-example-project/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	dbAdapter, err := database.NewAdapter(
		"postgres://postgres:password@localhost:5432/example?sslmode=disable",
		log,
	)
	if err != nil {
		log.Fatalf("creating bike repository: %v", err)
	}

	bikeService, err := bikes.NewService(dbAdapter.Bikes())
	if err != nil {
		log.Fatalf("creating bike service: %v", err)
	}

	srv, err := server.New(bikeService)
	if err != nil {
		log.Fatalf("creating new server: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	go func() {
		<-sigint
		cancel()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.RunHTTPServer(ctx, log, srv, ":8080")
		cancel()
		if err != nil {
			log.Errorf("http server: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = server.RunGRPCServer(ctx, log, srv, ":9090")
		cancel()
		if err != nil {
			log.Errorf("grpc server: %v", err)
		}
	}()

	wg.Wait()
}
