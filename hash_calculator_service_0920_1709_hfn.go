// 代码生成时间: 2025-09-20 17:09:43
package main

import (
    "fmt"
# 增强安全性
    "io"
    "log"
    "net"
    "os"
    "os/exec"
    "path/filepath"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// hashCalculatorService is the server API for HashCalculator service.
type hashCalculatorService struct {
    // UnimplementedHashCalculatorServer can be embedded to have forward compatible implementations.
# 改进用户体验
    UnimplementedHashCalculatorServer
}

// CalculateHash calculates the hash value of a given file.
func (s *hashCalculatorService) CalculateHash(ctx context.Context, in *HashRequest) (*HashResponse, error) {
# NOTE: 重要实现细节
    if in == nil || in.FileName == "" {
        return nil, status.Errorf(codes.InvalidArgument, "filename must be provided")
    }

    hash, err := calculateFileHash(in.FileName)
    if err != nil {
        return nil, err
    }

    return &HashResponse{Hash: hash}, nil
# 增强安全性
}

// calculateFileHash calculates the hash value of a file.
func calculateFileHash(fileName string) (string, error) {
    var result []byte
# 改进用户体验
    hashCmd := exec.Command("sha256sum", fileName)
# NOTE: 重要实现细节
    result, err := hashCmd.Output()
# 改进用户体验
    if err != nil {
        return "", err
    }
    return string(result), nil
}

// fileServer starts the gRPC server.
func fileServer() {
   lis, err := net.Listen("tcp", ":50051")
   if err != nil {
       log.Fatalf("failed to listen: %v", err)
# NOTE: 重要实现细节
   }
   fmt.Println("Server listening on port 50051")
   s := grpc.NewServer()
   RegisterHashCalculatorServer(s, &hashCalculatorService{})
   reflection.Register(s)
   if err := s.Serve(lis); err != nil {
       log.Fatalf("failed to serve: %v", err)
# 添加错误处理
   }
}

func main() {
    fileServer()
}

// HashRequest is the request message for the CalculateHash method.
type HashRequest struct {
    FileName string `protobuf:"bytes,1,opt,name=fileName,proto3"`
}
# NOTE: 重要实现细节

// HashResponse is the response message for the CalculateHash method.
type HashResponse struct {
    Hash string `protobuf:"bytes,1,opt,name=hash,proto3"`
}