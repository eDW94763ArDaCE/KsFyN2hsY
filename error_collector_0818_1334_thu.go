// 代码生成时间: 2025-08-18 13:34:00
package main

import (
    "fmt"
    "io"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// ErrorLogService defines the gRPC service for error logging.
type ErrorLogService struct{}

// LogError logs an error message to the service.
func (s *ErrorLogService) LogError(ctx context.Context, req *LogErrorRequest) (*LogErrorResponse, error) {
    // Check if the request is valid.
    if req == nil || req.ErrorMessage == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request")
    }

    // Log the error message to the standard logger.
    log.Printf("Error logged: %s", req.ErrorMessage)

    // Return a success response.
    return &LogErrorResponse{Success: true}, nil
}

// LogErrorRequest is the request message for the LogError RPC.
type LogErrorRequest struct {
    ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage"`
}

// LogErrorResponse is the response message for the LogError RPC.
type LogErrorResponse struct {
    Success bool `protobuf:"varint,1,opt,name=success"`
}

func main() {
    // Define the server address.
    serverAddress := ":50051"

    // Create a new gRPC server.
    lis, err := net.Listen("tcp", serverAddress)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Printf("Listening on %s", serverAddress)

    // Create a new gRPC server instance.
    server := grpc.NewServer()

    // Register the error log service.
    RegisterErrorLogServiceServer(server, &ErrorLogService{})

    // Start the server.
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterErrorLogServiceServer registers the error log service to the gRPC server.
func RegisterErrorLogServiceServer(s *grpc.Server, srv ErrorLogServiceServer) {
    s.RegisterService(&_ErrorLogService_serviceDesc, srv)
}

// _ErrorLogService_serviceDesc is the service descriptor for ErrorLogService.
var _ErrorLogService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "ErrorLogService",
    HandlerType: (*ErrorLogServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "LogError",
            Handler: _ErrorLogService_LogError_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "error_collector.proto",
}

// ErrorLogServiceServer must be implemented by the user.
type ErrorLogServiceServer interface {
    LogError(context.Context, *LogErrorRequest) (*LogErrorResponse, error)
}

// _ErrorLogService_LogError_Handler is the handler for the LogError method.
func _ErrorLogService_LogError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(LogErrorRequest)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil { return srv.(ErrorLogServiceServer).LogError(ctx, in) }
    info := &grpc.UnaryServerInfo{
        Server: srv,
         FullMethod: "/ErrorLogService/LogError",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(ErrorLogServiceServer).LogError(ctx, req.(*LogErrorRequest))
    }
    return interceptor(ctx, in, info, handler)
}
