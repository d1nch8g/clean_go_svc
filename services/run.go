package services

import (
	"fmt"
	"net"
	"users/postgres"
	"users/services/middleware"
	"users/services/pb"
	"users/services/users"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	GrpcPort int
	Postgres postgres.IPostgres
}

func Run(params Params) error {
	s := grpc.NewServer(middleware.GetUnary())
	srv := users.Server{Pg: params.Postgres}
	pb.RegisterUserStorageServer(s, srv)
	reflection.Register(s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", params.GrpcPort))
	if err != nil {
		return err
	}
	logrus.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}
