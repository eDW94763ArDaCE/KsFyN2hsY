// 代码生成时间: 2025-08-21 10:27:15
// form_validator.go

// Package formvalidator provides functionality to validate form data using gRPC.
package formvalidator

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FormValidatorService defines the gRPC service for form validation.
type FormValidatorService struct{}

// ValidateForm implements the gRPC method to validate form data.
func (s *FormValidatorService) ValidateForm(ctx context.Context, req *ValidateFormRequest) (*ValidateFormResponse, error) {
	// Check if the request is nil
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	// Validate email address
	if !isValidEmail(req.Email) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email: %s", req.Email)
	}

	// Validate username
	if !isValidUsername(req.Username) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid username: %s", req.Username)
	}

	// Return a success response
	return &ValidateFormResponse{
		IsValid: true,
	}, nil
}

// isValidEmail checks if the email address matches the expected pattern.
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// isValidUsername checks if the username only contains alphanumeric characters.
func isValidUsername(username string) bool {
	pattern := `^[a-zA-Z0-9]+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(username)
}

// ValidateFormRequest defines the request message for validating form data.
type ValidateFormRequest struct {
	Email    string
	Username string
}

// ValidateFormResponse defines the response message for validating form data.
type ValidateFormResponse struct {
	IsValid bool
}

// RegisterServer registers the form validator service with the gRPC server.
func RegisterServer(s *grpc.Server, service FormValidatorService) {
	RegisterFormValidatorServer(s, service)
}

// main function to start the gRPC server.
func main() {
	lis, err := net.Listen("net", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterServer(s, &FormValidatorService{})

	fmt.Println("Server is running on localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}