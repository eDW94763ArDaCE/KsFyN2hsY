// 代码生成时间: 2025-08-10 22:06:25
package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "path/filepath"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc/status"
)

// Define the LogEntry message structure
type LogEntry struct {
    Timestamp int64  "protobuf:"varint,1,opt,name=timestamp" json:"timestamp""
    Level     string "protobuf:"varint,2,opt,name=level" json:"level""
    Message   string "protobuf:"varint,3,opt,name=message" json:"message""
}

// Define the ErrorLogService service
type ErrorLogService struct {
    // Add other properties here if needed
}

// Define the ErrorLogServiceServer interface
type ErrorLogServiceServer interface {
    LogError(ctx context.Context, in *LogEntry) (*Empty, error)
}

// RegisterErrorLogServiceServer registers the ErrorLogServiceServer interface with the gRPC server
func RegisterErrorLogServiceServer(s *grpc.Server, srv ErrorLogServiceServer) {
    RegisterErrorLogServiceServer(s, srv)
}

// ErrorLogService implementation
func (e *ErrorLogService) LogError(ctx context.Context, in *LogEntry) (*Empty, error) {
    // Implement error logging logic
    fmt.Printf("Received log entry: %+v
", in)
    // Save the log entry to a file or database
    err := saveLogEntryToFile(in)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to save log entry: %v", err)
    }
    return &Empty{}, nil
}

// saveLogEntryToFile writes the log entry to a file
func saveLogEntryToFile(entry *LogEntry) error {
    // Define the log file path
    logFilePath := "error_logs.log"
    // Open the file for append
    file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    // Write the log entry to the file
    _, err = file.WriteString(fmt.Sprintf("%d %s %s
", entry.Timestamp, entry.Level, entry.Message))
    return err
}

func main() {
    // Define the server address
    serverAddress := ":50051"

    // Create a new gRPC server
    lis, err := net.Listen("tcp", serverAddress)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    fmt.Printf("Server listening on %s
", serverAddress)

    // Create a new gRPC server instance
    s := grpc.NewServer()

    // Register the ErrorLogServiceServer with the gRPC server
    RegisterErrorLogServiceServer(s, &ErrorLogService{})

    // Register reflection service on gRPC server.
    reflection.Register(s)

    // Serve the gRPC server
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
