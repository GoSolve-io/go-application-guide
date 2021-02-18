package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/nglogic/go-example-project/internal/adapter/database"
	"github.com/nglogic/go-example-project/internal/adapter/http/incidents"
	"github.com/nglogic/go-example-project/internal/adapter/http/weather"
	"github.com/nglogic/go-example-project/internal/app/bikes"
	"github.com/nglogic/go-example-project/internal/app/discount"
	"github.com/nglogic/go-example-project/internal/app/reservation"
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

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	weatherAdapter, err := weather.NewAdapter("https://www.metaweather.com", 10*time.Second, httpClient)
	if err != nil {
		log.Fatalf("creating weather adapter: %v", err)
	}

	incidentsAdapter, err := incidents.NewAdapter("https://bikewise.org/api", 10*time.Second, httpClient)
	if err != nil {
		log.Fatalf("creating incidents adapter: %v", err)
	}

	discountService, err := discount.NewService(weatherAdapter, incidentsAdapter)
	if err != nil {
		log.Fatalf("creating discount service: %v", err)
	}

	reservationService, err := reservation.NewService(discountService, bikeService, dbAdapter.Reservations())
	if err != nil {
		log.Fatalf("creating reservation service: %v", err)
	}

	srv, err := server.New(bikeService, reservationService)
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
