package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nglogic/go-example-project/internal/adapters/database"
	"github.com/nglogic/go-example-project/internal/app/bikes"
	"github.com/nglogic/go-example-project/internal/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
		err := runHTTPServer(ctx, log, srv, ":8080")
		cancel()
		if err != nil {
			log.Errorf("http server: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = runGRPCServer(ctx, log, srv, ":9090")
		cancel()
		if err != nil {
			log.Errorf("grpc server: %v", err)
		}
	}()

	wg.Wait()
}

func runHTTPServer(
	ctx context.Context,
	log logrus.FieldLogger,
	srv server.ServiceServer,
	addr string,
) error {
	mux := runtime.NewServeMux()

	err := server.RegisterServiceHandlerServer(ctx, mux, srv)
	if err != nil {
		return fmt.Errorf("registering http handlers for server: %w", err)
	}

	// See this great explanation on timeouts:
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	s := http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	go func() {
		<-ctx.Done()
		log.Infof("http server: shutting down")
		if err := s.Shutdown(context.Background()); err != nil {
			log.Errorf("http server: failed to shutdown http gateway server: %v", err)
		}
	}()

	log.Infof("http server: listening on %s", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func runGRPCServer(
	ctx context.Context,
	log logrus.FieldLogger,
	srv server.ServiceServer,
	addr string,
) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("creating net listener: %w", err)
	}

	s := grpc.NewServer()
	server.RegisterServiceServer(s, srv)
	go func() {
		<-ctx.Done()
		log.Infof("grpc server: shutting down")
		s.GracefulStop()
	}()

	log.Infof("grpc server: listening on %s", addr)
	return s.Serve(l)
}
