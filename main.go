package main

import (
	"users/postgres/migr"

	"users/postgres"
	"users/services"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

var cfg = Config{}

type Config struct {
	PostgresUser     string `env:"POSTGRES_USER"     envDefault:"user"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	PostgresHost     string `env:"POSTGRES_HOST"     envDefault:"localhost"`
	PostgresPort     int    `env:"POSTGRES_PORT"     envDefault:"5432"`
	PostgresDb       string `env:"POSTGRES_DB"       envDefault:"db"`
	GrpcPort         int    `env:"GRPC_PORT"         envDefault:"9080"`
	HttpPort         int    `env:"HTTP_PORT"         envDefault:"8080"`
}

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
}

func main() {
	logger := logrus.StandardLogger()

	migr.Migrate(migr.Params{
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		Db:       cfg.PostgresDb,
		Dir:      "migrations",
	})

	db := postgres.New(postgres.Params{
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		Db:       cfg.PostgresDb,
		Logger:   logger,
	})

	services.Run(services.Params{
		GrpcPort: cfg.GrpcPort,
		HttpPort: cfg.HttpPort,
		Postgres: db,
		Logger:   logger,
	})
}
