// 代码生成时间: 2025-09-14 13:40:22
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "io"
    "os"
    "path/filepath"
    "archive/zip"
)

// 文件压缩解压服务
type FileCompressionService struct{}

// DecompressZip 接口定义
type DecompressZipServer interface {
    Decompress(context.Context, *DecompressRequest) (*DecompressResponse, error)
}

// DecompressRequest 定义请求参数
type DecompressRequest struct {
    FileName string
}

// DecompressResponse 定义响应参数
type DecompressResponse struct {
    Success bool
    Message string
}

// 实现 DecompressZipServer 接口
func (f *FileCompressionService) Decompress(ctx context.Context, req *DecompressRequest) (*DecompressResponse, error) {
    // 检查文件是否存在
    if _, err := os.Stat(req.FileName); os.IsNotExist(err) {
        return &DecompressResponse{Success: false, Message: "File not found"}, nil
    }

    // 创建解压目录
    destDir, err := os.MkdirTemp("./", "zip-extract-*")
    if err != nil {
        return nil, fmt.Errorf("failed to create temp directory: %w", err)
    }
    defer os.RemoveAll(destDir) // 清理临时目录

    // 打开压缩文件
    file, err := zip.OpenReader(req.FileName)
    if err != nil {
        return nil, fmt.Errorf("failed to open zip file: %w", err)
    }
    defer file.Close()

    // 解压文件
    for _, file := range file.File {
        err := func(file *zip.File) error {
            filePath := filepath.Join(destDir, file.Name)
            if file.FileInfo().IsDir() {
                // 创建目录
                os.MkdirAll(filePath, os.ModePerm)
                return nil
            }
            
            // 创建文件
            outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
            if err != nil {
                return err
            }
            defer outFile.Close()
            
            // 读取压缩文件中的内容并写入新文件
            rc, err := file.Open()
            if err != nil {
                return err
            }
            defer rc.Close()
            _, err = io.Copy(outFile, rc)
            return err
        }(file)
        if err != nil {
            return nil, fmt.Errorf("failed to decompress file: %w", err)
        }
    }

    return &DecompressResponse{Success: true, Message: "Decompressed successfully"}, nil
}

func main() {
    // 服务地址
    address := ":50051"
    service := FileCompressionService{}

    // 监听服务
    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建 gRPC 服务
    s := grpc.NewServer()
    RegisterDecompressZipServer(s, service)

    // 启动服务
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
