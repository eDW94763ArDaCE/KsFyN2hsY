// 代码生成时间: 2025-09-01 06:57:57
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "google.golang.org/grpc"
    "net"
)

// TestReport defines the structure of a test report.
type TestReport struct {
    StartTime     time.Time
    EndTime       time.Time
    TestResults  []string
    OverallStatus string
}

// TestReportService is the server API for TestReport service.
type TestReportService struct {
    // No fields needed for this example service.
}

// GenerateTestReport RPC implementation.
func (s *TestReportService) GenerateTestReport(ctx context.Context, req *GenerateTestReportRequest) (*TestReport, error) {
    // Create a new test report instance.
    report := &TestReport{
        StartTime: time.Now(),
        EndTime:   time.Now().Add(5 * time.Minute),
        TestResults: []string{},
        OverallStatus: "PASS",
    }

    // Simulate test cases execution and populate test results.
    for _, testCase := range req.TestCases {
        // Simulate test execution.
        if testCase == "failing_test" {
            report.TestResults = append(report.TestResults, "Test failed: "+testCase)
            report.OverallStatus = "FAIL"
        } else {
            report.TestResults = append(report.TestResults, "Test passed: "+testCase)
        }
    }

    return report, nil
}

// GenerateTestReportRequest defines the request structure for GenerateTestReport RPC.
type GenerateTestReportRequest struct {
    TestCases []string
}

// TestReportServer is the gRPC server for TestReport service.
func TestReportServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is listening on port 50051")
    defer lis.Close()

    s := grpc.NewServer()
    RegisterTestReportServiceServer(s, &TestReportService{})
    fmt.Println("TestReportService registered")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    TestReportServer()
}

// RegisterTestReportServiceServer registers the TestReportService with the gRPC server.
func RegisterTestReportServiceServer(s *grpc.Server, srv *TestReportService) {
    v1.RegisterTestReportServiceServer(s, srv)
}

// v1 is the proto package for TestReportService.
package v1

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

// TestReportService provides methods for generating and retrieving test reports.
type TestReportServiceServer interface {
    GenerateTestReport(context.Context, *GenerateTestReportRequest) (*TestReport, error)
}

func RegisterTestReportServiceServer(s *grpc.Server, srv TestReportServiceServer) {
    s.RegisterService(&_TestReportService_serviceDesc, srv)
}

func _TestReportService_GenerateTestReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(GenerateTestReportRequest)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil { return srv.(TestReportServiceServer).GenerateTestReport(ctx, in) }
    info := &grpc.UnaryServerInfo{
       Server: srv,
        FullMethod: "/v1.TestReportService/GenerateTestReport",
        }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(TestReportServiceServer).GenerateTestReport(ctx, req.(*GenerateTestReportRequest))
    }
    return interceptor(ctx, in, info, handler)
}

var _TestReportService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "v1.TestReportService",
    HandlerType: (*TestReportServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "GenerateTestReport",
            Handler: _TestReportService_GenerateTestReport_Handler,
        },
    },
    Streams: []grpc.StreamDesc{
    },
    Metadata: "v1/test_report_service.proto",
}

// TestReport defines the structure of a test report.
type TestReport struct {
    StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:""`
    EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:""`
    TestResults []string `protobuf:"bytes,3,rep,name=test_results,json=testResults" json:""`
    OverallStatus string `protobuf:"bytes,4,opt,name=overall_status,json=overallStatus" json:""`
}

// GenerateTestReportRequest defines the request structure for GenerateTestReport RPC.
type GenerateTestReportRequest struct {
    TestCases []string `protobuf:"bytes,1,rep,name=test_cases,json=testCases" json:""`
}
