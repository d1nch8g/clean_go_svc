package main

import (
	"os"
	"users/config"
	"users/postgres"
	"users/services"

	"github.com/sirupsen/logrus"
)

func main() {
	database, err := postgres.New(postgres.Params{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		Db:       config.PostgresDb,
		Logger:   config.Logger,
	})
	if err != nil {
		logrus.Panic(err)
		os.Exit(1)
	}

	services.Run(services.Params{
		GrpcPort: config.GrpcPort,
		HttpPort: config.HttpPort,
		Postgres: database,
		Logger:   config.Logger,
	})
}
