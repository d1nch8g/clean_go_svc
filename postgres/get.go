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

type IPostgres interface {
	sqlc.Querier

	WithTx(tx pgx.Tx) IPostgres
	Begin(ctx context.Context) (pgx.Tx, error)
	RollBack(ctx context.Context, tx pgx.Tx)
}

type Params struct {
	ConnString string
	MigrDir    string
}

type postgres struct {
	*pgxpool.Pool
	params Params
	sqlc.Queries
}

func Get(params Params) (IPostgres, error) {
	goose.SetLogger(logrus.StandardLogger())
	db, err := goose.OpenDBWithDriver(`pgx`, params.ConnString+`?sslmode=disable`)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = goose.Up(db, params.MigrDir)
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
	pg := &postgres{
		Queries: *sqlc,
		Pool:    pool,
		params:  params,
	}
	if err != nil {
		return nil, err
	}
	return pg, nil
}
