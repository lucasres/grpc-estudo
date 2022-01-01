package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1 - Unary\n 2 - Response stream\n 3 - Sent with stream")

	choose, _ := reader.ReadString('\n')

	if choose == "1\n" {
		callAddUser(client)
	} else if choose == "2\n" {
		callAddUserVerbose(client)
	} else {
		callSendStreamUser(client)
	}

}

func callAddUser(c pb.UserServiceClient) {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Id:")
	id, _ := r.ReadString('\n')
	fmt.Println("Name:")
	name, _ := r.ReadString('\n')
	fmt.Println("Email:")
	email, _ := r.ReadString('\n')

	u := &pb.User{
		Id:    id,
		Name:  name,
		Email: email,
	}

	res, err := c.AddUser(context.Background(), u)

	if err != nil {
		log.Fatalf("Cannot make request: %v", err)
	}

	fmt.Printf("Response:\n Id: %s\n Name: %s\n Email: %s", res.GetId(), res.GetName(), res.GetEmail())
}

func callAddUserVerbose(c pb.UserServiceClient) {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Id:")
	id, _ := r.ReadString('\n')
	fmt.Println("Name:")
	name, _ := r.ReadString('\n')
	fmt.Println("Email:")
	email, _ := r.ReadString('\n')

	u := &pb.User{
		Id:    id,
		Name:  name,
		Email: email,
	}

	res, err := c.AddUserVerbose(context.Background(), u)

	if err != nil {
		log.Fatalf("Cannot make request: %v", err)
	}

	for {
		stream, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Cannot recive: %v", err)
		}

		fmt.Println(stream.GetStatus())
	}
}

func callSendStreamUser(c pb.UserServiceClient) {
	users := []*pb.User{
		&pb.User{
			Id:    "123",
			Name:  "Lucas",
			Email: "lucas@email.com",
		},
		&pb.User{
			Id:    "456",
			Name:  "Everson Zoio",
			Email: "777fazsol@email.com",
		},
		&pb.User{
			Id:    "789",
			Name:  "Matue",
			Email: "voufazerumamaquinadotempo@email.com",
		},
	}

	stream, err := c.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("erro in send stream client %v", err)
	}

	for _, val := range users {
		stream.Send(val)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error when close stream %v", err)
	}

	fmt.Println("response:")
	fmt.Print(res)
}
