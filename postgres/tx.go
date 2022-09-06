package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
)

func (p *Db) WithTx(tx pgx.Tx) *Db {
	dbtx := p.Queries.WithTx(tx)
	return &Db{
		Queries: *dbtx,
		Pool:    p.Pool,
	}
}

func (p *Db) RollBack(ctx context.Context, tx pgx.Tx) {
	err := tx.Rollback(ctx)
	if err != nil {
		if !errors.Is(err, pgx.ErrTxClosed) {
			log.Printf(`error occured %e`, pgx.ErrTxClosed)
		}
	}
}
