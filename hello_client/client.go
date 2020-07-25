package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fenriz07/grpc-boilerplate/hello_client/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client run")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err.Error())
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)

	helloUnary(c)

}

func helloUnary(c hellopb.HelloServiceClient) {

	fmt.Println("Starting unary grpc Hello")

	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{FirstName: "Servio", Prefix: "Sr"},
	}

	res, err := c.Hello(context.Background(), req)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Response Hello %v", res.CustomHello)
}
