// 代码生成时间: 2025-09-21 18:49:27
package main

import (
    "fmt"
    "net/url"
    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
)

// URLValidationService is the server API for URLValidation service.
type URLValidationService struct{}

// ValidateURL checks if the provided URL is valid.
func (s *URLValidationService) ValidateURL(ctx context.Context, req *URLValidationRequest) (*URLValidationResponse, error) {
    // Parse the URL from the request
    u, err := url.Parse(req.GetUrl())
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "invalid URL: %v", err)
    }
    if u.Scheme == "" || u.Host == "" {
        return nil, status.Errorf(codes.InvalidArgument, "URL must have a scheme and a host")
    }
    // If all checks pass, the URL is valid
    return &URLValidationResponse{Valid: true}, nil
}

// URLValidationRequest is the request message for the ValidateURL RPC.
type URLValidationRequest struct {
    Url string
}

// URLValidationResponse is the response message for the ValidateURL RPC.
type URLValidationResponse struct {
    Valid bool
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on :50051")

    s := grpc.NewServer()
    pb.RegisterURLValidationServiceServer(s, &URLValidationService{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// The following lines are for the protobuf definitions,
// which should be placed in a .proto file and compiled using protoc.

// syntax = "proto3";
//
// package urlvalidation;
//
// service URLValidation {
//     rpc ValidateURL(URLValidationRequest) returns (URLValidationResponse);
// }
//
// message URLValidationRequest {
//     string url = 1;
// }
//
// message URLValidationResponse {
//     bool valid = 1;
// }