// 代码生成时间: 2025-08-15 22:48:23
package sort

import (
    "fmt"
    "sort"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
    "log"
)

// Define the SortServiceServer which implements the SortService interface.
type SortServiceServer struct{}

// Implement the Sort method from the SortService service.
func (s *SortServiceServer) Sort(ctx context.Context, in *SortRequest) (*SortResponse, error) {
    if in == nil || len(in.Numbers) == 0 {
        return nil, fmt.Errorf("empty input")
    }

    // Sort the input slice using Go's built-in sort function.
    sort.Ints(in.Numbers)

    return &SortResponse{Numbers: in.Numbers}, nil
}

// SortRequest defines the request message for the Sort method.
type SortRequest struct {
    Numbers []int32
}

// SortResponse defines the response message for the Sort method.
type SortResponse struct {
    Numbers []int32
}

// server is used to implement sort.Service.
func RegisterServer(server *grpc.Server) {
    RegisterSortServiceServer(server, &SortServiceServer{})
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    RegisterServer(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
