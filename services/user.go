package services

import (
	"context"

	"github.com/lucasres/grpc-estudo/pb/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, u *pb.User) (*pb.User, error) {
	return &pb.User{
		Id:    u.GetId(),
		Name:  u.GetName(),
		Email: u.GetEmail(),
	}, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
