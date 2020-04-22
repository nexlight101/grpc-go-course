package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nexlight101/gRPC_course/greet/greetpb"
	"google.golang.org/grpc"
)

// Create a server type
type server struct{}

// Implement unary server function
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Request received in server")
	fullName := req.Greeting.GetFirstName() + " " + req.Greeting.GetLastName()
	result := "Hello " + fullName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello from gRPC-Server")

	// Create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//Create a new gRPC server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	// Check if the server is serving the listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
