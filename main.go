package main

import (
	"users/config"
	"users/goose"
	"users/postgres"
	"users/services"
)

func main() {
	goose.Migrate(goose.Params{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		Db:       config.PostgresDb,
		Dir:      "migrations",
		Logger:   config.Logger,
	})

	database := postgres.New(postgres.Params{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		Db:       config.PostgresDb,
		Logger:   config.Logger,
	})

	services.Run(services.Params{
		GrpcPort: config.GrpcPort,
		HttpPort: config.HttpPort,
		Postgres: database,
		Logger:   config.Logger,
	})
}
