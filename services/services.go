package services

import (
	"fmt"
	"log"
	"net"
	"users/gen/pb"
	"users/postgres"
	"users/services/middleware"
	"users/services/users"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Params struct {
	GrpcPort int
	Postgres postgres.IPostgres
}

func Run(params Params) {
	s := grpc.NewServer(middleware.Get())
	srv := users.Server{Pg: params.Postgres}
	pb.RegisterUserStorageServer(s, srv)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", params.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	logrus.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
