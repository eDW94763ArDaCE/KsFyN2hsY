// 代码生成时间: 2025-08-18 20:00:31
package main

import (
    "fmt"
    "log"
    "net"
    "runtime"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    p "your_project_path/your_package_name" // 替换为你的proto文件的package路径
)

// MemoryUsageService 定义了一个服务，用于内存使用情况分析
type MemoryUsageService struct {
    // 这里可以添加需要的字段
}

// MemoryUsage 服务服务器
type MemoryUsageServer struct {
    p.UnimplementedMemoryUsageServer
}

// GetMemoryUsage 实现了MemoryUsageService服务的GetMemoryUsage方法
func (s *MemoryUsageServer) GetMemoryUsage(ctx context.Context, in *p.GetMemoryUsageRequest) (*p.GetMemoryUsageResponse, error) {
    // 获取当前内存使用情况
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    // 计算内存使用情况
    usedMemory := m.Alloc
    totalMemory := m.Sys
    freeMemory := totalMemory - usedMemory

    // 构造响应
    response := &p.GetMemoryUsageResponse{
        UsedMemory: usedMemory,
        FreeMemory: freeMemory,
        TotalMemory: totalMemory,
    }

    return response, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建gRPC服务器
    s := grpc.NewServer()

    // 将MemoryUsageServer注册到gRPC服务器
    p.RegisterMemoryUsageServer(s, &MemoryUsageServer{})

    // 注册gRPC反射服务
    reflection.Register(s)

    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// NOTE: 这里假设你的proto文件中定义了MemoryUsage服务和相应的GetMemoryUsage方法
// 你的proto文件应该如下所示：
/*
service MemoryUsage {
    rpc GetMemoryUsage(GetMemoryUsageRequest) returns (GetMemoryUsageResponse) {}
}

message GetMemoryUsageRequest {}

message GetMemoryUsageResponse {
    int64 UsedMemory = 1;
    int64 FreeMemory = 2;
    int64 TotalMemory = 3;
}
*/
