// 代码生成时间: 2025-09-07 03:22:54
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/reflection"
    pb "path/to/your/protobuf/package" // Replace with your actual protobuf package path
)

type server struct {
    pb.UnimplementedYourServiceServer // Replace with your service name
}

func (s *server) YourServiceMethod(ctx context.Context, req *pb.YourServiceRequest) (*pb.YourServiceResponse, error) {
    // Implement your service method here
    return &pb.YourServiceResponse{}, nil
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterYourServiceServer(grpcServer, &server{}) // Register your service
    reflection.Register(grpcServer) // Register reflection service

    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

    // Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch
    log.Println("Shutting down gRPC server...