package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "server/task/v1"
)

// Implement UpdateTask, DeleteTask, and ListTasks methods similarly

func main() {
	port := 5000

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Server listening on port %d", port)

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &TaskServiceServer{
		Tasks: make(map[int64]*pb.Task),
	})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
