package main

import (
	"users/goose"
	"users/postgres"
	"users/services"

	"github.com/caarlos0/env/v6"

	"github.com/sirupsen/logrus"
)

var (
	cfg    = Config{}
	logger = logrus.StandardLogger()
)

type Config struct {
	PostgresUser     string `env:"POSTGRES_USER"     envDefault:"user"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	PostgresHost     string `env:"POSTGRES_HOST"     envDefault:"localhost"`
	PostgresPort     int    `env:"POSTGRES_PORT"     envDefault:"5432"`
	PostgresDb       string `env:"POSTGRES_DB"       envDefault:"db"`
	GrpcPort         int    `env:"GRPC_PORT"         envDefault:"9080"`
	HttpPort         int    `env:"HTTP_PORT"         envDefault:"8080"`
	JsonLogs         bool   `env:"JSON_LOGS"         envDefault:"true"`
}

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	logger.SetFormatter(&logrus.JSONFormatter{})
	if !cfg.JsonLogs {
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:  true,
			DisableQuote: true,
		})
	}
	logger.SetLevel(logrus.DebugLevel)
}

func main() {
	goose.Migrate(goose.Params{
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		Db:       cfg.PostgresDb,
		Dir:      "migrations",
		Logger:   logger,
	})

	database := postgres.New(postgres.Params{
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
		Postgres: database,
		Logger:   logger,
	})
}
