package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type taskServiceServer struct {
	mu     sync.Mutex
	tasks  map[int64]*pb.Task
	nextID int64
}

func (s *taskServiceServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &pb.Task{
		Id:          s.nextID,
		Title:       req.Title,
		Description: req.Description,
	}
	s.tasks[s.nextID] = task
	s.nextID++

	return task, nil
}

func (s *taskServiceServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[req.Id]
	if !ok {
		return nil, fmt.Errorf("Task not found")
	}

	return task, nil
}

// Implement UpdateTask, DeleteTask, and ListTasks methods similarly

func main() {
	port := 5000

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Server listening on port %d", port)

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &taskServiceServer{
		tasks: make(map[int64]*pb.Task),
	})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
