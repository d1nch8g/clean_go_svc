package users

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"users/gen/go/pb"
	"users/postgres/container"

	"github.com/stretchr/testify/assert"
)

var (
	ctx        = context.Background()
	pg         = container.Postgres
	testServer = server{
		IPostgres: container.Postgres,
	}
)

func TestCreate(t *testing.T) {
	randName := fmt.Sprintf("testname_%d", rand.Intn(1000000))
	resp, err := testServer.Create(ctx, &pb.User{
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
