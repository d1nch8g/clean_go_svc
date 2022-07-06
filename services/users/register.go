package users

import (
	"users/services/pb"

	"gitlab.c2g.pw/back/modelrepo/postgres"
)

type server struct {
	pb.UnimplementedUsersServer
	postgres.IPostgres
}

var serverInstance server

func init() {
	
}
