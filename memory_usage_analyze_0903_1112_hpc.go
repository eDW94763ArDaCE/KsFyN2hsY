// 代码生成时间: 2025-09-03 11:12:24
package main

import (
    "fmt"
    "log"
    "net"
    "runtime"
    "strings"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// MemoryUsageService defines the service
type MemoryUsageService struct {}

// MemoryUsageResponse defines the response message for MemoryUsage
type MemoryUsageResponse struct {
    // MemoryUsage is the memory usage in bytes
    MemoryUsage uint64 `protobuf:"varint,1,opt,name=memory_usage,json=memoryUsage,proto3"`
}

// MemoryUsage defines the RPC to get memory usage
func (s *MemoryUsageService) MemoryUsage(ctx context.Context, _ *emptypb.Empty) (*MemoryUsageResponse, error) {
    // Get the memory usage
    memUsage := new(runtime.MemStats)
    runtime.ReadMemStats(memUsage)
    return &MemoryUsageResponse{
        MemoryUsage: memUsage.Alloc, // or memUsage.Sys for total memory allocated by the process
    }, nil
}

// server is used to implement memory_usage.MemoryUsageServer
type server struct{
    UnimplementedMemoryUsageServer
}

// MemoryUsage is a server method for MemoryUsage
func (s *server) MemoryUsage(ctx context.Context, in *emptypb.Empty) (*MemoryUsageResponse, error) {
    return &MemoryUsageResponse{
        MemoryUsage: runtimeMemStats.Alloc, // or runtimeMemStats.Sys for total memory allocated by the process
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    grpcServer := grpc.NewServer()
    memory_usage.RegisterMemoryUsageServer(grpcServer, &server{})
    reflection.Register(grpcServer)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// NOTE: This code assumes the existence of a protobuf generated code for the MemoryUsage service
// which includes the definitions of MemoryUsageResponse and MemoryUsageServer.
// The above code provides a basic structure for the implementation.
// You will need to create the `.proto` file, generate the Go code using `protoc`,
// and include the generated code in this project.
