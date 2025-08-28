// 代码生成时间: 2025-08-29 06:48:29
package main

import (
    "fmt"
    "google.golang.org/grpc"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
)

// FileSyncService provides the interface for file backup and synchronization.
type FileSyncService struct{}

// SyncFiles is a method that synchronizes files from source to destination.
func (f *FileSyncService) SyncFiles(ctx context.Context, req *FileSyncRequest) (*FileSyncResponse, error) {
    // Check for valid request
    if req == nil || req.SourceDirectory == "" || req.DestinationDirectory == "" {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid sync request")
    }

    // Ensure source directory exists
    if _, err := os.Stat(req.SourceDirectory); os.IsNotExist(err) {
        return nil, status.Errorf(codes.NotFound, "Source directory not found")
    }

    // Ensure destination directory exists or create it
    if _, err := os.Stat(req.DestinationDirectory); os.IsNotExist(err) {
        if err := os.MkdirAll(req.DestinationDirectory, 0755); err != nil {
            return nil, status.Errorf(codes.Internal, "Failed to create destination directory")
        }
    }

    // Sync files
    err := filepath.Walk(req.SourceDirectory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        relPath, err := filepath.Rel(req.SourceDirectory, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(req.DestinationDirectory, relPath)
        if info.IsDir() {
            return os.MkdirAll(destPath, 0755)
        } else {
            return copyFile(path, destPath)
        }
    })

    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to sync files: %v", err)
    }

    return &FileSyncResponse{
        Success: true,
        Message: "Files synced successfully",
    }, nil
}

// copyFile copies a file from source to destination.
func copyFile(src, dst string) error {
    data, err := ioutil.ReadFile(src)
    if err != nil {
        return err
    }
    if err := ioutil.WriteFile(dst, data, 0644); err != nil {
        return err
    }
    return nil
}

func main() {
    // Define the server address
    serverAddress := "localhost:50051"

    // Create a new gRPC server
    lis, err := net.Listen("tcp", serverAddress)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Printf("Listening on %s", serverAddress)

    // Create a new FileSyncService
    syncService := &FileSyncService{}

    // Create a new gRPC server and register the FileSyncService
    s := grpc.NewServer()
    pb.RegisterFileSyncServer(s, syncService)

    // Serve the server
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Protobuf definitions would be in a separate file, typically with a .proto extension,
// and compiled into Go code using the protoc compiler with the Go plugin.
// Here is a simplified illustration of what the .proto file might contain:

// syntax = "proto3";
//
// package file_sync;
//
// service FileSync {
//   rpc SyncFiles(FileSyncRequest) returns (FileSyncResponse) {};
// }
//
// message FileSyncRequest {
//   string sourceDirectory = 1;
//   string destinationDirectory = 2;
// }
//
// message FileSyncResponse {
//   bool success = 1;
//   string message = 2;
// }
