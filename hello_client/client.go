package main

import (
	"context"
	"fmt"
	"io"
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

	//helloUnary(c)
	helloServerStreaming(c)

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

func helloServerStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("invoked helloServerStreaming")

	req := &hellopb.HelloManyLanguagesRequest{
		Hello: &hellopb.Hello{
			FirstName: "Servio",
			Prefix:    "Don",
		},
	}

	restStream, err := c.HelloManyLanguages(context.Background(), req)

	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		msg, err := restStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error reading stream")
		}

		log.Printf("Res from HML: %v \n", msg.GetHelloLanguage())
	}
}
