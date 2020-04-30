package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/nexlight101/gRPC_course/greet/greetpb"
	"google.golang.org/grpc"
)

// Create a server type
type server struct{}

// LongGreet client stream
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Greet stream Request received in server")
	result := ""
	for {
		req, rErr := stream.Recv()
		if rErr == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if rErr != nil {
			log.Fatalf("Could not receive from client: %v\n", rErr)
			return rErr
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

// Implement unary server function
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Request received in server %v", req)
	fullName := req.Greeting.GetFirstName() + " " + req.Greeting.GetLastName()
	result := "Hello " + fullName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

// Implement server streaming function
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet Request received in GreetManyTimes server %v", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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
