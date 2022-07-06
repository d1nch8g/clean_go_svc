package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"users/gen/sqlc"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IPostgres interface {
	sqlc.Querier

	WithTx(tx pgx.Tx) IPostgres
	Begin(ctx context.Context) (pgx.Tx, error)
	RollBack(ctx context.Context, tx pgx.Tx)
}

type Params struct {
	User     string
	Password string
	Host     string
	Port     int
	Db       string
}

type postgres struct {
	*pgxpool.Pool
	params Params
	sqlc.Queries
}

func New(params Params) IPostgres {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		`postgresql://%s:%s@%s:%d/%s`,
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.Db,
	))
	if err != nil {
		panic(fmt.Errorf(`unable to parse postgres connection string, %e`, err))
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(fmt.Errorf(`unable to connect pool, %e`, err))
	}
	sqlc := sqlc.New(pool)
	pg := &postgres{
		Queries: *sqlc,
		Pool:    pool,
		params:  params,
	}

	if err != nil {
		panic(fmt.Errorf(`unable to set postgres constants, %e`, err))
	}

	return pg
}

func (p *postgres) WithTx(tx pgx.Tx) IPostgres {
	dbtx := p.Queries.WithTx(tx)
	return &postgres{
		Queries: *dbtx,
		Pool:    p.Pool,
	}
}

func (p *postgres) RollBack(ctx context.Context, tx pgx.Tx) {
	err := tx.Rollback(ctx)
	if err != nil {
		if !errors.Is(err, pgx.ErrTxClosed) {
			log.Printf(`error occured %e`, pgx.ErrTxClosed)
		}
	}
}
