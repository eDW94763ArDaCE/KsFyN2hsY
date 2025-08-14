// 代码生成时间: 2025-08-14 16:39:33
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "log"
)

// AccessControlService 定义访问控制服务
type AccessControlService struct {
    // 可以添加其他属性，比如用户权限数据库等
}

// CheckPermission 实现权限检查方法
func (service *AccessControlService) CheckPermission(ctx context.Context, req *PermissionRequest) (*PermissionResponse, error) {
    // 这里添加权限检查逻辑，例如验证用户token等
    // 以下仅为示例，实际逻辑需要根据具体需求编写
    if req.Token == "valid_token" {
        return &PermissionResponse{Allowed: true}, nil
    } else {
        return nil, fmt.Errorf("unauthorized access")
    }
}

// PermissionRequest 定义权限请求参数
type PermissionRequest struct {
    Token string
}

// PermissionResponse 定义权限响应结果
type PermissionResponse struct {
    Allowed bool
}

func main() {
    // 监听本地端口
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建gRPC服务器
    s := grpc.NewServer()

    // 注册访问控制服务
    s.RegisterService(&_AccessControlService_serviceDesc, &AccessControlService{})

    // 启动服务器
    fmt.Println("server starting")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// 以下是gRPC服务定义和注册的伪代码，需要根据实际protobuf定义生成
var _AccessControlService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "AccessControlService",
    HandlerType: (*AccessControlService)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "CheckPermission",
            Handler: _AccessControlService_CheckPermission_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "metadata",
}

// _AccessControlService_CheckPermission_Handler 是CheckPermission方法的handler
func _AccessControlService_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(PermissionRequest)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(AccessControlService).CheckPermission(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
         FullMethod: "/AccessControlService/CheckPermission",
    }
    
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(AccessControlService).CheckPermission(ctx, req.(*PermissionRequest))
    }
    return interceptor(ctx, in, info, handler)
}