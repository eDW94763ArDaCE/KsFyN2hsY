// 代码生成时间: 2025-08-07 19:17:15
This file defines a simple GRPC service for a user interface component library.
It includes a service definition and a server implementation.
*/

package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// Define our own errors for better clarity
var (
    ErrComponentNotFound = grpc.Errorf(codes.NotFound, "Component not found")
)

// ComponentLibraryService represents the server that will handle requests.
type ComponentLibraryService struct{}

// Define the service
type ComponentLibraryServer interface {
    GetComponent(context.Context, *ComponentRequest) (*Component, error)
    ListComponents(context.Context, *emptypb.Empty) (*ComponentList, error)
}

// Implement the service
func (s *ComponentLibraryService) GetComponent(ctx context.Context, req *ComponentRequest) (*Component, error) {
    // Check if the component is registered
    if _, ok := registeredComponents[req.Name]; !ok {
        return nil, ErrComponentNotFound
    }

    // Return a copy of the component to prevent modification
    return registeredComponents[req.Name], nil
}

func (s *ComponentLibraryService) ListComponents(ctx context.Context, _ *emptypb.Empty) (*ComponentList, error) {
    var components []*Component
    for _, component := range registeredComponents {
        components = append(components, component)
    }
    return &ComponentList{Components: components}, nil
}

// Define the protobuf messages
type ComponentRequest struct {
    Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

type Component struct {
    Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
    Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
    // Other fields can be added here as needed
}

type ComponentList struct {
    Components []*Component `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
}

// registeredComponents is a map of the components we have registered.
var registeredComponents = map[string]*Component{
    "Button":   {"Name": "Button", "Version": "1.0.0"},
    "Textbox": {"Name": "Textbox", "Version": "1.0.0"},
    // Add more components here
}

// main function to start the server
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    // Register the service
    grpcServer.RegisterService(&_ComponentLibrary_serviceDesc, &ComponentLibraryService{})
    // Register reflection service on gRPC server.
    reflection.Register(grpcServer)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
