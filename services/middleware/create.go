package middleware

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Get(logger *logrus.Logger) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		errorsInterceptor,
		ctxTagInterceptor(),
		getLoggingInterceptor(logger),
	)
}
