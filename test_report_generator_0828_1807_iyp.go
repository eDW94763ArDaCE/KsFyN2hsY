// 代码生成时间: 2025-08-28 18:07:56
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "os"
    "path/filepath"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/timestamppb"
    "google.golang.org/protobuf/types/known/anypb"
)

// TestReport defines the structure of a test report
type TestReport struct {
    Timestamp   *timestamppb.Timestamp
    TestResults []*TestResult
}

// TestResult defines the structure of an individual test result
type TestResult struct {
    TestName    string
    Status      string
    Description string
    Error       string
}

// ReportServiceServer is the server API for ReportService service
type ReportServiceServer struct {
    //
}

// NewReportServiceServer creates a new instance of ReportServiceServer
func NewReportServiceServer() *ReportServiceServer {
    return &ReportServiceServer{}
}

// GenerateTestReport generates a test report based on the provided test results
func (s *ReportServiceServer) GenerateTestReport(ctx context.Context, req *GenerateTestReportRequest) (*GenerateTestReportResponse, error) {
    // Validate the request
    if req == nil || req.TestResults == nil {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request")
    }

    // Create a new test report
    report := &TestReport{
        Timestamp: timestamppb.Now(),
        TestResults: make([]*TestResult, 0),
    }

    // Add test results to the report
    for _, result := range req.TestResults {
        report.TestResults = append(report.TestResults, &TestResult{
            TestName:    result.TestName,
            Status:      result.Status,
            Description: result.Description,
            Error:       result.Error,
        })
    }

    // Save the report to a file
    filename := fmt.Sprintf("test_report_%s.json", time.Now().Format("20060102150405"))
    data, err := json.MarshalIndent(report, "", "  ")
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to marshal report: %v", err)
    }
    if err := ioutil.WriteFile(filename, data, 0644); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to write report: %v", err)
    }

    // Return the response
    return &GenerateTestReportResponse{
        ReportFilename: filename,
    }, nil
}

// main is the entry point of the program
func main() {
    // Define the gRPC server address
    listen, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    defer listen.Close()

    // Create a new gRPC server
    srv := grpc.NewServer()

    // Register the ReportServiceServer
    RegisterReportServiceServer(srv, NewReportServiceServer())

    // Start the gRPC server
    log.Printf("server listening at %v", listen.Addr())
    if err := srv.Serve(listen); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// GenerateTestReportRequest is the request message for GenerateTestReport method
type GenerateTestReportRequest struct {
    TestResults []*TestResult `protobuf:"bytes,1,rep,name=testResults,proto3"`
}

// GenerateTestReportResponse is the response message for GenerateTestReport method
type GenerateTestReportResponse struct {
    ReportFilename string `protobuf:"bytes,1,opt,name=report_filename,proto3"`
}

// ReportService is the gRPC service definition
type ReportServiceServer interface {
    GenerateTestReport(context.Context, *GenerateTestReportRequest) (*GenerateTestReportResponse, error)
}

func RegisterReportServiceServer(s *grpc.Server, srv ReportServiceServer) {
    s.RegisterService(&_ReportService_serviceDesc, srv)
}

var _ReportService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "ReportService",
    HandlerType: (*ReportServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "GenerateTestReport",
            Handler: _ReportService_GenerateTestReport_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
   Metadata: "test_report_generator.proto",
}
