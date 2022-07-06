package migr

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Params struct {
	User     string
	Password string
	Host     string
	Port     int
	Db       string
	Dir      string
}

var ErrMissingMigrations = errors.New(`migrations are outdated`)

func Migrate(params Params) {
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
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Millisecond * 500)
		if i == 9 {
			panic(`unable to connect to database`)
		}
	}
	err = goose.Up(db, params.Dir)
	if err != nil {
		panic(fmt.Errorf(`migrate err: %s`, err))
	}
}
