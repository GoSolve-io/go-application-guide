package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nglogic/go-application-guide/internal/adapter/database"
	"github.com/nglogic/go-application-guide/internal/adapter/http/incidents"
	"github.com/nglogic/go-application-guide/internal/adapter/http/weather"
	"github.com/nglogic/go-application-guide/internal/app/bikerental/bikes"
	"github.com/nglogic/go-application-guide/internal/app/bikerental/discount"
	"github.com/nglogic/go-application-guide/internal/app/bikerental/reservation"
	"github.com/nglogic/go-application-guide/internal/transport/grpc"
	"github.com/nglogic/go-application-guide/internal/transport/grpc/httpgateway"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const (
	maxHTTPClientTimeout = 30 * time.Second
)

func main() {
	log := logrus.New()

	conf, err := newConfig()
	if err != nil {
		log.Fatalf("initializing config: %v", err)
	}

	logLevel, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		log.Fatalf("initializing config: %v", err)
	}
	log.SetLevel(logLevel)

	dbAdapter, err := database.NewAdapter(conf.PostgresHostPort, conf.PostgresDB, conf.PostgresUser, conf.PostgresPass, conf.PostgresMigrationsDir, log)
	if err != nil {
		log.Fatalf("creating bike repository: %v", err)
	}

	bikeService, err := bikes.NewService(dbAdapter.Bikes())
	if err != nil {
		log.Fatalf("creating bike service: %v", err)
	}

	httpClient := &http.Client{
		Timeout: maxHTTPClientTimeout,
	}

	weatherAdapter, err := weather.NewAdapter(conf.MetaweatherAddr, conf.MetaweatherTimeout, httpClient)
	if err != nil {
		log.Fatalf("creating weather adapter: %v", err)
	}

	incidentsAdapter, err := incidents.NewAdapter(conf.BikewiseAddr, conf.BikewiseTimeout, httpClient)
	if err != nil {
		log.Fatalf("creating incidents adapter: %v", err)
	}

	discountService, err := discount.NewService(weatherAdapter, incidentsAdapter)
	if err != nil {
		log.Fatalf("creating discount service: %v", err)
	}

	reservationService, err := reservation.NewService(
		discountService,
		bikeService,
		dbAdapter.Reservations(),
		dbAdapter.Customers(),
	)
	if err != nil {
		log.Fatalf("creating reservation service: %v", err)
	}

	srv, err := grpc.NewServer(bikeService, reservationService, log)
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

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := httpgateway.RunServer(ctx, log, srv, conf.HTTPServerAddr); err != nil {
			return fmt.Errorf("http server: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		l, err := net.Listen("tcp", conf.GRPCServerAddr)
		if err != nil {
			return fmt.Errorf("creating net listener: %w", err)
		}
		if err = grpc.RunServer(ctx, log, srv, l); err != nil {
			return fmt.Errorf("grpc server: %w", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Error(err)
	}
}
