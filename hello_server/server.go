package main

import (
	"fmt"
	"log"
	"net"

	"../hellopb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Hello go server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	s := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}

}
