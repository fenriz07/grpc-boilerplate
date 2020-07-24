package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main()  {
	fmt.Println("Go client run")	

	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil{
		log.Fatal(err.Error())
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)
}