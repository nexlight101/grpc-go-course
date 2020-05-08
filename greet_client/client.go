package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/nexlight101/gRPC_course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// doUnary(c)

	// doServerStreaming
	// doServerStreaming(c)

	// doStreamingClient
	// doClientStreaming(c)

	// doUnaryWithDeadline implement deadline
	doUnaryWithDeadline(c, 5*time.Second) // Should complete
	doUnaryWithDeadline(c, 1*time.Second) // Should timeout

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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Sending the streaming Greet request to server")
	req := &greetpb.GreetManyTimesRequest{Greeting: &greetpb.Greeting{
		FirstName: "Hendrik",
		LastName:  "Pienaar",
	},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v\n", err)
	}

	for {
		msg, streamErr := resStream.Recv()
		if streamErr == io.EOF {
			break
		}
		if streamErr != nil {
			log.Fatalf("Error while reading stream: %v\n", streamErr)
		}
		log.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}

// Client streaming function
func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Sending the streaming Greet request to server")
	requests := []*greetpb.LongGreetRequest{{
		Greeting: &greetpb.Greeting{
			FirstName: "Hendrik",
			LastName:  "Pienaar",
		},
	},
		{Greeting: &greetpb.Greeting{
			FirstName: "Henriette",
			LastName:  "Pienaar",
		},
		},
		{Greeting: &greetpb.Greeting{
			FirstName: "Danielle",
			LastName:  "Pienaar",
		},
		},
		{Greeting: &greetpb.Greeting{
			FirstName: "Michele",
			LastName:  "Pienaar",
		},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Cannot send to server: %v\n", err)
	}
	for _, v := range requests {
		sErr := stream.Send(v)
		if sErr != nil {
			log.Fatalf("Cannot send request: %v\n", sErr)
		}
	}
	res, serErr := stream.CloseAndRecv()
	if serErr != nil {
		log.Fatalf("Cannot receive from server: %v\n", serErr)
	}
	fmt.Printf(" %v\n ", res.GetResult())
}

// doUnaryWithDeadline
func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Sending the Unary GreetWithDeadline request to server")
	req := &greetpb.GreetWithDeadlineRequest{Greeting: &greetpb.Greeting{
		FirstName: "Hendrik",
		LastName:  "Pienaar",
	},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)

	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("Unxpected error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error while calling GreetWithDeadline RPC: %v\n", err)
		}
		return
	}
	fmt.Printf("Response from GreeWithDeadline: %v\n", res.GetResult())
}
