// 代码生成时间: 2025-08-07 06:58:00
// scheduler_service.go
// This file provides the implementation of a gRPC service for a scheduling task.

package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/timestamppb"
)

// Task represents a scheduled task.
type Task struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    NextRun  time.Time `json:"nextRun"`
    Interval time.Duration `json:"interval"`
}

// SchedulerService is the server API for SchedulerService service.
type SchedulerService struct {
    // Contains filtered or unexported fields.
}

// RunTask runs the provided task at the scheduled time.
func (s *SchedulerService) RunTask(ctx context.Context, req *Task) (*timestamppb.Timestamp, error) {
    log.Printf("Received a task with name: %s", req.Name)
    
    // Schedule the next run of the task.
    go func() {
        time.Sleep(req.NextRun.Sub(time.Now()))
        fmt.Printf("Running task: %s at: %v", req.Name, time.Now())
    }()

    return timestamppb.Now(), nil
}

// RegisterGRPCServer registers the scheduler service on the provided gRPC server.
func RegisterGRPCServer(server *grpc.Server) {
    // Register the service on the server.
    // Assuming the service definition is in the same package.
    // The service name is the full protobuf service name.
    // grpc.NewServer registers the service on the provided server.
}

// startGRPCServer starts the gRPC server on the specified port.
func startGRPCServer(address string) {
    listener, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    defer listener.Close()
    
    server := grpc.NewServer()
    RegisterGRPCServer(server)
    reflection.Register(server) // Register reflection service on gRPC server.
    
    log.Printf("Server listening at %s", address)
    if err := server.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

func main() {
    // Define the address to serve on.
    address := ":50051"
    
    // Start the gRPC server.
    startGRPCServer(address)
}