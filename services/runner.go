package services

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"users/gen/go/pb"
	"users/postgres"
	"users/services/middleware"
	"users/services/users"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Params struct {
	GrpcPort int
	HttpPort int
	Postgres postgres.IPostgres
	Logger   *logrus.Logger
}

func Run(params Params) {
	s := grpc.NewServer(middleware.Get(params.Logger))

	users.Register(s, params.Postgres)

	go runGrpc(params.GrpcPort, s)
	runHttp(params.HttpPort, params.GrpcPort)
}

func runGrpc(port int, s *grpc.Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logrus.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func runHttp(httpport int, grpcport int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	grpcadress := fmt.Sprintf("localhost:%d", grpcport)
	err := pb.RegisterUserStorageHandlerFromEndpoint(ctx, mux, grpcadress, opts)
	if err != nil {
		panic(err)
	}

	logrus.Printf("http server listening at %d", httpport)
	httpPort := fmt.Sprintf(":%d", httpport)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		panic(err)
	}
}
