// 代码生成时间: 2025-09-12 02:23:12
// math_service.go 定义了一个数学服务，该服务使用gRPC框架提供远程数学计算功能。
    
package main

import (
    "fmt"
    "log"
    "math"
    
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    
    pb "path/to/your/protobufs" // 导入protobuf生成的pb.go文件
)

// MathService 定义了数学服务的方法。
type MathService struct {
    // 这里可以添加一些服务需要的字段
}

// Add 实现了添加两个数的功能。
func (s *MathService) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
    if in.GetA() == nil || in.GetB() == nil {
        return nil, fmt.Errorf("invalid input: both numbers are required")
    }
    
    sum := in.GetA().GetValue() + in.GetB().GetValue()
    return &pb.AddResponse{Result: sum}, nil
}

// Div 实现了两个数相除的功能。
func (s *MathService) Div(ctx context.Context, in *pb.DivRequest) (*pb.DivResponse, error) {
    if in.GetA() == nil || in.GetB() == nil {
        return nil, fmt.Errorf("invalid input: both numbers are required")
    }
    
    if in.GetB().GetValue() == 0 {
        return nil, fmt.Errorf("division by zero is not allowed")
    }
    
    div := in.GetA().GetValue() / in.GetB().GetValue()
    return &pb.DivResponse{Result: div}, nil
}

// server 是一个gRPC服务器实例。
type server struct {
    pb.UnimplementedMathServiceServer
}

// RegisterServer 注册数学服务到gRPC服务器。
func RegisterServer(s *grpc.Server) {
    pb.RegisterMathServiceServer(s, &server{})
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on 50051")
    
    s := grpc.NewServer()
    RegisterServer(s)
    reflection.Register(s) // 注册反射服务
    
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
