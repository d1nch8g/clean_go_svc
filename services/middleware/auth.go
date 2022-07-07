package middleware

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	dummyToken = "12345"
)

var ErrAuthFailed = status.Errorf(codes.PermissionDenied, "buildDummyAuthFunction bad token")

func auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	if token != dummyToken {
		return nil, ErrAuthFailed
	}
	return ctx, nil
}
