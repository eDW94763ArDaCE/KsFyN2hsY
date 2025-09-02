// 代码生成时间: 2025-09-02 12:49:10
package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

// CSVProcessorService 定义处理CSV文件的gRPC服务
type CSVProcessorService struct{}

// ProcessCSVFiles 处理提供的CSV文件列表
func (s *CSVProcessorService) ProcessCSVFiles(ctx context.Context, req *ProcessRequest) (*ProcessResponse, error) {
    // 检查请求中的文件路径列表是否为空
    if len(req.FilePaths) == 0 {
        return nil, fmt.Errorf("no file paths provided")
    }

    // 遍历文件路径列表进行处理
    for _, filePath := range req.FilePaths {
        if err := processSingleCSVFile(filePath); err != nil {
            return nil, fmt.Errorf("failed to process file %s: %w", filePath, err)
        }
    }

    // 返回成功响应
    return &ProcessResponse{
        Success: true,
        Message: "All files processed successfully",
    }, nil
}

// processSingleCSVFile 处理单个CSV文件
func processSingleCSVFile(filePath string) error {
    // 打开CSV文件
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filePath, err)
    }
    defer file.Close()

    // 遍历文件中的每一行
    scanner := csv.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // 处理CSV行（例如，转换或验证数据）
        fmt.Println("Processing line:", line)
    }
    
    // 检查是否有错误发生
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("failed to read file %s: %w", filePath, err)
    }

    return nil
}

// ProcessRequest 定义处理CSV文件的请求结构
type ProcessRequest struct {
    FilePaths []string `protobuf:"bytes,1,rep,name=file_paths,json=filePaths"`
}

// ProcessResponse 定义处理CSV文件的响应结构
type ProcessResponse struct {
    Success bool   `protobuf:"varint,1,opt,name=success,proto3"`
    Message string `protobuf:"bytes,2,opt,name=message,proto3"`
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    s := grpc.NewServer()
    pb.RegisterCSVProcessorServiceServer(s, &CSVProcessorService{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// 注意：此代码需要与protobuf文件一起使用，该protobuf文件定义了ProcessRequest和ProcessResponse消息类型。