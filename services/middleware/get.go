package middleware

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

func Get() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		getLoggingInterceptor(),
		grpc_auth.UnaryServerInterceptor(auth),
		grpc_recovery.UnaryServerInterceptor(),
		errorsInterceptor,
	)
}
