package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type config struct {
	HTTPServerAddr string `env:"HTTP_SERVER_ADDR" envDefault:":8080"`
	GRPCServerAddr string `env:"GRPC_SERVER_ADDR" envDefault:":9090"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"info"`

	PostgresDSN string `env:"POSTGRES_DSN" envDefault:"postgres://postgres:password@localhost:5432/example?sslmode=disable"`

	MetaweatherAddr    string        `env:"METAWEATHER_ADDR" envDefault:"https://www.metaweather.com"`
	MetaweatherTimeout time.Duration `env:"METAWEATHER_TIMEOUT" envDefault:"10s"`

	BikewiseAddr    string        `env:"BIKEWISE_ADDR" envDefault:"https://bikewise.org/api"`
	BikewiseTimeout time.Duration `env:"BIKEWISE_TIMEOUT" envDefault:"10s"`
}

func newConfig() (config, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("decoding config from env: %v", err)
	}
	return cfg, nil
}
