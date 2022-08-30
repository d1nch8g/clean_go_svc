package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type config struct {
	PostgresStr string `env:"POSTGRES_CONN_STRING" envDefault:"postgresql://user:password@host.docker.internal:5432/db"`
	Migrations  string `env:"MIGRATIONS_DIR"       envDefault:"postgres/migrations"`
	JsonLogs    bool   `env:"JSON_LOGS"            envDefault:"true"`
	GrpcPort    int    `env:"GRPC_PORT"            envDefault:"9080"`
}

func Get() (*config, error) {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	log := logrus.StandardLogger()
	if cfg.JsonLogs {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}
	return &cfg, nil
}
