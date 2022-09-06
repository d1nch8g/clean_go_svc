package users

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"users/config"
	"users/postgres"
	"users/postgres/sqlc"
	"users/services/pb"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var (
	ctx = context.Background()
	pg  = getPg()
	s   = Server{Pg: pg}
)

func getPg() *postgres.Db {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}
	pg, err := postgres.Get(postgres.Params{
		ConnString: cfg.PostgresStr,
		MigrDir:    "../../postgres/migrations",
	})
	if err != nil {
		panic(err)
	}
	return pg
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

type mockStream struct {
	grpc.ServerStream
	users []*pb.User
}

func (f *mockStream) Send(in *pb.User) error {
	f.users = append(f.users, in)
	return nil
}

func (f *mockStream) Context() context.Context {
	return ctx
}

func TestList(t *testing.T) {
	stream := &mockStream{}
	pg.InsertUser(ctx, sqlc.InsertUserParams{
		Name: `test name`,
	})
	err := s.List(&pb.Empty{}, stream)
	assert.Nil(t, err)
	found := false
	for _, u := range stream.users {
		if u.Name == `test name` {
			found = true
		}
	}
	assert.True(t, found)
}
