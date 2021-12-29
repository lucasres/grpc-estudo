package cmd

import (
	"log"
	"net"

	"github.com/lucasres/grpc-estudo/pb/pb"
	"github.com/lucasres/grpc-estudo/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func GrpcServer() {
	lis, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("Cannot list %v", err)
	}

	grpcServer := grpc.NewServer()
	// registrando os services
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	// habilitando o reflection
	reflection.Register(grpcServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot server: %v", err)
	}
}
