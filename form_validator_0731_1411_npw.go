// 代码生成时间: 2025-07-31 14:11:57
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// FormValidationService 定义了表单验证器服务
type FormValidationService struct{}

// ValidateForm 实现验证表单数据的方法
func (s *FormValidationService) ValidateForm(ctx context.Context, req *FormRequest) (*FormResponse, error) {
    // 检查表单数据是否为空
    if req.GetName() == "" || req.GetEmail() == "" || req.GetAge() == 0 {
        return nil, fmt.Errorf("missing required form fields")
    }
    // 检查姓名是否合法，这里只是简单示例，实际中可能需要更复杂的规则
    if len(req.GetName()) < 2 || len(req.GetName()) > 50 {
        return nil, fmt.Errorf("invalid name length")
    }
    // 检查邮箱是否合法，这里使用简单的正则表达式作为示例
    if ok, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, req.GetEmail()); !ok {
        return nil, fmt.Errorf("invalid email format")
    }
    // 检查年龄是否合法
    if req.GetAge() < 0 || req.GetAge() > 120 {
        return nil, fmt.Errorf("invalid age")
    }
    // 如果所有验证通过，则返回成功响应
    return &FormResponse{Success: true}, nil
}

// FormRequest 定义了表单请求的proto消息
type FormRequest struct {
    Name  string `protobuf:"bytes,1,opt,name=name,proto3"`
    Email string `protobuf:"bytes,2,opt,name=email,proto3"`
    Age   int32  `protobuf:"varint,3,opt,name=age,proto3"`
}

// FormResponse 定义了表单响应的proto消息
type FormResponse struct {
    Success bool `protobuf:"varint,1,opt,name=success,proto3"`
}

// main 函数初始化并运行gRPC服务
func main() {
   lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")
    s := grpc.NewServer()
    // 注册服务
    RegisterFormValidationServiceServer(s, &FormValidationService{})
    // 注册gRPC反射服务
    reflection.Register(s)
    // 启动服务
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
