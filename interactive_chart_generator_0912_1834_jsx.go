// 代码生成时间: 2025-09-12 18:34:23
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// ChartService defines the structure of the service
type ChartService struct {
    // embedding the UnimplementedChartServiceServer to implement the server interface
    // This is where we will implement the methods of the service
    grpc.UnimplementedChartServiceServer
}

// Define the ChartRequest message
type ChartRequest struct {
    // This message should include all the necessary fields to generate a chart
    // For example, data points, chart type, etc.
    DataPoints []float64 `protobuf:"varint,1,rep,packed==false,name=data_points,json=dataPoints" json:"dataPoints"`
    ChartType string `protobuf:"bytes,2,opt,name=chart_type,json=chartType" json:"chartType"`
}

// Define the ChartResponse message
type ChartResponse struct {
    // This message will contain the generated chart image data or a success message
    ImageData []byte `protobuf:"bytes,1,opt,name=image_data,json=imageData" json:"imageData"`
}

// GenerateChart generates an interactive chart based on the provided request
func (s *ChartService) GenerateChart(ctx context.Context, req *ChartRequest) (*ChartResponse, error) {
    // Check if the request is valid
    if req == nil || len(req.DataPoints) == 0 {
        return nil, status.Error(codes.InvalidArgument, "Invalid chart request")
    }

    // Generate the chart based on the provided data points and chart type
    // This is a placeholder for the actual chart generation logic
    // In a real-world scenario, you would use a charting library to generate the chart
    // and then encode the image data to send back in the response
    imageData := []byte{} // Placeholder for the actual image data

    // Return a ChartResponse with the generated image data
    return &ChartResponse{ImageData: imageData}, nil
}

func main() {
    // Define the server address and port
    serverAddress := ":50051"

    // Create a new gRPC server
    srv := grpc.NewServer()

    // Register the ChartService with the server
    grpc.RegisterChartServiceServer(srv, &ChartService{})

    // Start the server
    lis, err := net.Listen("tcp", serverAddress)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    go func() {
        if err := srv.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

    // Wait for the interrupt signal to gracefully shutdown the server
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Printf("Shutting down gRPC server...
")
    srv.GracefulStop()
    log.Printf("gRPC server exited
")
}
