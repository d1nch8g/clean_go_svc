package users

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"users/config"
	"users/gen/go/pb"
	"users/postgres"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
	pg  postgres.IPostgres
	s   server
)

func init() {
	testpg, err := postgres.New(postgres.Params{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		Db:       config.PostgresDb,
		MigrDir:  `../../migrations`,
		Logger:   config.Logger,
	})
	if err != nil {
		panic(err)
	}
	pg = testpg
	s = server{IPostgres: pg}
}

func TestCreate(t *testing.T) {
	randName := fmt.Sprintf("testname_%d", rand.Intn(1000000))
	resp, err := s.Create(ctx, &pb.User{
		Name:        randName,
		Age:         25,
		Description: "test_descr",
	})
	assert.Nil(t, err)
	assert.NotEqual(t, 0, resp.Id)
	users, err := pg.SelectUsers(ctx)
	assert.Nil(t, err)
	found := false
	for _, user := range users {
		if user.Name == randName {
			found = true
		}
	}
	assert.True(t, found)
}

// TODO add your tests here
