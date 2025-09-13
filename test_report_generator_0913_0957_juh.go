// 代码生成时间: 2025-09-13 09:57:33
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/timestamppb"
)

// TestReport is the message type for test reports
type TestReport struct {
    TestId      string    `protobuf:"bytes,1,opt,name=test_id,json=taskId"`
    TestName    string    `protobuf:"bytes,2,opt,name=test_name,json=testName"`
    StartTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime"`
    Duration    float64   `protobuf:"fixed64,4,opt,name=duration"`
    Status      string    `protobuf:"bytes,5,opt,name=status"`
    ErrorMessage string `protobuf:"bytes,6,opt,name=error_message,json=errorMessage"`
}

// TestReportServiceServer defines the gRPC service
type TestReportServiceServer struct {
    // embedding unexported fields
}

// GenerateTestReport generates a test report
func (s *TestReportServiceServer) GenerateTestReport(ctx context.Context, req *TestReport) (*TestReport, error) {
    // Check if the request is valid
    if req == nil || req.TestId == "" || req.TestName == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request")
    }

    // Simulate report generation logic
    // In a real-world scenario, this would involve more complex logic,
    // potentially involving database operations, file I/O, etc.

    // Set the start time to the current time
    req.StartTime = timestamppb.Now()

    // Set a mock duration and status (for demonstration purposes)
    req.Duration = 120.5 // 2 minutes
    req.Status = "PASSED"

    // Return the generated report
    return req, nil
}

// RegisterTestReportServiceServer registers the service with the gRPC server
func RegisterTestReportServiceServer(s *grpc.Server, server TestReportServiceServer) {
    // Register the service with the gRPC server
    pb.RegisterTestReportServiceServer(s, server)
}

// main is the entry point of the program
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")

    s := grpc.NewServer()
    RegisterTestReportServiceServer(s, TestReportServiceServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
