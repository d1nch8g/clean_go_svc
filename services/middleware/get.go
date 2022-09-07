package middleware

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

func GetUnary() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		getUnaryLogger(),
		grpc_auth.UnaryServerInterceptor(auth),
		grpc_recovery.UnaryServerInterceptor(),
	)
}

func GetStream() grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		getStreamLogger(),
		grpc_auth.StreamServerInterceptor(auth),
		grpc_recovery.StreamServerInterceptor(),
	)
}
