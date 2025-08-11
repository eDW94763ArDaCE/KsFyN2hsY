// 代码生成时间: 2025-08-11 12:39:33
package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
)

// SortingAlgorithmService defines the gRPC service
type SortingAlgorithmService struct{}

// SortNumbers sorts the given slice of integers using a simple sorting algorithm
func (s *SortingAlgorithmService) SortNumbers(ctx context.Context, req *SortNumbersRequest) (*SortNumbersResponse, error) {
    // Validate the request
    if req.Numbers == nil || len(req.Numbers.Numbers) == 0 {
        return nil, fmt.Errorf("empty or nil numbers slice provided")
    }

    // Sort the slice of numbers
    numbers := make([]int32, len(req.Numbers.Numbers))
    copy(numbers, req.Numbers.Numbers)
    // A simple bubble sort algorithm for demonstration purposes
    for i := 0; i < len(numbers); i++ {
        for j := 0; j < len(numbers)-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
            }
        }
    }

    // Return the sorted numbers in the response
    return &SortNumbersResponse{Numbers: &Numbers{Numbers: numbers}}, nil
}

// Main function to start the gRPC server
func main() {
    // Generate random numbers and create a request
    rand.Seed(time.Now().UnixNano())
    numbers := make([]int32, 10)
    for i := range numbers {
        numbers[i] = int32(rand.Intn(100))
    }
    req := &SortNumbersRequest{Numbers: &Numbers{Numbers: numbers}}

    // Create a connection to the server
    conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := NewSortingAlgorithmServiceClient(conn)

    // Call the SortNumbers method
    resp, err := c.SortNumbers(context.Background(), req)
    if err != nil {
        log.Fatalf("could not sort numbers: %v", err)
    }
    fmt.Printf("Sorted numbers: %v", resp.Numbers.Numbers)
}

// SortNumbersRequest defines the request message for the SortNumbers RPC
type SortNumbersRequest struct {
    Numbers *Numbers `protobuf:"bytes,1,opt,name=numbers"`
}

// SortNumbersResponse defines the response message for the SortNumbers RPC
type SortNumbersResponse struct {
    Numbers *Numbers `protobuf:"bytes,1,opt,name=numbers"`
}

// Numbers is a wrapper message for a repeated field of int32s
type Numbers struct {
    Numbers []int32 `protobuf:"varint,1,rep,packed,name=numbers"`
}
