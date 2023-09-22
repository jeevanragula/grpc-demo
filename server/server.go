package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "server/task/v1"
	"sync"
)

type TaskServiceServer struct {
	mu     sync.Mutex
	Tasks  map[int64]*pb.Task
	nextID int64
	pb.UnimplementedTaskServiceServer
}

func (s *TaskServiceServer) DeleteTask(ctx context.Context, request *pb.DeleteTaskRequest) (*emptypb.Empty, error) {
	delete(s.Tasks, request.GetId())
	return &emptypb.Empty{}, nil
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &pb.Task{
		Id:          s.nextID,
		Title:       req.Title,
		Description: req.Description,
	}
	s.Tasks[s.nextID] = task
	s.nextID++
	return task, nil
}

func (s *TaskServiceServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.Tasks[req.Id]
	if !ok {
		return nil, fmt.Errorf("task not found")
	}

	return task, nil
}

func (s *TaskServiceServer) ListTasks(context.Context, *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks := make([]*pb.Task, 0)
	for id, task := range s.Tasks {
		task := &pb.Task{
			Id:          id,
			Title:       task.Title,
			Description: task.Description,
		}
		tasks = append(tasks, task)
	}

	return &pb.ListTasksResponse{
		Tasks: tasks,
	}, nil
}
