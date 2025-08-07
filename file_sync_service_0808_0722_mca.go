// 代码生成时间: 2025-08-08 07:22:43
package main

import (
    "context"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os"
    "path/filepath"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// FileSyncService defines the service for file backup and sync
type FileSyncService struct {
    // Other service properties can be included here
}

// FileSyncServiceServer is the server API for FileSyncService service
type FileSyncServiceServer struct {
    FileSyncService
}

// Proto message for file sync request
type FileSyncRequest struct {
    SourcePath string `protobuf:"bytes,1,opt,name=source_path,json=sourcePath,proto3"`
    DestinationPath string `protobuf:"bytes,2,opt,name=destination_path,json=destinationPath,proto3"`
}

// Proto message for file sync response
type FileSyncResponse struct {
    Success bool   `protobuf:"varint,1,opt,name=success,proto3"`
    Message string `protobuf:"bytes,2,opt,name=message,proto3"`
}

// SyncFile is called by the client to initiate file sync
func (s *FileSyncServiceServer) SyncFile(ctx context.Context, req *FileSyncRequest) (*FileSyncResponse, error) {
    // Check if source path is valid
    if _, err := os.Stat(req.SourcePath); os.IsNotExist(err) {
        return nil, status.Errorf(codes.NotFound, "Source file not found: %s", req.SourcePath)
    }

    // Ensure destination directory exists
    if err := os.MkdirAll(filepath.Dir(req.DestinationPath), 0755); err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to create destination directory: %s", err)
    }

    // Perform file copy
    if err := copyFile(req.SourcePath, req.DestinationPath); err != nil {
        return nil, status.Errorf(codes.Internal, "File sync failed: %s", err)
    }

    return &FileSyncResponse{Success: true, Message: "File sync successful"}, nil
}

// copyFile copies a file from src to dst. If dst already exists, it will be overwritten.
func copyFile(src, dst string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
