package main

import (
	"context"
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

	// CLose the connection at exit
	defer cc.Close()

	// Establish a new client
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Client activated: %v\n", c)

	// send request to unary client
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Sending the Unary Greet request to server")
	req := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{
		FirstName: "Hendrik",
		LastName:  "Pienaar",
	},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		fmt.Printf("Error while calling Greet RPC: %v\n", err)
	}
	fmt.Println(res.GetResult())
}
