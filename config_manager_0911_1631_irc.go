// 代码生成时间: 2025-09-11 16:31:43
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "os"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// Configuration defines the data structure for holding configuration data.
type Configuration struct {
    Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

// ConfigManagerServer defines the server interface for managing configurations.
type ConfigManagerServer struct {
    // embedded for protobuf's server implementation
    unimplementedConfigManagerServer int
}

// NewConfigManagerServer creates a new instance of ConfigManagerServer.
func NewConfigManagerServer() *ConfigManagerServer {
    return &ConfigManagerServer{}
}

// ReadConfig reads the configuration data from a file.
func (s *ConfigManagerServer) ReadConfig(ctx context.Context, in *emptypb.Empty) (*Configuration, error) {
    // Read configuration file content
    configFile, err := os.ReadFile("config.txt")
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
    // Return configuration data as a response
    return &Configuration{Data: string(configFile)}, nil
}

// WriteConfig writes the configuration data to a file.
func (s *ConfigManagerServer) WriteConfig(ctx context.Context, in *Configuration) (*emptypb.Empty, error) {
    // Write configuration data to a file
    err := ioutil.WriteFile("config.txt", []byte(in.Data), 0644)
    if err != nil {
        return nil, fmt.Errorf("failed to write config file: %w", err)
    }
    
    // Return empty response on success
    return &emptypb.Empty{}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")
    
    // Create a new server
    srv := grpc.NewServer()
    reflection.Register(srv)
    
    // Register the service with the server
    configManagerServer := NewConfigManagerServer()
    ConfigManagerServiceServer(srv).RegisterService(configManagerServer)
    
    // Start the server
    if err := srv.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}