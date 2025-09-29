// 代码生成时间: 2025-09-29 16:09:10
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "github.com/microcosm-cc/bluemonday"
)

// XSSService is a service that provides XSS protection.
type XSSService struct{}

// IsSafe checks if the given text is safe from XSS attacks.
func (s *XSSService) IsSafe(ctx context.Context, in *IsSafeRequest) (*IsSafeResponse, error) {
    // Use bluemonday to sanitize the input text.
    safeText := bluemonday.UGCPolicy().Sanitize(in.GetText())
    if safeText == in.GetText() {
        // If the sanitized text matches the original text, it is safe.
        return &IsSafeResponse{IsSafe: true}, nil
    } else {
        // If the sanitized text does not match, there were potential XSS threats.
        return &IsSafeResponse{IsSafe: false}, nil
    }
}

// server is used to implement the XSSServiceServer interface.
type server struct{
    XSSServiceServer
}
os
// IsSafeRequest is the request for checking XSS safety.
type IsSafeRequest struct {
    Text string `protobuf:"bytes,1,opt,name=text"` // The text to be checked for XSS.
}
os
// IsSafeResponse is the response for checking XSS safety.
type IsSafeResponse struct {
    IsSafe bool `protobuf:"varint,1,opt,name=is_safe,json=isSafe"` // Whether the text is safe from XSS.
}
os
// Run starts the GRPC server with the XSSService.
func Run() error {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
        return err
    }

    fmt.Println("Listening on port 50051")

    grpcServer := grpc.NewServer()
    RegisterXSSServiceServer(grpcServer, &server{})

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
        return err
    }
    return nil
}

func main() {
    if err := Run(); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}

// The following would be your protobuf definitions and registration code, which are
// omitted for brevity.
// You would also need to generate the Go code from your .proto files using the
// protoc compiler with the appropriate gRPC plugins.
