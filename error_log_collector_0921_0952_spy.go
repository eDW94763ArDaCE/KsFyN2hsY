// 代码生成时间: 2025-09-21 09:52:01
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"
    "google.golang.org/grpc"
)

// ErrorLogService is a service for collecting error logs.
type ErrorLogService struct {
    // The channel for collecting logs.
    logChannel chan string
}

// NewErrorLogService creates a new instance of ErrorLogService.
func NewErrorLogService() *ErrorLogService {
    return &ErrorLogService{
        logChannel: make(chan string, 100), // Buffer size to hold logs
    }
}

// CollectLog is a method to collect error logs from clients.
func (s *ErrorLogService) CollectLog(ctx context.Context, in *LogRequest) (*LogResponse, error) {
    // Check if the incoming log is empty
    if in.GetLog() == "" {
        return nil, fmt.Errorf("log cannot be empty")
    }

    // Emit the log to the channel
    s.logChannel <- in.GetLog()

    // Acknowledge the receipt of the log
    return &LogResponse{Acknowledged: true}, nil
}

// Start starts the error log service.
func (s *ErrorLogService) Start(port string) {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    defer lis.Close()

    // Create a gRPC server
    s := grpc.NewServer()

    // Register the service
    RegisterErrorLogServiceServer(s, s)

    // Start serving requests
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Stop stops the error log service.
func (s *ErrorLogService) Stop() {
    close(s.logChannel)
}

// LogRequest is a message for sending error logs.
type LogRequest struct {
    Log string `protobuf:"bytes,1,opt,name=log,proto3"`
}

// LogResponse is a message for acknowledging log reception.
type LogResponse struct {
    Acknowledged bool `protobuf:"varint,1,opt,name=acknowledged,proto3"`
}

// ErrorLogServiceServer is the server API for ErrorLogService service.
type ErrorLogServiceServer interface {
    CollectLog(context.Context, *LogRequest) (*LogResponse, error)
}

// UnimplementedErrorLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedErrorLogServiceServer struct{}

// Must implement CollectLog.
func (*UnimplementedErrorLogServiceServer) CollectLog(
    context.Context, *LogRequest) (*LogResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method CollectLog not implemented")
}

// RegisterErrorLogServiceServer registers the service on the server.
func RegisterErrorLogServiceServer(s *grpc.Server, srv ErrorLogServiceServer) {
    s.RegisterService(&_ErrorLogService_serviceDesc, srv)
}

// _ErrorLogService_serviceDesc is the service descriptor for ErrorLogService.
var _ErrorLogService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "ErrorLogService",
    HandlerType: (*ErrorLogServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "CollectLog",
            Handler: _ErrorLogService_CollectLog_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "error_log_collector.proto",
}

// Main function to run the error log collector.
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Define the service
    service := NewErrorLogService()

    // Define the port to listen on
    port := ":50051"

    // Start the service
    go func() {
        service.Start(port)
    }()

    // Listen for termination signals
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Wait for signal
    <-sigs

    // Stop the service
    service.Stop()
    fmt.Println("Error log collector stopped")
}

// LogRequest and LogResponse are defined in the error_log_collector.proto file.
// The proto file should be compiled using the protoc command to generate Go code.

// The error_log_collector.proto file should be defined as follows:
// 
// syntax = "proto3";
// 
// package errorlog;
// 
// // The ErrorLogService provides methods for collecting error logs.
// service ErrorLogService {
//   // CollectLog allows clients to send error logs.
//   rpc CollectLog (LogRequest) returns (LogResponse);
// }
// 
// // A request message containing the log.
// message LogRequest {
//   string log = 1;
// }
// 
// // A response message acknowledging the log collection.
// message LogResponse {
//   bool acknowledged = 1;
// }

