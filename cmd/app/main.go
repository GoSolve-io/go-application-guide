package main

import (
	"context"
	"fmt"
	"github.com/nglogic/go-application-guide/internal/adapter/metrics"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/caarlos0/env/v6"

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

	metricProvider := metrics.NewDummy(log)

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
		if err := httpgateway.RunServer(ctx, log, metricProvider, srv, conf.HTTPServerAddr); err != nil {
			return fmt.Errorf("http server: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		l, err := net.Listen("tcp", conf.GRPCServerAddr)
		if err != nil {
			return fmt.Errorf("creating net listener: %w", err)
		}
		if err = grpc.RunServer(ctx, log, metricProvider, srv, l); err != nil {
			return fmt.Errorf("grpc server: %w", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Error(err)
	}
}

type config struct {
	HTTPServerAddr string `env:"HTTP_SERVER_ADDR" envDefault:":8080"`
	GRPCServerAddr string `env:"GRPC_SERVER_ADDR" envDefault:":9090"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"info"`

	PostgresDB            string `env:"POSTGRES_DB" envDefault:"testdb"`
	PostgresUser          string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPass          string `env:"POSTGRES_PASS" envDefault:"password"`
	PostgresHostPort      string `env:"POSTGRES_HOSTPORT" envDefault:"localhost:5432"`
	PostgresMigrationsDir string `env:"POSTGRES_MIGRATIONS_DIR" envDefault:"configs/postgresql"`

	MetaweatherAddr    string        `env:"METAWEATHER_ADDR" envDefault:"https://www.metaweather.com"`
	MetaweatherTimeout time.Duration `env:"METAWEATHER_TIMEOUT" envDefault:"10s"`

	BikewiseAddr    string        `env:"BIKEWISE_ADDR" envDefault:"https://bikewise.org/api"`
	BikewiseTimeout time.Duration `env:"BIKEWISE_TIMEOUT" envDefault:"10s"`
}

func newConfig() (config, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("decoding config from env: %w", err)
	}
	return cfg, nil
}
