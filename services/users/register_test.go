package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestRegisterSuccess(t *testing.T) {
	assert.NotPanics(t, func() {
		s := grpc.NewServer()
		Register(s, pg)
	})
}

func TestRegisterFail(t *testing.T) {
	assert.Panics(t, func() {
		s := grpc.NewServer()
		Register(s, nil)
	})
}
