package config

import (
	"github.com/caarlos0/env/v6"
)

type config struct {
	PostgresStr string `env:"POSTGRES_CONN_STRING" envDefault:"postgresql://user:password@host.docker.internal:5432/db"`
	Migrations  string `env:"MIGRATIONS_DIR"       envDefault:"migrations"`
	GrpcPort    int    `env:"GRPC_PORT"            envDefault:"9080"`
	JsonLogs    bool   `env:"JSON_LOGS"            envDefault:"true"`
}

func Get() (*config, error) {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
