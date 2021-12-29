package services

import (
	"context"
	"fmt"
	"time"

	"github.com/lucasres/grpc-estudo/pb/pb"
)

// AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUser(ctx context.Context, u *pb.User) (*pb.User, error) {
	fmt.Println("Recived request")
	return &pb.User{
		Id:    u.GetId(),
		Name:  u.GetName(),
		Email: u.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(u *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	fmt.Println("Recived request verbose")
	stream.Send(&pb.UserResultStream{
		Status: "init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "prepare_insert_db",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "inserted_db",
		User: &pb.User{
			Id:    u.GetId(),
			Name:  u.GetName(),
			Email: u.GetName(),
		},
	})

	return nil
}

func NewUserService() *UserService {
	return &UserService{}
}
