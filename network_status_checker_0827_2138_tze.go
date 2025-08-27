// 代码生成时间: 2025-08-27 21:38:57
// network_status_checker.go
// 该程序是一个基于GRPC框架的网络连接状态检查器
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// NetworkStatusCheckerService 定义检查网络连接的服务
type NetworkStatusCheckerService struct{}

// CheckConnection 检查指定主机的网络连接状态
func (s *NetworkStatusCheckerService) CheckConnection(ctx context.Context, in *ConnectionRequest) (*ConnectionResponse, error) {
	// 尝试拨号
	conn, err := net.Dial("tcp", in.Host)
	if err != nil {
		// 如果拨号失败，返回错误状态
		return nil, status.Errorf(codes.Unavailable, "failed to connect to host: %v", err)
	}
	defer conn.Close()

	// 如果拨号成功，返回成功状态
	return &ConnectionResponse{
		Success: true,
		Message: "Connection established successfully",
	}, nil
}

// ConnectionRequest 定义检查网络连接的请求参数
type ConnectionRequest struct {
	Host string
}

// ConnectionResponse 定义检查网络连接的响应参数
type ConnectionResponse struct {
	Success bool   
	Message string
}

// serviceDesc 是服务描述
var serviceDesc = grpc.ServiceDesc{
	ServiceName: "networkstatus.NetworkStatusChecker",
	HandlerType: (*NetworkStatusCheckerService)(nil),
	Methods: []grpc.MethodDesc{
	{
		MethodName: "CheckConnection",
		Handler: grpcUnaryServerInterceptor(
			func(ctx context.Context, req interface{}) (interface{}, error) {
				return srv.CheckConnection(ctx, req.(*ConnectionRequest))
			}),
	},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network_status_checker.proto",
}

// RegisterNetworkStatusCheckerService 注册服务
func RegisterNetworkStatusCheckerService(s *grpc.Server, srv *NetworkStatusCheckerService) {
	grpc.RegisterServiceDesc(s, &serviceDesc, srv)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterNetworkStatusCheckerService(grpcServer, &NetworkStatusCheckerService{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
