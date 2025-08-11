// 代码生成时间: 2025-08-12 07:05:39
package main

import (
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// FormValidatorService 用于表单验证的服务
type FormValidatorService struct{}

// ValidateForm 验证表单数据
func (s *FormValidatorService) ValidateForm(ctx grpc.Context, req *FormRequest) (*emptypb.Empty, error) {
    // 检查表单中的字段是否符合要求
    if err := Validate(req); err != nil {
        // 如果验证失败，返回错误
        return nil, status.Errorf(codes.InvalidArgument, "Validation failed: %v", err)
    }
    // 如果验证成功，返回空对象
    return &emptypb.Empty{}, nil
}

// Validate 验证表单请求中的每个字段
func Validate(req *FormRequest) error {
    // 示例：验证字段 'email'
    if req.Email == "" {
        return fmt.Errorf("email is required")
    }
    // 示例：验证字段 'username'
    if req.Username == "" {
        return fmt.Errorf("username is required")
    }
    // 可以添加更多的字段验证逻辑
    // ...
    return nil
}

// FormRequest 是一个包含表单数据的请求结构
type FormRequest struct {
    Email    string
    Username string
    // 可以添加更多的字段
    // ...
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")
    s := grpc.NewServer()
    // 注册服务
    pb.RegisterFormValidatorServer(s, &FormValidatorService{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// 要使用这个服务，你需要定义proto文件，并使用protoc生成Go代码。下面的代码是一个简单的示例，展示了如何定义proto文件。
// service.proto
//
// syntax = "proto3";
//
// package pb;
//
// import "google/protobuf/empty.proto";
//
// message FormRequest {
//     string email = 1;
//     string username = 2;
// }
//
// service FormValidator {
//     rpc ValidateForm (FormRequest) returns (google.protobuf.Empty) {};
// }