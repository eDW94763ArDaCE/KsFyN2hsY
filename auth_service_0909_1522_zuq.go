// 代码生成时间: 2025-09-09 15:22:14
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
)
# NOTE: 重要实现细节

// User represents a user entity with ID and Username
type User struct {
    ID       string
    Username string
}

// AuthServiceServer is the server API for AuthService service
type AuthServiceServer struct {
    // UserMap to store user credentials for demonstration purposes
    UserMap map[string]string
}
# TODO: 优化性能

// NewAuthServiceServer creates a new instance of AuthServiceServer
func NewAuthServiceServer() *AuthServiceServer {
    return &AuthServiceServer{
# NOTE: 重要实现细节
        UserMap: make(map[string]string),
# 增强安全性
    }
}
# 增强安全性

// Authenticate checks the provided credentials and returns the authenticated user ID
func (s *AuthServiceServer) Authenticate(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
    // Check if the user exists and the password is correct
    if password, exists := s.UserMap[req.Username]; exists && password == req.Password {
        // Return the user ID if authentication is successful
        return &AuthResponse{
            UserId: fmt.Sprintf("user-%s", req.Username),
        }, nil
    }
    // Return an error if authentication fails
    return nil, status.Errorf(codes.Unauthenticated, "invalid username or password")
}

// AuthRequest is the request message for the Authenticate RPC
type AuthRequest struct {
    Username string
    Password string
}

// AuthResponse is the response message for the Authenticate RPC
type AuthResponse struct {
    UserId string
}

// RegisterUser registers a new user with the given username and password
func (s *AuthServiceServer) RegisterUser(ctx context.Context, req *User) (*User, error) {
    // Check if the user already exists
    if _, exists := s.UserMap[req.Username]; exists {
        return nil, status.Errorf(codes.AlreadyExists, "user already exists")
    }
    // Store the new user in the UserMap
    s.UserMap[req.Username] = req.ID
    return req, nil
}

func main() {
    server := NewAuthServiceServer()
    server.UserMap["testuser"] = "testpass" // Pre-register a user for demonstration

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on :50051")

    s := grpc.NewServer()
    RegisterAuthServiceServer(s, server)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterAuthServiceServer is a utility function to register the AuthServiceServer to a gRPC server
func RegisterAuthServiceServer(s *grpc.Server, srv *AuthServiceServer) {
    RegisterAuthServiceServer(s, srv)
}
