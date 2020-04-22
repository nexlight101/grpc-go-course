package main

import (
	"fmt"
	"log"

	"github.com/nexlight101/gRPC_course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	// Create connection to the server
	options := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Client activated: %f\n", c)
}
