// 代码生成时间: 2025-08-19 23:14:05
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
    "time"
)

// User represents a user entity with a username and password
type User struct {
    Username string
    Password string
}

// AuthServiceServer is the server API for AuthService service
type AuthServiceServer struct{}

// Authenticate method checks if the provided username and password are valid
func (s *AuthServiceServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
    // Simulate user database
    users := map[string]string{
        "admin": "password",
    }

    // Check if the username and password are valid
    storedPassword, ok := users[req.Username]
    if !ok || storedPassword != req.Password {
        return nil, status.Errorf(codes.Unauthenticated, "invalid username or password")
    }

    // Return a success response
    return &AuthenticateResponse{Success: true}, nil
}

// AuthenticateRequest is the request for Authenticate method
type AuthenticateRequest struct {
    Username string
    Password string
}

// AuthenticateResponse is the response for Authenticate method
type AuthenticateResponse struct {
    Success bool
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    s := grpc.NewServer()
    RegisterAuthServiceServer(s, &AuthServiceServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterAuthServiceServer registers the AuthServiceServer with the gRPC server
func RegisterAuthServiceServer(s *grpc.Server, srv *AuthServiceServer) {
    s.RegisterService(&_AuthService_serviceDesc, srv)
}

// AuthServiceServer must embed UnimplementedAuthServiceServer to forward requests to the storage implementation
type UnimplementedAuthServiceServer struct{}

// UnimplementedAuthServiceServer must be embedded to have forward compatible methods
func (*UnimplementedAuthServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

// The following are generated code by the protocol buffers compiler.
// Service descriptor
var _AuthService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "AuthService",
    HandlerType: (*AuthServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "Authenticate",
            Handler: grpc.HandlerFunc(func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
                return srv.(AuthServiceServer).Authenticate(ctx, dec(&AuthenticateRequest{}))
            }),
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "auth.proto",
}
