// 代码生成时间: 2025-08-09 23:35:31
package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/protobuf/types/known/timestamppb"
)

// AuditLog represents a single audit log entry
type AuditLog struct {
    Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
    Event      string               `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
    Details    string               `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

// AuditLogServiceServer defines the gRPC server methods for audit logging
type AuditLogServiceServer struct {
    // Embed the default server implementation
    UnimplementedAuditLogServiceServer 
}

// NewAuditLogServiceServer creates a new instance of the AuditLogServiceServer
func NewAuditLogServiceServer() *AuditLogServiceServer {
    return &AuditLogServiceServer{}
}

// LogEvent logs a new audit event
func (s *AuditLogServiceServer) LogEvent(ctx context.Context, req *LogEventRequest) (*emptypb.Empty, error) {
    // Validate the request
    if req.Event == nil {
        return nil, status.Errorf(codes.InvalidArgument, "Event must be provided")
    }

    // Create a new audit log entry
    logEntry := &AuditLog{
        Timestamp: timestamppb.Now(),
        Event:     req.Event.Name,
        Details:   req.Event.Details,
    }

    // Write the audit log entry to a file or database (implementation depends on requirements)
    // For simplicity, we'll just write to a file
    err := writeLogToFile(logEntry)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to write log entry: %v", err)
    }

    return &emptypb.Empty{}, nil
}

// LogEventRequest represents the request for logging an event
type LogEventRequest struct {
    Event *AuditLog `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

// AuditLogService provides the interface for the audit logging service
type AuditLogService interface {
    LogEvent(ctx context.Context, req *LogEventRequest) (*emptypb.Empty, error)
}

// writeLogToFile writes the audit log entry to a file
func writeLogToFile(logEntry *AuditLog) error {
    file, err := os.OpenFile("audit_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(fmt.Sprintf("%s - %s - %s
", logEntry.Timestamp.Seconds, logEntry.Event, logEntry.Details))
    return err
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Create a gRPC server
    grpcServer := grpc.NewServer()

    // Register the audit log service
    auditLogService := NewAuditLogServiceServer()
    pb.RegisterAuditLogServiceServer(grpcServer, auditLogService)

    // Register reflection service on gRPC server.
    reflection.Register(grpcServer)

    // Start serving
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}