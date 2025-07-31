// 代码生成时间: 2025-07-31 20:37:24
package main

import (
    "context"
    "log"
    "os"
    "path/filepath"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

// FileRenameService is the server API for batch file renaming.
type FileRenameService struct{}

// RenameFiles is a method to rename multiple files using provided mapping.
func (s *FileRenameService) RenameFiles(ctx context.Context, req *RenameRequest) (*emptypb.Empty, error) {
    for _, renameOp := range req.Operations {
        err := renameFile(renameOp.OldName, renameOp.NewName)
        if err != nil {
            // Handle error appropriately or return it to the client.
            return &emptypb.Empty{}, err
        }
    }
    return &emptypb.Empty{}, nil
}

// renameFile renames a single file from oldName to newName.
func renameFile(oldName, newName string) error {
    // Check if the old file exists.
    if _, err := os.Stat(oldName); os.IsNotExist(err) {
        return err
    }
    // Check if the new file name already exists.
    if _, err := os.Stat(newName); err == nil {
        return &os.PathError{Op: "rename", Path: newName, Err: os.ErrExist}
    }
    // Perform the rename operation.
    if err := os.Rename(oldName, newName); err != nil {
        return err
    }
    return nil
}

// main is the entry point for the batch rename tool.
func main() {
    // Set up gRPC server and listen on port 50051.
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    defer listener.Close()

    // Create a new server.
    server := grpc.NewServer()

    // Register the file rename service.
    RegisterFileRenameServiceServer(server, &FileRenameService{})

    // Start the server.
    log.Printf("server listening at %v", listener.Addr())
    if err := server.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RenameRequest is the request message for the RenameFiles method.
type RenameRequest struct {
    Operations []*RenameOperation
}

// RenameOperation is the operation message for file renaming.
type RenameOperation struct {
    OldName string
    NewName string
}

// RegisterFileRenameServiceServer registers the FileRenameServiceServer to a gRPC server.
func RegisterFileRenameServiceServer(s *grpc.Server, srv *FileRenameService) {
    RegisterFileRenameServiceServer(s, srv)
}
