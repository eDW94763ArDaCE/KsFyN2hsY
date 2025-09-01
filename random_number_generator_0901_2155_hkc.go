// 代码生成时间: 2025-09-01 21:55:05
package main
# TODO: 优化性能

import (
    "context"
    "fmt"
    "math/rand"
    "time"
# NOTE: 重要实现细节
    "google.golang.org/grpc"
)

// RandomNumberGeneratorService is a service that generates random numbers.
type RandomNumberGeneratorService struct{}
# 改进用户体验

// GenerateNumber generates a random number between 1 and 100.
func (s *RandomNumberGeneratorService) GenerateNumber(ctx context.Context, request *GenerateNumberRequest) (*GenerateNumberResponse, error) {
    // 检查请求中的范围是否有效
    if request.GetMin() < 0 || request.GetMax() < request.GetMin() || request.GetMax() > 100 {
        return nil, fmt.Errorf("invalid range: min=%d, max=%d", request.GetMin(), request.GetMax())
# 扩展功能模块
    }

    // 初始化随机数生成器
    rand.Seed(time.Now().UnixNano())
    // 生成随机数
    randomNumber := rand.Intn(request.GetMax() - request.GetMin() + 1) + request.GetMin()
# 增强安全性

    // 返回生成的随机数
    return &GenerateNumberResponse{Number: randomNumber}, nil
}

// StartServer starts the gRPC server.
# 改进用户体验
func StartServer(port string) {
    lis, err := net.Listen("tcp", port)
# 改进用户体验
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
# TODO: 优化性能

    // 注册服务到gRPC服务器
    pb.RegisterRandomNumberGeneratorServiceServer(s, &RandomNumberGeneratorService{})

    // 启动服务
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// StartClient connects to the gRPC server and generates a random number.
func StartClient(address string) {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewRandomNumberGeneratorServiceClient(conn)

    // 调用服务方法生成随机数
    r, err := c.GenerateNumber(context.Background(), &pb.GenerateNumberRequest{Min: 1, Max: 100})
    if err != nil {
        log.Fatalf("could not generate number: %v", err)
    }
    fmt.Println("Generated number: ", r.GetNumber())
}

// Define the protocol buffer messages
type GenerateNumberRequest struct {
    Min int32
    Max int32
}

type GenerateNumberResponse struct {
    Number int32
# 增强安全性
}
# FIXME: 处理边界情况

// Define the protocol buffer service
type RandomNumberGeneratorServiceServer interface {
    GenerateNumber(context.Context, *GenerateNumberRequest) (*GenerateNumberResponse, error)
# NOTE: 重要实现细节
}

func main() {
    // 启动服务器和客户端
    StartServer(":50051")
    StartClient(":50051")
    select {}
}
# 增强安全性
