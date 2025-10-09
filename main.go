package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lao-tseu-is-alive/go-cloud-grpc-client-ai/pkg/inference" // Change to your module name
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	c := pb.NewInferencerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	prompt := "Hello, my name is Carlos ! What's yours ?"
	log.Printf("Sending prompt to Python service: %s", prompt)

	r, err := c.GenerateText(ctx, &pb.GenerateRequest{Prompt: prompt})
	if err != nil {
		log.Fatalf("could not generate text: %v", err)
	}

	log.Printf("Received from Python gRPC server:\n%s\n", r.GetGeneratedText())
}
