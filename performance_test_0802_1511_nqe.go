// 代码生成时间: 2025-08-02 15:11:20
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// 定义服务端和客户端之间通信的RPC服务
type BenchmarkServiceClient interface {
    Ping(context.Context, *PingRequest, ...grpc.CallOption) (*PingResponse, error)
}

// RPC请求和响应的protobuf定义
type PingRequest struct {
    Message string
}

type PingResponse struct {
    Message string
}

func main() {
    // 服务端地址
    serverAddress := "localhost:50051"
    // 创建连接
    conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // 创建客户端
    client := NewBenchmarkServiceClient(conn)
    
    // 性能测试参数
    testCount := 10000
    start := time.Now()
    for i := 0; i < testCount; i++ {
        _, err := client.Ping(context.Background(), &PingRequest{Message: "Hello"}, grpc.WaitForReady(true))
        if err != nil {
            if status.Code(err) == codes.Unavailable {
                fmt.Println("Server is unavailable")
                return
            }
            log.Fatalf("could not ping: %v", err)
        }
    }
    // 计算并输出性能测试结果
    duration := time.Since(start)
    fmt.Printf("Performed %d pings in %s
", testCount, duration)
    fmt.Printf("Average ping time: %f ms
", float64(duration.Milliseconds())/float64(testCount))
}

// BenchmarkServiceClient是BenchmarkServiceClient接口的实现
type benchmarkServiceClient struct {
    cc *grpc.ClientConn
}

func NewBenchmarkServiceClient(cc *grpc.ClientConn) BenchmarkServiceClient {
    return &benchmarkServiceClient{cc}
}

func (c *benchmarkServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
    out := new(PingResponse)
    err := grpc.Invoke(ctx, "/BenchmarkService/Ping", in, out, c.cc, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}
