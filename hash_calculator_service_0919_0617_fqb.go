// 代码生成时间: 2025-09-19 06:17:45
package main

import (
    "context"
    "fmt"
    "io"
    "log"

    "crypto/sha256"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// Define the service definition in the protobuf file and generate the Go code using protoc.

// HashCalculatorService provides RPC methods for calculating hash values.
type HashCalculatorService struct {}

// CalculateHash receives input data and returns its SHA-256 hash.
func (s *HashCalculatorService) CalculateHash(ctx context.Context, in *HashRequest) (*HashResponse, error) {
    if in == nil || in.Data == nil || len(in.Data) == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Input data is missing or empty")
    }

    hash := sha256.Sum256(in.Data)
    return &HashResponse{Hash: hash[:]}, nil
}

// HashRequest defines the request structure for calculating a hash.
type HashRequest struct {
    Data []byte `protobuf:"varint,1,opt,name=data,proto3"`
}

// HashResponse defines the response structure for the hash calculation.
type HashResponse struct {
    Hash []byte `protobuf:"varint,1,opt,name=hash,proto3"`
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")

    s := grpc.NewServer()
    pb.RegisterHashCalculatorServiceServer(s, &HashCalculatorService{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
