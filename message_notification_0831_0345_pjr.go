// 代码生成时间: 2025-08-31 03:45:50
package main

import (
    "context"
    "fmt"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "pb" // 假设pb是消息通知系统的proto文件生成的Go文件
)

// NotificationService 是一个实现了消息通知服务的gRPC服务器
type NotificationService struct {
    // 可以在这里添加更多的字段，比如数据库连接，缓存等
}

// NewNotificationService 创建一个新的NotificationService实例
func NewNotificationService() *NotificationService {
    return &NotificationService{}
}

// SendNotification 实现了消息通知的gRPC方法
func (s *NotificationService) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
    // 检查请求参数
    if req.GetMessage() == "" {
        return nil, status.Error(codes.InvalidArgument, "Message cannot be empty")
    }

    // 模拟消息发送逻辑
    fmt.Println("Sending notification: ", req.GetMessage())

    // 正常情况下，这里会是与消息队列或者数据库的交互
    // 例如：err := SendMessageToQueue(req.GetMessage())
    // if err != nil {
    //     return nil, status.Error(codes.Internal, "Failed to send message")
    // }

    // 返回成功响应
    return &pb.SendNotificationResponse{Status: "success"}, nil
}

func main() {
    // 监听端口
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建gRPC服务器
    s := grpc.NewServer()
    pb.RegisterNotificationServiceServer(s, NewNotificationService())

    // 启动服务器
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
