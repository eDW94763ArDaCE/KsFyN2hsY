// 代码生成时间: 2025-09-05 10:13:32
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

// ApiResponseFormatterService 定义了API响应格式化服务的接口
type ApiResponseFormatterService interface {
	FormatResponse(ctx context.Context, req *FormatRequest) (*FormatResponse, error)
}

// FormatRequest 定义了格式化请求的结构
type FormatRequest struct {
	RawResponse string
}

// FormatResponse 定义了格式化响应的结构
type FormatResponse struct {
	FormattedMessage string
}

// server 实现了ApiResponseFormatterService接口
type server struct {
}

// NewServer 创建一个新的server实例
func NewServer() *server {
	return &server{}
}

// FormatResponse 实现了ApiResponseFormatterService接口的FormatResponse方法
func (s *server) FormatResponse(ctx context.Context, req *FormatRequest) (*FormatResponse, error) {
	// 检查请求是否为空
	if req == nil || req.RawResponse == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	// 格式化响应
	formattedMessage := fmt.Sprintf("Formatted Response: %s", req.RawResponse)

	// 返回格式化后的响应
	return &FormatResponse{FormattedMessage: formattedMessage}, nil
}

// startServer 启动GRPC服务器
func startServer(port string, service ApiResponseFormatterService) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建GRPC服务器
	grpcServer := grpc.NewServer()

	// 注册服务
	RegisterApiResponseFormatterServiceServer(grpcServer, service)

	// 启动服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// 创建服务实例
	service := NewServer()

	// 启动GRPC服务器
	startServer(":50051", service)
}
