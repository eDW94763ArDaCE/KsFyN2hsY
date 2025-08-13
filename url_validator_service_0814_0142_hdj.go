// 代码生成时间: 2025-08-14 01:42:24
package main

import (
    "context"
    "fmt"
    "net/url"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// URLValidatorService is the service that implements the URL validation functionality.
type URLValidatorService struct{}

// ValidateURL checks if the provided URL is valid or not.
func (s *URLValidatorService) ValidateURL(ctx context.Context, in *ValidateURLRequest) (*ValidateURLResponse, error) {
    // Parse the URL to check its validity
    u, err := url.ParseRequestURI(in.Url)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "invalid URL: %v", err)
    }

    // Additional checks can be implemented here if necessary
    // For example, check if the URL scheme is allowed, or if the host is valid, etc.

    // Return a response with the result of the URL validation
    return &ValidateURLResponse{IsValid: u.Scheme != ""}, nil
}

// ValidateURLRequest is the request message for the ValidateURL method.
type ValidateURLRequest struct {
    Url string
}

// ValidateURLResponse is the response message for the ValidateURL method.
type ValidateURLResponse struct {
    IsValid bool
}

func main() {
    // Define the server and listen for incoming connections
    server := grpc.NewServer()
    RegisterURLValidatorServiceServer(server, &URLValidatorService{})
    fmt.Println("Server is running...")
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        panic(err)
    }
    if err := server.Serve(lis); err != nil {
        panic(err)
    }
}

// RegisterURLValidatorServiceServer registers the server to handle incoming requests.
func RegisterURLValidatorServiceServer(s *grpc.Server, srv *URLValidatorService) {
    pb.RegisterURLValidatorServiceServer(s, srv)
}

// Note: The above code assumes the existence of a protobuf file named 'url_validator_service.proto'
// which defines the URLValidatorService, ValidateURLRequest, and ValidateURLResponse.
// The generated code from this protobuf file should be imported as 'pb'.
