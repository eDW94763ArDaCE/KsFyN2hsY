// 代码生成时间: 2025-10-02 03:04:22
package main

import (
    "context"
    "fmt"
    "net/url"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// UrlValidatorService is the server API for UrlValidator service.
type UrlValidatorService struct{
}

// Validate checks if the given URL is valid.
func (s *UrlValidatorService) Validate(ctx context.Context, in *ValidateRequest) (*ValidateResponse, error) {
    // Parse the URL to check its validity.
    u, err := url.ParseRequestURI(in.Url)
    if err != nil {
        // If the URL cannot be parsed, return an error with INVALID_ARGUMENT code.
        return nil, status.Errorf(codes.InvalidArgument, "invalid URL: %v", err)
    }
    
    // Check if the URL scheme is valid (e.g., http, https).
    if u.Scheme != "http" && u.Scheme != "https" {
        return nil, status.Errorf(codes.InvalidArgument, "URL scheme must be http or https")
    }

    // Return a success response indicating the URL is valid.
    return &ValidateResponse{Valid: true}, nil
}

// ValidateRequest is the request message for the Validate method.
type ValidateRequest struct {
    Url string
}

// ValidateResponse is the response message for the Validate method.
type ValidateResponse struct {
    Valid bool
}

func main() {
    // Define the server.
    server := grpc.NewServer()
    
    // Create a new UriValidatorService instance.
    urlValidatorService := &UrlValidatorService{}
    
    // Register the service with the server.
    // Assuming a protocol buffer definition exists for the UrlValidatorService
    // and it's been compiled and registered correctly.
    
    // Start the server.
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// NOTE: The actual registration of the service with the gRPC server depends on the protocol buffer
// definitions and the compiled code generated from those definitions.
