package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/lucasres/grpc-estudo/pb/pb"
	"google.golang.org/grpc"
)

func GrpcClient() {
	con, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot connect to server: %v", err)
	}

	defer con.Close()

	client := pb.NewUserServiceClient(con)

	callAddUser(client)

}

func callAddUser(c pb.UserServiceClient) {
	u := &pb.User{
		Id:    "123456",
		Name:  "Everson Zoio",
		Email: "azideia@email.com",
	}

	res, err := c.AddUser(context.Background(), u)

	if err != nil {
		log.Fatalf("Cannot make request: %v", err)
	}

	fmt.Printf("Response: Id: %s, Name: %s, Email: %s", res.GetId(), res.GetName(), res.GetEmail())
}
