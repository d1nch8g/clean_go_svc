package middleware

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknown  = status.Error(codes.Unknown, `unknown error`)
	ErrNotFound = status.Error(codes.NotFound, `not found in database`)

	// TODO add your errors and cases checks for them  here
)

func errorsInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, ErrUnknown
	}
	return resp, nil
}
