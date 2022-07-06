package users

import (
	"users/postgres"
)

type server struct {
	pb.UnimplementedUsersServer
	postgres.IPostgres
}

var serverInstance server

func init() {

}
