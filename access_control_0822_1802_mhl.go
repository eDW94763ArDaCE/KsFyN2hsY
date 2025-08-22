// 代码生成时间: 2025-08-22 18:02:52
package main

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Define the service with an RPC method that requires authorization.
type MyService struct{}

// SayHello implements the RPC method.
func (s *MyService) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	// Check for authorization in the context's metadata.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "There is no metadata in the context")
	}

	// Assume we expect a specific token in the metadata.
	token, exists := md[\