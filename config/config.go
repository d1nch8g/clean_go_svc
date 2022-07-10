package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type config struct {
	PostgresUser     string `env:"POSTGRES_USER"     envDefault:"user"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	PostgresHost     string `env:"POSTGRES_HOST"     envDefault:"localhost"`
	PostgresPort     int    `env:"POSTGRES_PORT"     envDefault:"5432"`
	PostgresDb       string `env:"POSTGRES_DB"       envDefault:"db"`
	GrpcPort         int    `env:"GRPC_PORT"         envDefault:"9080"`
	HttpPort         int    `env:"HTTP_PORT"         envDefault:"8080"`
	JsonLogs         bool   `env:"JSON_LOGS"         envDefault:"true"`
}

var (
	Logger = logrus.StandardLogger()

	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     int
	PostgresDb       string
	GrpcPort         int
	HttpPort         int
	JsonLogs         bool
)

func init() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	PostgresUser = cfg.PostgresUser
	PostgresPassword = cfg.PostgresPassword
	PostgresHost = cfg.PostgresHost
	PostgresPort = cfg.PostgresPort
	PostgresDb = cfg.PostgresDb
	GrpcPort = cfg.GrpcPort
	HttpPort = cfg.HttpPort
	JsonLogs = cfg.JsonLogs

	Logger.SetFormatter(&logrus.JSONFormatter{})
	if !JsonLogs {
		Logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:  true,
			DisableQuote: true,
		})
	}
	Logger.SetLevel(logrus.DebugLevel)
}
