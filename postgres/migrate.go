package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pressly/goose/v3"
)

func Migrate(params Params) error {
	goose.SetLogger(params.Logger)
	connectionString := `postgresql://%s:%s@%s:%d/%s?sslmode=disable`
	connectionString = fmt.Sprintf(
		connectionString,
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.Db,
	)
	var err error
	var db *sql.DB
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			continue
		}
		defer db.Close()
		err = goose.Up(db, params.MigrDir)
		if err != nil {
			params.Logger.Warn(`unable to operate migrations`)
			continue
		}
		time.Sleep(time.Millisecond * 25)
	}
	return err
}
