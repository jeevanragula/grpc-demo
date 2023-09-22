package main

import (
	pb "client/task/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	serverAddress := "localhost:5000"

	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewTaskServiceClient(conn)

	// Example: Create a Task
	createTaskRequest := &pb.CreateTaskRequest{
		Title:       "Sample Task",
		Description: "This is a sample task.",
	}
	createTaskResponse, err := client.CreateTask(context.Background(), createTaskRequest)
	if err != nil {
		log.Fatalf("CreateTask failed: %v", err)
	}
	fmt.Printf("Created Task: %v\n", createTaskResponse)

	// Example: Get a Task by ID
	getTaskRequest := &pb.GetTaskRequest{
		Id: createTaskResponse.Id,
	}
	getTaskResponse, err := client.GetTask(context.Background(), getTaskRequest)
	if err != nil {
		log.Fatalf("GetTask failed: %v", err)
	}
	fmt.Printf("Retrieved Task: %v\n", getTaskResponse)

	conn.Close()
}
