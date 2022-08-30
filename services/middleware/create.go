package middleware

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Get() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		errorsInterceptor,
		getLoggingInterceptor(logrus.New()),
		grpc_recovery.UnaryServerInterceptor(),
		grpc_auth.UnaryServerInterceptor(auth),
	)
}
