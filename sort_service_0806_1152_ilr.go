// 代码生成时间: 2025-08-06 11:52:26
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
    "math/rand"
    "sort"
    "time"
)

// SortService defines the service for sorting
type SortService struct{}

// SortNumbers sorts a slice of integers using the sort package
func (s *SortService) SortNumbers(ctx context.Context, req *SortRequest) (*SortResponse, error) {
    // Check if the request contains any integers
    if len(req.Numbers) == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "empty slice of numbers")
    }

    // Copy the slice to avoid modifying the original slice
    numbers := make([]int, len(req.Numbers))
    copy(numbers, req.Numbers)

    // Sort the slice
    sort.Ints(numbers)

    // Return the sorted slice
    return &SortResponse{Numbers: numbers}, nil
}

// StartServer starts the gRPC server
func StartServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")
    grpcServer := grpc.NewServer()
    RegisterSortServiceServer(grpcServer, &SortService{})
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// GenerateRandomNumbers generates a slice of random integers
func GenerateRandomNumbers(count, max int) []int {
    rand.Seed(time.Now().UnixNano())
    numbers := make([]int, count)
    for i := range numbers {
        numbers[i] = rand.Intn(max)
    }
    return numbers
}

// SortRequest defines the request message for the SortNumbers method
type SortRequest struct {
    Numbers []int
}

// SortResponse defines the response message for the SortNumbers method
type SortResponse struct {
    Numbers []int
}

// RegisterSortServiceServer registers the sort service on the gRPC server
func RegisterSortServiceServer(s *grpc.Server, srv *SortService) {
    RegisterSortServiceServer(s, srv)
}

func main() {
    StartServer()
}
