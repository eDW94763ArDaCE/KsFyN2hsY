// 代码生成时间: 2025-08-27 12:33:53
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "log"

    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// 文件重命名请求
type RenameRequest struct {
    SourceName string "json:\"sourceName\""
    DestName   string "json:\"destName\""
}

// 文件重命名响应
type RenameResponse struct {
    Success bool   "json:\"success\""
    Message string "json:\"message\""
}

// 文件服务定义
type FileServiceServer struct{}

// RenameFile 重命名文件
func (s *FileServiceServer) RenameFile(ctx context.Context, req *RenameRequest) (*RenameResponse, error) {
    // 检查源文件是否存在
    srcPath := filepath.Join(req.SourceName)
    if _, err := os.Stat(srcPath); os.IsNotExist(err) {
        return nil, status.Errorf(codes.NotFound, "源文件不存在: %s", req.SourceName)
    }
    // 检查目标文件是否已存在
    destPath := filepath.Join(req.DestName)
    if _, err := os.Stat(destPath); err == nil {
        return nil, status.Errorf(codes.AlreadyExists, "目标文件已存在: %s", req.DestName)
    }
    // 重命名文件
    if err := os.Rename(srcPath, destPath); err != nil {
        return nil, status.Errorf(codes.Internal, "文件重命名失败: %s", err)
    }
    return &RenameResponse{Success: true, Message: "文件重命名成功"}, nil
}

// BatchRenameFile 批量重命名文件
func (s *FileServiceServer) BatchRenameFile(ctx context.Context, req *BatchRenameRequest) (*emptypb.Empty, error) {
    for _, item := range req.Items {
        if _, err := s.RenameFile(ctx, &RenameRequest{SourceName: item.SourceName, DestName: item.DestName}); err != nil {
            return nil, err
        }
    }
    return &emptypb.Empty{}, nil
}

// 服务定义
type batchRenameServiceServer struct {
    FileServiceServer
}

func main() {
    // 监听端口
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    // 创建gRPC服务器
    s := grpc.NewServer()
    // 注册服务
    batchrenamepb.RegisterBatchRenameServiceServer(s, &batchRenameServiceServer{})
    // 启动服务
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}