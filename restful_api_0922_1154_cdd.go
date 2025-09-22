// 代码生成时间: 2025-09-22 11:54:29
package main

import (
    "context"
    "net/http"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// GRPCServer 模拟一个GRPC服务端
type GRPCServer struct {
    // 嵌入unimplementedService来实现接口
    unimplementedService UnimplementedServiceServer
}

// UnimplementedServiceServer 是服务接口
type UnimplementedServiceServer interface {
    GRPCMethod(context.Context, *Request) (*Response, error)
}

// Request 是GRPC方法的请求类型
type Request struct{
    // 定义请求字段
    Data string
}

// Response 是GRPC方法的响应类型
type Response struct{
    // 定义响应字段
    Message string
}

// GRPCMethod 是GRPC服务的方法实现
func (s *GRPCServer) GRPCMethod(ctx context.Context, req *Request) (*Response, error) {
    // 处理请求并返回响应
    return &Response{Message: "Hello from GRPC!"}, nil
}

// RestfulAPIServer 封装HTTP服务端
type RestfulAPIServer struct {
    // 嵌入http.Server结构体
    http.Server
}

// NewRestfulAPIServer 创建一个新的RestfulAPIServer实例
func NewRestfulAPIServer() *RestfulAPIServer {
    mux := http.NewServeMux()
    server := &RestfulAPIServer{
        Server: http.Server{
            Addr:    ":8080",
            Handler: mux,
        },
    }
    // 注册RESTful API接口
    mux.HandleFunc("/api/grpc", grpcHandler)
    return server
}

// grpcHandler 处理GRPC到RESTful API的转换
func grpcHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return

    }
    // 解析请求体
    var req Request
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // 调用GRPC方法
    grpcResponse, err := grpcServer.GRPCMethod(r.Context(), &req)
    if err != nil {
        grpcStatus, _ := status.FromError(err)
        http.Error(w, grpcStatus.Message(), int(grpcStatus.Code()))
        return
    }

    // 将GRPC响应转换为HTTP响应
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&grpcResponse)
}

// main 函数启动HTTP服务端
func main() {
    log.Println("Starting RESTful API server...")
    server := NewRestfulAPIServer()
    log.Fatal(server.ListenAndServe())
}
