// 代码生成时间: 2025-09-10 02:16:45
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

// AnalyzerService 服务定义
type AnalyzerService struct{}

// AnalyzeFile 分析文本文件
func (s *AnalyzerService) AnalyzeFile(ctx context.Context, in *AnalyzeFileRequest) (*AnalyzeFileResponse, error) {
    if in.GetFilename() == "" {
        return nil, fmt.Errorf("filename cannot be empty")
    }

    content, err := ioutil.ReadFile(in.GetFilename())
    if err != nil {
        if os.IsNotExist(err) {
            return nil, fmt.Errorf("file not found: %s", in.GetFilename())
        }
        return nil, fmt.Errorf("failed to read file: %s", err.Error())
    }

    response := &AnalyzeFileResponse{
        Filename: in.GetFilename(),
        Content:  string(content),
    }
    return response, nil
}

// AnalyzeFileRequest 文件分析请求
type AnalyzeFileRequest struct {
    Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

// AnalyzeFileResponse 文件分析响应
type AnalyzeFileResponse struct {
    Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
    Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

// main 函数启动gRPC服务器
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    s := grpc.NewServer()
    pb.RegisterAnalyzerServiceServer(s, &AnalyzerService{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
