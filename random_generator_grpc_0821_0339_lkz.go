// 代码生成时间: 2025-08-21 03:39:44
// random_generator_grpc.go
// 该文件定义了一个使用GRPC框架的随机数生成器服务

package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

// RandomNumberService 是一个GRPC服务，用于生成随机数
type RandomNumberService struct {}

// GenerateRandomNumber 响应客户端请求，生成一个随机数
func (s *RandomNumberService) GenerateRandomNumber(ctx context.Context, req *GenerateRandomNumberRequest) (*GenerateRandomNumberResponse, error) {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 根据请求中的参数生成随机数
	num := rand.Intn(req.GetMax()) + req.GetMin()

	// 创建响应并返回
	return &GenerateRandomNumberResponse{
		Number: int32(num),
	}, nil
}
on
// main 函数定义了GRPC服务器的启动逻辑
func main() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("Failed to listen: ", err)
		os.Exit(1)
	}

	// 创建gRPC服务器
	server := grpc.NewServer()

	// 注册服务
	RegisterRandomNumberServiceServer(server, &RandomNumberService{})

	// 启动服务器
	if err := server.Serve(lis); err != nil {
		fmt.Println("Server exited with error: ", err)
	}
}

// GenerateRandomNumberRequest 是生成随机数的请求消息
type GenerateRandomNumberRequest struct {
	Min int32 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max int32 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
}

// GenerateRandomNumberResponse 是生成随机数的响应消息
type GenerateRandomNumberResponse struct {
	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

// RandomNumberServiceServer 是随机数服务的GRPC服务器接口
type RandomNumberServiceServer interface {
	GenerateRandomNumber(context.Context, *GenerateRandomNumberRequest) (*GenerateRandomNumberResponse, error)
}

// RegisterRandomNumberServiceServer 将服务注册到GRPC服务器
func RegisterRandomNumberServiceServer(s *grpc.Server, srv *RandomNumberService) {
	RegisterRandomNumberServiceServer(s, srv)
}
