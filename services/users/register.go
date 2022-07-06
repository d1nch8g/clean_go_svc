package users

import (
	"users/gen/go/pb"
	"users/postgres"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserStorageServer
	postgres.IPostgres
}

func Register(s grpc.ServiceRegistrar, db postgres.IPostgres) {
	if db == nil {
		panic(`attempt to register users server with nil db`)
	}

	pb.RegisterUserStorageServer(s, &server{
		IPostgres: db,
	})
}
