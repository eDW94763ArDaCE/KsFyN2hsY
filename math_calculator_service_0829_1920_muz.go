// 代码生成时间: 2025-08-29 19:20:01
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// Define the math service
type MathServiceServer struct {}

// Define the Protobuf generated service
type mathServiceServer struct{
    MathServiceServer
}

// Define the service methods
func (s *mathServiceServer) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
    if in == nil {
        return nil, status.Errorf(codes.InvalidArgument, "received nil AddRequest")
    }
    result := in.A + in.B
    return &AddResponse{Result: result}, nil
}

func (s *mathServiceServer) Subtract(ctx context.Context, in *SubtractRequest) (*SubtractResponse, error) {
    if in == nil {
        return nil, status.Errorf(codes.InvalidArgument, "received nil SubtractRequest")
    }
    result := in.Minuend - in.Subtrahend
    return &SubtractResponse{Result: result}, nil
}

func (s *mathServiceServer) Multiply(ctx context.Context, in *MultiplyRequest) (*MultiplyResponse, error) {
    if in == nil {
        return nil, status.Errorf(codes.InvalidArgument, "received nil MultiplyRequest")
    }
    result := in.A * in.B
    return &MultiplyResponse{Result: result}, nil
}

func (s *mathServiceServer) Divide(ctx context.Context, in *DivideRequest) (*DivideResponse, error) {
    if in == nil {
        return nil, status.Errorf(codes.InvalidArgument, "received nil DivideRequest")
    }
    if in.Divisor == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "cannot divide by zero")
    }
    result := in.Dividend / in.Divisor
    return &DivideResponse{Result: result}, nil
}

// Define the main function to run the gRPC server
func main() {
   lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    s := grpc.NewServer()
    RegisterMathServiceServer(s, &mathServiceServer{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
