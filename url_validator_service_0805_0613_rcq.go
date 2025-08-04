// 代码生成时间: 2025-08-05 06:13:37
package main

import (
    "context"
    "fmt"
    "net/url"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
)

// URLValidationService is the server API for URLValidation service.
type URLValidationService struct{}

// ValidateURL checks if the provided URL is valid.
func (s *URLValidationService) ValidateURL(ctx context.Context, req *URLValidationRequest) (*URLValidationResponse, error) {
    // Parse the URL to validate its structure.
    u, err := url.ParseRequestURI(req.Url)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "invalid URL: %v", err)
    }

    // Check for the presence of both scheme and host to consider a URL valid.
    if u.Scheme == "" || u.Host == "" {
        return &URLValidationResponse{Valid: false}, nil
    }

    // If all checks pass, return a valid response.
    return &URLValidationResponse{Valid: true}, nil
}

// URLValidationRequest is the request message for the ValidateURL method.
type URLValidationRequest struct {
    Url string `protobuf:"bytes,1,opt,name=url,proto3"`
}

// URLValidationResponse is the response message for the ValidateURL method.
type URLValidationResponse struct {
    Valid bool `protobuf:"varint,1,opt,name=valid,proto3"`
}

func main() {
    // Set up the gRPC server.
    server := grpc.NewServer()
    // Register the URLValidationService on the gRPC server.
    pb.RegisterURLValidationServiceServer(server, &URLValidationService{})

    // Listen on port 50051 for incoming connections.
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
