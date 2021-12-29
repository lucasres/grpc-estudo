package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/lucasres/grpc-estudo/cmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose:\n 1 - server\n 2 - client\n")

	choose, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Cannot read choose: %v", err)
	}

	if choose == "1\n" {
		fmt.Println("Start server at :5000")
		cmd.GrpcServer()
	} else if choose == "2\n" {
		cmd.GrpcClient()
	} else {
		log.Fatalf("Invalid choose")
	}
}
