package users

import (
	"context"
	"users/postgres"
	"users/postgres/sqlc"
	"users/services/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUnknown = status.Error(codes.NotFound, `unknown error`)

type Server struct {
	pb.UnimplementedUserStorageServer
	Pg *postgres.Db
}

func (s Server) Create(ctx context.Context, in *pb.User) (*pb.User, error) {
	id, err := s.Pg.InsertUser(ctx, sqlc.InsertUserParams{
		Description: in.Description,
		Name:        in.Name,
		Age:         in.Age,
	})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:          id,
		Name:        in.Name,
		Age:         in.Age,
		Description: in.Description,
	}, nil
}

func (s Server) List(in *pb.Empty, str pb.UserStorage_ListServer) error {
	users, err := s.Pg.SelectUsers(str.Context())
	if err != nil {
		return ErrUnknown
	}
	for _, sur := range users {
		str.Send(&pb.User{
			Id:          sur.Id,
			Name:        sur.Name,
			Age:         sur.Age,
			Description: sur.Description,
		})
	}
	return nil
}

func (s Server) Remove(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	err := s.Pg.DeleteUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s Server) Update(ctx context.Context, in *pb.User) (*pb.User, error) {
	err := s.Pg.UpdateUser(ctx, sqlc.UpdateUserParams{
		Id:          in.Id,
		Name:        in.Name,
		Age:         in.Age,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:          in.Id,
		Name:        in.Name,
		Age:         in.Age,
		Description: in.Description,
	}, nil
}
