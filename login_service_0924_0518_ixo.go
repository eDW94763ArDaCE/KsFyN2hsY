// 代码生成时间: 2025-09-24 05:18:13
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// User defines the structure of a user.
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// UserServer is the server API for User service.
type UserServer struct {
    // Add your implementation here.
    // Storage for the users
    users map[string]string
}

// NewUserServer creates a new instance of the UserServer.
func NewUserServer() *UserServer {
    return &UserServer{
        users: make(map[string]string),
    }
}

// CheckUser checks if a user exists and their password is correct.
func (s *UserServer) CheckUser(ctx context.Context, req *User) (*emptypb.Empty, error) {
    // Check if the user exists in the map and if the password is correct.
    if password, ok := s.users[req.Username]; ok && password == req.Password {
        return &emptypb.Empty{}, nil
    }
    return nil, status.Errorf(codes.Unauthenticated, "invalid username or password")
}

// RegisterUser registers a new user with the provided username and password.
func (s *UserServer) RegisterUser(ctx context.Context, req *User) (*emptypb.Empty, error) {
    // Check if the user already exists.
    if _, ok := s.users[req.Username]; ok {
        return nil, status.Errorf(codes.AlreadyExists, "user already exists")
    }
    // Store the user in the map.
    s.users[req.Username] = req.Password
    return &emptypb.Empty{}, nil
}

func main() {
   lis, err := net.Listen("tcp", "localhost:50051")
   if err != nil {
       log.Fatalf("failed to listen: %v", err)
   }

   creds, err := credentials.NewServerTLSFromFile("server.pem", "server.key")
   if err != nil {
       log.Fatalf("failed to generate credentials: %v", err)
   }

   grpcServer := grpc.NewServer(grpc.Creds(creds))
   fmt.Println("Server started on port 50051")

   // Register the reflection service on gRPC server.
   reflection.Register(grpcServer)

   // Create a new instance of the UserServer.
   userServer := NewUserServer()

   // Register the UserServer.
   // Assuming you have the user.proto compiled into a Go file named user.pb.go.
   userpb.RegisterUserServer(grpcServer, userServer)

   // Serve the gRPC server.
   if err := grpcServer.Serve(lis); err != nil {
       log.Fatalf("failed to serve: %v", err)
   }
}
