// 代码生成时间: 2025-08-26 01:01:27
package main

import (
    "context"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os"
    "path/filepath"
    "time"
    "google.golang.org/grpc"
)

// FileSyncService 文件同步服务接口
type FileSyncService interface {
    FileSync(ctx context.Context, in *FileSyncRequest) (*FileSyncResponse, error)
}

// FileSyncRequest 文件同步请求
type FileSyncRequest struct {
    SourcePath string
    DestinationPath string
}

// FileSyncResponse 文件同步响应
type FileSyncResponse struct {
    Result string
}

// fileSyncServer 文件同步服务服务器
type fileSyncServer struct{}

// FileSync 实现FileSyncService接口
func (s *fileSyncServer) FileSync(ctx context.Context, in *FileSyncRequest) (*FileSyncResponse, error) {
    // 检查源路径是否存在
    _, err := os.Stat(in.SourcePath)
    if err != nil {
        return &FileSyncResponse{Result: "Source path not found"}, nil
    }

    // 创建目标路径目录
    if err := os.MkdirAll(filepath.Dir(in.DestinationPath), 0755); err != nil {
        return &FileSyncResponse{Result: "Failed to create destination directory"}, nil
    }

    // 复制文件
    src, err := os.Open(in.SourcePath)
    if err != nil {
        return &FileSyncResponse{Result: "Failed to open source file"}, nil
    }
    defer src.Close()

    dst, err := os.Create(in.DestinationPath)
    if err != nil {
        return &FileSyncResponse{Result: "Failed to create destination file"}, nil
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return &FileSyncResponse{Result: "Failed to copy file"}, nil
    }
    if err := dst.Sync(); err != nil {
        return &FileSyncResponse{Result: "Failed to sync file"}, nil
    }

    return &FileSyncResponse{Result: "File synced successfully"}, nil
}

func main() {
    // 监听本地端口
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建gRPC服务器
    s := grpc.NewServer()
    fileSyncService := &fileSyncServer{}

    // 注册文件同步服务
    RegisterFileSyncServiceServer(s, fileSyncService)

    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
