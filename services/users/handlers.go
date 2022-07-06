package users

import (
	"context"
	"users/gen/pb"
	"users/gen/sqlc"
)

func (s *server) Create(ctx context.Context, in *pb.User) (*pb.User, error) {
	s.InsertUser(ctx, sqlc.InsertUserParams{})
	return &pb.User{Id: in.Id}, nil
}

func (s *server) Get(ctx context.Context, in *pb.Id) (*pb.User, error) {
	return &pb.User{Id: in.Id}, nil
}

func (s *server) Remove(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.User) (*pb.User, error) {
	return &pb.User{Id: in.Id}, nil
}
