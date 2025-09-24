// 代码生成时间: 2025-09-24 17:18:44
 * interactive_chart_generator.go
 * This program uses Go and the gRPC framework to create an interactive chart generator service.
 */

package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"

    "path/to/your/interactive_chart_generator/pb" // Replace with the actual path to your generated protobuf package
)

// server represents the interactive chart generator server.
type server struct {
    pb.UnimplementedInteractiveChartGeneratorServer

    // Add fields here to handle the state of the server
}

// NewServer creates a new interactive chart generator server.
func NewServer() *server {
    return &server{}
}

// GenerateChart implements the GenerateChart RPC method.
func (s *server) GenerateChart(ctx context.Context, req *pb.ChartRequest) (*pb.ChartResponse, error) {
    // Check for nil request
    if req == nil {
        return nil, fmt.Errorf("request cannot be nil")
    }

    // Process the request and generate the chart
    // This is where you would implement the logic to generate the chart based on the request

    // For demonstration purposes, we're just returning a success response with a dummy chart ID
    chartID := fmt.Sprintf("chart_%d", 1) // Replace with actual chart ID generation logic

    // Return a success response
    return &pb.ChartResponse{ChartId: chartID}, nil
}

// startServer starts the gRPC server and begins listening for connections.
func startServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    fmt.Println("Starting interactive chart generator service... on port 50051")

    s := grpc.NewServer()
    pb.RegisterInteractiveChartGeneratorServer(s, NewServer())
    reflection.Register(s) // Enable server reflection for debugging
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    startServer()
}
