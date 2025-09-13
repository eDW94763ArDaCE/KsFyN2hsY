// 代码生成时间: 2025-09-14 04:02:47
package main

import (
    "context"
    "log"
    "net"
    "os"
    "os/exec"
    "syscall"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
    "github.com/golang/protobuf/ptypes"
)

// ProcessManagerService defines the gRPC service interface for managing processes.
type ProcessManagerService struct{}

// StartProcess starts a process with the specified command.
func (s *ProcessManagerService) StartProcess(ctx context.Context, in *StartProcessRequest) (*emptypb.Empty, error) {
    if in == nil || in.Command == "" {
        return nil, status.Errorf(codes.InvalidArgument, "Command must be provided")
    }

    // Split the command into individual parts.
    args := strings.Split(in.Command, " ")
    if len(args) == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid command format")
    }

    // Start the process.
    process := exec.Command(args[0], args[1:]...)
    if err := process.Start(); err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to start process: %v", err)
    }

    log.Printf("Process started with command: %s", in.Command)
    return &emptypb.Empty{}, nil
}

// StopProcess stops a process with the specified ID.
func (s *ProcessManagerService) StopProcess(ctx context.Context, in *StopProcessRequest) (*emptypb.Empty, error) {
    if in == nil || in.ProcessID == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Process ID must be provided")
    }

    process, err := os.FindProcess(in.ProcessID)
    if err != nil {
        return nil, status.Errorf(codes.NotFound, "Process not found: %v", err)
    }

    if err := process.Signal(syscall.SIGTERM); err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to stop process: %v", err)
    }

    log.Printf("Process stopped with ID: %d", in.ProcessID)
    return &emptypb.Empty{}, nil
}

// GetProcessStatus retrieves the status of a process with the specified ID.
func (s *ProcessManagerService) GetProcessStatus(ctx context.Context, in *GetProcessStatusRequest) (*GetProcessStatusResponse, error)
    if in == nil || in.ProcessID == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Process ID must be provided\)
    }

    process, err := os.FindProcess(in.ProcessID)
    if err != nil {
        return nil, status.Errorf(codes.NotFound, "Process not found: %v", err)
    }

    state, err := process.Wait()
    if err != nil && !errors.Is(err, os.ErrProcessDone) {
        return nil, status.Errorf(codes.Internal, "Failed to retrieve process status: %v", err)
    }

    exitCode := 0
    if exitErr, ok := state.Sys().(syscall.WaitStatus); ok && exitErr.Exited() {
        exitCode = exitErr.ExitStatus()
    }

    return &GetProcessStatusResponse{ExitCode: int32(exitCode), State: state.String()}, nil
}

func main() {
    server := grpc.NewServer()
    processManagerService := &ProcessManagerService{}

    // Register the service with the gRPC server.
    RegisterProcessManagerServiceServer(server, processManagerService)
    reflection.Register(server)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    log.Printf("Server listening on port %s", "50051")
    if err := server.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

// Define Protobuf messages and service for process management.

// StartProcessRequest is the request message for starting a process.
type StartProcessRequest struct {
    Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

// StopProcessRequest is the request message for stopping a process.
type StopProcessRequest struct {
    ProcessID int32 `protobuf:"varint,1,opt,name=processID,proto3" json:"processID,omitempty"`
}

// GetProcessStatusRequest is the request message for getting the status of a process.
type GetProcessStatusRequest struct {
    ProcessID int32 `protobuf:"varint,1,opt,name=processID,proto3" json:"processID,omitempty"`
}

// GetProcessStatusResponse is the response message for getting the status of a process.
type GetProcessStatusResponse struct {
    ExitCode int32 `protobuf:"varint,1,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
    State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

// RegisterProcessManagerServiceServer registers the service with the gRPC server.
func RegisterProcessManagerServiceServer(s *grpc.Server, srv *ProcessManagerService) {
    RegisterProcessManagerServiceServer(s, srv)
}
