// 代码生成时间: 2025-09-01 15:19:28
package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
)

// 定义清洗数据的请求类型
type CleanDataRequest struct {
	Data string `protobuf:"bytes,1,opt,name=data,proto3"`
}

// 定义清洗数据的响应类型
type CleanDataResponse struct {
	CleanedData string `protobuf:"bytes,1,opt,name=cleaned_data,proto3"`
}

// 定义服务接口
type DataCleaningServiceServer struct {
}

// 实现数据清洗的方法
func (s *DataCleaningServiceServer) CleanData(ctx context.Context, req *CleanDataRequest) (*CleanDataResponse, error) {
	if req.Data == "" {
		return nil, fmt.Errorf("data is empty")
	}

	// 示例：去除字符串中的空格
	cleanedData := strings.TrimSpace(req.Data)

	return &CleanDataResponse{CleanedData: cleanedData}, nil
}

// 实现gRPC服务服务器
func serve() {
	ls, err := grpc.NewServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	// 注册服务
	pb.RegisterDataCleaningServiceServer(ls, &DataCleaningServiceServer{})

	// 监听并服务
	if err := ls.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// 设置监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 启动服务
	go serve()

	// 等待服务器停止信号
	select {}
}
