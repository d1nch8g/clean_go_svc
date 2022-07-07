package container

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
	"users/postgres"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx     = context.Background()
	Logger  = logrus.StandardLogger()
	Timeout = 10

	DbUser = "user"
	DbPass = "password"
	DbHost = "localhost"
	DbPort = rand.Intn(60000)
	DbName = "db"

	TestContainerParams = testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{fmt.Sprintf("%d:5432", DbPort)},
		WaitingFor:   wait.ForLog("system is ready to accept connections"),
		Env: map[string]string{
			"POSTGRES_USER":     DbUser,
			"POSTGRES_PASSWORD": DbPass,
			"POSTGRES_DB":       DbName,
		},
	}
	GenericContainer = testcontainers.GenericContainerRequest{
		ContainerRequest: TestContainerParams,
		Started:          true,
		Logger:           Logger,
	}

	LocalHost  = `localhost`
	DomainHost = `postgres`

	LocalConnectionStr = fmt.Sprintf(
		`postgresql://%s:%s@%s:%d/%s?sslmode=disable`,
		DbUser, DbPass, LocalHost, DbPort, DbName,
	)
	DomainConnectionStr = fmt.Sprintf(
		`postgresql://%s:%s@%s:%d/%s?sslmode=disable`,
		DbUser, DbPass, DomainHost, DbPort, DbName,
	)

	Postgres postgres.IPostgres
)

func init() {
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})

	_, err := testcontainers.GenericContainer(ctx, GenericContainer)
	if err != nil {
		panic(err)
	}

	localDb, err := sql.Open("postgres", LocalConnectionStr)
	if err != nil {
		panic(err)
	}
	domainDb, err := sql.Open("postgres", DomainConnectionStr)
	if err != nil {
		panic(err)
	}

	for {
		err = localDb.Ping()
		if err == nil {
			DbHost = LocalHost
			break
		}

		err = domainDb.Ping()
		if err == nil {
			DbHost = DomainHost
			break
		}

		Logger.Warn(`db ping failed: `, Timeout)
		if Timeout == 0 {
			panic(`db ping failed for too long`)
		}
		Timeout -= 1
		time.Sleep(time.Second)
	}

	Postgres = postgres.New(postgres.Params{
		User:     DbUser,
		Password: DbPass,
		Host:     DbHost,
		Port:     DbPort,
		Db:       DbName,
		Logger:   Logger,
	})
}
