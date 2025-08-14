// 代码生成时间: 2025-08-14 08:44:58
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
    "gopkg.in/yaml.v2"
    "google.golang.org/grpc"
)

// LogEntry defines the structure for a log entry.
type LogEntry struct {
    Timestamp time.Time `yaml:"timestamp"`
    Level     string    `yaml:"level"`
    Message   string    `yaml:"message"`
}

// LogParserService is the service that will handle log parsing.
type LogParserService struct{}

// ParseLog parses a log file and returns structured log entries.
func (s *LogParserService) ParseLog(ctx context.Context, in *ParseLogRequest) (*ParseLogResponse, error) {
    if in.LogFile == "" {
        return nil, status.Errorf(codes.InvalidArgument, "file path cannot be empty")
    }

    data, err := ioutil.ReadFile(in.LogFile)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to read log file: %v", err)
    }

    logEntries := make([]*LogEntry, 0)
    for _, line := range strings.Split(string(data), "
") {
        entry := &LogEntry{}
        if err := yaml.Unmarshal([]byte(line), entry); err != nil {
            log.Printf("Failed to parse log entry: %s", line)
            continue
        }
        logEntries = append(logEntries, entry)
    }

    return &ParseLogResponse{Entries: logEntries}, nil
}

// ParseLogRequest is the request message for ParseLog.
type ParseLogRequest struct {
    LogFile string `protobuf:"bytes,1,opt,name=log_file,json=LogFile" json:"log_file,omitempty"`
}

// ParseLogResponse is the response message for ParseLog.
type ParseLogResponse struct {
    Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

// RegisterService registers the LogParserService with the gRPC server.
func RegisterService(s *grpc.Server, svc *LogParserService) {
    pb.RegisterLogParserServer(s, svc)
}

func main() {
    // Define the server and service
    server := grpc.NewServer()
    service := &LogParserService{}

    // Register the service with the server
    RegisterService(server, service)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    // Start serving requests
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
