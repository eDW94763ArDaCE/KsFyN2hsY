// 代码生成时间: 2025-09-15 13:00:04
// stat_analysis_service.go

// Package stat provides a gRPC service for statistical data analysis.
package stat

import (
    "context"
    "fmt"
    "math"
    "log"
    "sort"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

// Define the StatAnalysisServiceServer which will implement the StatAnalysisService
type StatAnalysisServiceServer struct {
    // embedded fields can be added here for additional functionality
}

// Define the function signatures required by the service
type StatAnalysisServiceServer interface {
    CalculateMean(ctx context.Context, in *MeanRequest) (*MeanResponse, error)
    CalculateMedian(ctx context.Context, in *MedianRequest) (*MedianResponse, error)
    CalculateMode(ctx context.Context, in *ModeRequest) (*ModeResponse, error)
}

// MeanRequest is the message type for the CalculateMean RPC
type MeanRequest struct {
    Data []float64 `protobuf:"fixed64,1,rep,name=data" json:"data"`
}

// MeanResponse is the message type for the CalculateMean RPC
type MeanResponse struct {
    Mean float64 `protobuf:"fixed64,1,opt,name=mean,proto3" json:"mean"`
}

// MedianRequest is the message type for the CalculateMedian RPC
type MedianRequest struct {
    Data []float64 `protobuf:"fixed64,1,rep,name=data" json:"data"`
}

// MedianResponse is the message type for the CalculateMedian RPC
type MedianResponse struct {
    Median float64 `protobuf:"fixed64,1,opt,name=median,proto3" json:"median"`
}

// ModeRequest is the message type for the CalculateMode RPC
type ModeRequest struct {
    Data []float64 `protobuf:"fixed64,1,rep,name=data" json:"data"`
}

// ModeResponse is the message type for the CalculateMode RPC
type ModeResponse struct {
    Mode float64 `protobuf:"fixed64,1,opt,name=mode,proto3" json:"mode"`
}

// Implement the StatAnalysisServiceServer interface
func (s *StatAnalysisServiceServer) CalculateMean(ctx context.Context, in *MeanRequest) (*MeanResponse, error) {
    // Calculate the mean of the provided data
    mean := calculateMean(in.Data)
    return &MeanResponse{Mean: mean}, nil
}

func (s *StatAnalysisServiceServer) CalculateMedian(ctx context.Context, in *MedianRequest) (*MedianResponse, error) {
    // Calculate the median of the provided data
    median := calculateMedian(in.Data)
    return &MedianResponse{Median: median}, nil
}

func (s *StatAnalysisServiceServer) CalculateMode(ctx context.Context, in *ModeRequest) (*ModeResponse, error) {
    // Calculate the mode of the provided data
    mode := calculateMode(in.Data)
    return &ModeResponse{Mode: mode}, nil
}

// Helper function to calculate the mean
func calculateMean(data []float64) float64 {
    var sum float64
    for _, value := range data {
        sum += value
    }
    return sum / float64(len(data))
}

// Helper function to calculate the median
func calculateMedian(data []float64) float64 {
    sortedData := make([]float64, len(data))
    copy(sortedData, data)
    sort.Float64s(sortedData)
    middle := len(sortedData) / 2
    if len(sortedData)%2 == 0 {
        return (sortedData[middle-1] + sortedData[middle]) / 2
    }
    return sortedData[middle]
}

// Helper function to calculate the mode
func calculateMode(data []float64) float64 {
    var counts = make(map[float64]int)
    for _, number := range data {
        counts[number]++
    }
    var maxCount int
    var mode float64
    for number, count := range counts {
        if count > maxCount {
            maxCount = count
            mode = number
        }
    }
    return mode
}

// Register and serve the gRPC service
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")
    s := grpc.NewServer()
    // Register the service with the gRPC server
    RegisterStatAnalysisServiceServer(s, &StatAnalysisServiceServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}