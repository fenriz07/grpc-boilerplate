package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fenriz07/grpc-boilerplate/hello_server/hellopb"
	"google.golang.org/grpc"
)

type server struct {
	hellopb.UnimplementedHelloServiceServer
}

func (*server) Hello(cxt context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Println("Hello function was called with %v \n", req)

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	customHello := "Bienvenido " + prefix + " " + firstName

	res := &hellopb.HelloResponse{
		CustomHello: customHello,
	}

	return res, nil

}

func main() {
	fmt.Println("Iniciando el server go")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
