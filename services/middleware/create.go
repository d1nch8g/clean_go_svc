package middleware

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Get(logger *logrus.Logger) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		errorsInterceptor,
		getLoggingInterceptor(logger),
		grpc_recovery.UnaryServerInterceptor(),
	)
}
