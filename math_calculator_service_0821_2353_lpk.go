// 代码生成时间: 2025-08-21 23:53:49
package main

import (
    "context"
    "fmt"
    "log"
    "math"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/proto"
)

// MathCalculatorService defines the math calculator service
type MathCalculatorService struct{}

// MathCalculatorServiceServer defines the gRPC server
type MathCalculatorServiceServer struct {
    MathCalculatorService
}

// Add performs addition
func (s *MathCalculatorServiceServer) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
    if in == nil {
        return nil, fmt.Errorf("input cannot be nil")
    }
    result := in.A + in.B
    return &AddResponse{Result: result}, nil
}

// Subtract performs subtraction
func (s *MathCalculatorServiceServer) Subtract(ctx context.Context, in *SubtractRequest) (*SubtractResponse, error) {
    if in == nil {
        return nil, fmt.Errorf("input cannot be nil")
    }
    result := in.A - in.B
    return &SubtractResponse{Result: result}, nil
}

// Multiply performs multiplication
func (s *MathCalculatorServiceServer) Multiply(ctx context.Context, in *MultiplyRequest) (*MultiplyResponse, error) {
    if in == nil {
        return nil, fmt.Errorf("input cannot be nil")
    }
    result := in.A * in.B
    return &MultiplyResponse{Result: result}, nil
}

// Divide performs division
func (s *MathCalculatorServiceServer) Divide(ctx context.Context, in *DivideRequest) (*DivideResponse, error) {
    if in == nil {
        return nil, fmt.Errorf("input cannot be nil")
    }
    if in.B == 0 {
        return nil, fmt.Errorf("denominator cannot be zero")
    }
    result := in.A / in.B
    return &DivideResponse{Result: result}, nil
}

// StartServer starts the gRPC server
func StartServer(address string) {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Printf("Server listening on %s
", address)

    grpcServer := grpc.NewServer()
    // Register reflection service on gRPC server.
    reflection.Register(grpcServer)

    // Register MathCalculatorServiceServer on gRPC server.
    grpcServer.RegisterService(&MathCalculatorService_ServiceDesc, &MathCalculatorServiceServer{})

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// AddRequest represents the request for addition
type AddRequest struct {
    A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
    B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

// AddResponse represents the response for addition
type AddResponse struct {
    Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

// SubtractRequest represents the request for subtraction
type SubtractRequest struct {
    A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
    B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

// SubtractResponse represents the response for subtraction
type SubtractResponse struct {
    Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

// MultiplyRequest represents the request for multiplication
type MultiplyRequest struct {
    A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
    B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

// MultiplyResponse represents the response for multiplication
type MultiplyResponse struct {
    Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

// DivideRequest represents the request for division
type DivideRequest struct {
    A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
    B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

// DivideResponse represents the response for division
type DivideResponse struct {
    Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func main() {
    // Start the gRPC server
    StartServer(":50051")
}
