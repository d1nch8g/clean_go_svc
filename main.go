package main

import (
	"os"
	"users/config"
	"users/postgres"
	"users/services"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Get()
	check(err, `config`)

	pg, err := postgres.Get(postgres.Params{
		ConnString: cfg.PostgresStr,
		MigrDir:    cfg.Migrations,
	})
	check(err, `config`)

	services.Run(services.Params{
		Postgres: pg,
		GrpcPort: cfg.GrpcPort,
	})
	check(err, `services`)
}

func check(err error, module string) {
	if err != nil {
		logrus.Error(`error recieving module`, module, err)
		os.Exit(1)
	}
	logrus.Info(module, ` recieved successfully`)
}
