package users

import (
	"testing"
	"users/postgres/container"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestRegisterSuccess(t *testing.T) {
	assert.NotPanics(t, func() {
		s := grpc.NewServer()
		Register(s, container.Postgres)
	})
}

func TestRegisterFail(t *testing.T) {
	assert.Panics(t, func() {
		s := grpc.NewServer()
		Register(s, nil)
	})
}
