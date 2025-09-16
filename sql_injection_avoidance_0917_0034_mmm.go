// 代码生成时间: 2025-09-17 00:34:11
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "strings"
    "time"

    _ "github.com/go-sql-driver/mysql"  // MySQL driver
    "google.golang.org/grpc"
)

// SQLInjectionService 定义了一个防止SQL注入的服务
type SQLInjectionService struct{}

// 防止SQL注入的一个简单方法就是使用参数化查询
// 以下为gRPC服务端的实现

// GetUserInfo 是一个示例方法，它接受用户ID，并返回用户信息
func (s *SQLInjectionService) GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (*UserInfoResponse, error) {
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        return nil, err
    }
    defer db.Close()

    // 使用参数化查询防止SQL注入
    var userID int64
    var username string
    var email string
    query := `SELECT id, username, email FROM users WHERE id = ?`
    err = db.QueryRow(query, req.UserId).Scan(&userID, &username, &email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, grpc.Errorf(codes.NotFound, "User not found")
        }
        return nil, grpc.Errorf(codes.Internal, "Error fetching user info: %v", err)
    }

    // 构造响应
    return &UserInfoResponse{Username: username, Email: email}, nil
}

// main 函数初始化gRPC服务并启动服务端
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    sqlInjectionService := &SQLInjectionService{}

    // 注册服务
    RegisterSQLInjectionServiceServer(grpcServer, sqlInjectionService)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

// 数据结构定义
type GetUserInfoRequest struct {
    UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId"`
}

type UserInfoResponse struct {
    Username string `protobuf:"string,1,opt,name=username"`
    Email    string `protobuf:"string,2,opt,name=email"`
}
