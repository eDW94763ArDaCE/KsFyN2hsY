// 代码生成时间: 2025-08-17 19:56:26
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "io"
    "log"
    "net"
    "os"
    "time"
)

// FileService is the server API for backup and sync operations.
type FileService struct {
    // Embedded struct to implement the gRPC interface.
    UnimplementedFileServiceServer
}

// File represents the data structure for file information.
type File struct {
    Name    string
    Content string
}

// Request is the data structure for handling requests.
type Request struct {
    Source string
    Destination string
}

// Response is the data structure for handling responses.
type Response struct {
    Success bool
    Message string
}

// NewFileService creates a new instance of the FileService.
func NewFileService() *FileService {
    return &FileService{}
}

// BackupFile handles the backup of a file from source to destination.
func (s *FileService) BackupFile(ctx context.Context, in *Request) (*Response, error) {
    // Open the source file.
    srcFile, err := os.Open(in.Source)
    if err != nil {
        return nil, err
    }
    defer srcFile.Close()

    // Create the destination file.
    dstFile, err := os.Create(in.Destination)
    if err != nil {
        return nil, err
    }
    defer dstFile.Close()

    // Copy file content from source to destination.
    if _, err := io.Copy(dstFile, srcFile); err != nil {
        return nil, err
    }

    // Return a success response.
    return &Response{Success: true, Message: "Backup completed successfully."}, nil
}

// SyncFiles handles the synchronization of files between source and destination.
func (s *FileService) SyncFiles(ctx context.Context, in *Request) (*Response, error) {
    // Check if the source file exists.
    if _, err := os.Stat(in.Source); os.IsNotExist(err) {
        return nil, fmt.Errorf("source file does not exist: %w", err)
    }

    // Check if the destination file exists.
    if _, err := os.Stat(in.Destination); os.IsNotExist(err) {
        // If the destination file does not exist, create it.
        if _, err := os.Create(in.Destination); err != nil {
            return nil, err
        }
    }

    // Compare file modification times and copy if necessary.
    srcModTime, err := os.Stat(in.Source).ModTime()
    if err != nil {
        return nil, err
    }
    dstModTime, err := os.Stat(in.Destination).ModTime()
    if err != nil {
        return nil, err
    }

    if srcModTime.After(dstModTime) {
        return s.BackupFile(ctx, in)
    }

    // Return a success response if no action needed.
    return &Response{Success: true, Message: "Files are already in sync."}, nil
}

// RunServer starts the gRPC server on the provided address.
func RunServer(address string) error {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        return err
    }
    fmt.Println("Server listening on", address)

    grpcServer := grpc.NewServer()
    // Register the FileService on the server.
    RegisterFileServiceServer(grpcServer, NewFileService())

    // Serve the server.
    if err := grpcServer.Serve(lis); err != nil {
        return err
    }
    return nil
}

func main() {
    // Run the server on the specified address.
    address := ":50051"
    if err := RunServer(address); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
