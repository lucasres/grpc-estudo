package services

import (
	"context"
	"fmt"
	"io"
	"log"
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

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Error when recive stream msg %v", err)
		}

		fmt.Printf("Adding new user: %s\n", req.GetId())

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error when recive stream in bi directional stream  %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added user",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Error when send stream in bi directional stream %v", err)
		}
	}

}

func NewUserService() *UserService {
	return &UserService{}
}
