// 代码生成时间: 2025-08-24 11:48:53
 * integration_test_tool.go
 * This file contains an example of a GRPC server and client for integration testing.
 */

package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc/test/bufconn"
)

// Define the TestService service
type TestService struct{}

// Define the TestServiceServer which implements TestServiceServer interface
type TestServiceServer struct{}

// SayHello implements TestServiceServer
func (s *TestServiceServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
    // Simple logic for demonstration purposes
    return &HelloResponse{Message: "Hello, " + in.Name}, nil
}

// HelloRequest is the request message for SayHello
type HelloRequest struct {
    Name string
}

// HelloResponse is the response message for SayHello
type HelloResponse struct {
    Message string
}

// TestServiceServer must embed UnimplementedTestServiceServer for forward compatibility
type UnimplementedTestServiceServer struct{}

func (*UnimplementedTestServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

// RegisterTestServiceServer registers a TestServiceServer to a gRPC server
func RegisterTestServiceServer(s *grpc.Server, srv *TestServiceServer) {
    RegisterTestServiceServer(s, srv)
}

// bufDialer creates a connection to a gRPC server
func bufDialer(ctx context.Context, addr string) (net.Conn, error) {
    return bufListener.Dialer(ctx, addr)
}

func main() {
    // Create a listener using a buffer
    bufListener := bufconn.Listen(1024 * 1024)
    server := grpc.NewServer()

    // Create a TestServiceServer and register it to the server
    testServiceServer := &TestServiceServer{}
    RegisterTestServiceServer(server, testServiceServer)
    reflection.Register(server)

    // Start the server
    go func() {
        if err := server.Serve(bufListener); err != nil {
            log.Fatalf("Failed to serve: %v", err)
        }
    }()

    // Create a client connection to the server using the buffer dialer
    conn, err := grpc.Dial("bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create a TestService client
    client := NewTestServiceClient(conn)

    // Make a test request
    r, err := client.SayHello(context.Background(), &HelloRequest{Name: "Integration Test"})
    if err != nil {
        log.Fatalf("Error when calling SayHello: %v", err)
    }
    fmt.Printf("Greeting: %s
", r.Message)
}
