package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
)

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
