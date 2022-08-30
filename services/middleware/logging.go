package middleware

import (
	"time"

	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getLoggingInterceptor() grpc.UnaryServerInterceptor {
	entry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc_logrus.Option{
		grpc_logrus.WithCodes(grpc_logging.DefaultErrorToCode),
		grpc_logrus.WithLevels(grpc_logrus.DefaultClientCodeToLevel),
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
		grpc_logrus.WithMessageProducer(grpc_logrus.DefaultMessageProducer),
	}
	grpc_logrus.ReplaceGrpcLogger(entry)
	return grpc_logrus.UnaryServerInterceptor(entry, opts...)
}
