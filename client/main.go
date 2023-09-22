package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "task-proto/task/v1"
)

func main() {
	serverAddress := "localhost:50051" // Replace with the address of your gRPC server

	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
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

	// Implement other CRUD operations similarly

	// Close the connection gracefully
	conn.Close()
}
