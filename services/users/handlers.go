package users

import (
	"context"
	"users/services/pb"
)

func (s *server) Create(ctx context.Context, in *pb.User) (*pb.User, error) {
	return &pb.User{Id: in.Id}, nil
}

func (s *server) Get(ctx context.Context, in *pb.User) (*pb.User, error) {
	return &pb.User{Id: in.Id}, nil
}

func (s *server) Remove(ctx context.Context, in *pb.User) (*pb.User, error) {
	return &pb.User{Id: in.Id}, nil
}

func (s *server) Update(ctx context.Context, in *pb.User) (*pb.User, error) {
	return &pb.User{Id: in.Id}, nil
}
