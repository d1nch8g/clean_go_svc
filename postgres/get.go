package postgres

import (
	"context"
	"users/postgres/sqlc"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

type Params struct {
	ConnString string
	MigrDir    string
}

type Db struct {
	*pgxpool.Pool
	params Params
	sqlc.Queries
}

func Get(params Params) (*Db, error) {
	goose.SetLogger(logrus.StandardLogger())
	migrDb, err := goose.OpenDBWithDriver(`pgx`, params.ConnString+`?sslmode=disable`)
	if err != nil {
		return nil, err
	}
	defer migrDb.Close()
	err = goose.Up(migrDb, params.MigrDir)
	if err != nil {
		return nil, err
	}

	config, err := pgxpool.ParseConfig(params.ConnString)
	if err != nil {
		return nil, err
	}
	config.ConnConfig.LogLevel = pgx.LogLevelError
	config.ConnConfig.Logger = logrusadapter.NewLogger(logrus.StandardLogger())
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	sqlc := sqlc.New(pool)
	return &Db{
		Queries: *sqlc,
		Pool:    pool,
		params:  params,
	}, nil
}
