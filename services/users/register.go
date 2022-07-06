package users

import (
	"users/gen/pb"
	"users/postgres"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUsersServer
	postgres.IPostgres
}

func Register(s grpc.ServiceRegistrar, db postgres.IPostgres) {
	if db == nil {
		panic(`attempt to register users server with nil db`)
	}
	pb.RegisterUsersServer(s, &server{
		IPostgres: db,
	})
}
