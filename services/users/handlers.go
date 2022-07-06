package users

import (
	"context"
	"users/gen/pb"
	"users/gen/sqlc"
)

func (s *server) Create(ctx context.Context, in *pb.User) (*pb.User, error) {
	id, err := s.InsertUser(ctx, sqlc.InsertUserParams{
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

func (s *server) List(ctx context.Context, in *pb.Empty) (*pb.Users, error) {
	users, err := s.SelectUsers(ctx)
	if err != nil {
		return nil, err
	}

	out := &pb.Users{}
	for _, user := range users {
		out.Users = append(out.Users, &pb.User{
			Id:          user.ID,
			Name:        user.Name,
			Age:         user.Age,
			Description: user.Description,
		})
	}

	return out, nil
}

func (s *server) Remove(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	err := s.DeleteUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.User) (*pb.User, error) {
	err := s.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:          in.Id,
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
