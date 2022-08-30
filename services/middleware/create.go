package middleware

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

func Get() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		errorsInterceptor,
		getLoggingInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
		grpc_auth.UnaryServerInterceptor(auth),
	)
}
