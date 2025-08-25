// 代码生成时间: 2025-08-25 17:19:55
package main

import (
    "fmt"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/protobuf/types/known/timestamppb"
    "path/filepath"
    "os"
    "strings"

    "github.com/xuri/excelize/v2"
)

// ExcelGeneratorService defines the gRPC service interface.
type ExcelGeneratorService struct{}

// GenerateExcelSpreadsheet generates an Excel spreadsheet based on the given data.
func (s *ExcelGeneratorService) GenerateExcelSpreadsheet(ctx context.Context, req *GenerateExcelRequest) (*GenerateExcelResponse, error) {
    f := excelize.File{}
    sheetName := "Sheet1"
    if err := f.NewSheet(sheetName); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to create sheet: %v", err)
    }

    // Write data to the sheet.
    for i, data := range req.Data {
        for j, cell := range data {
            if err := f.SetCellValue(sheetName, fmt.Sprintf("A%d", j+1), cell.Value); err != nil {
                return nil, status.Errorf(codes.Internal, "failed to set cell value: %v", err)
            }
        }
    }

    // Save the spreadsheet to a file.
    filePath := filepath.Join(req.OutputPath, req.FileName)
    if err := f.SaveAs(filePath); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to save file: %v", err)
    }

    return &GenerateExcelResponse{FilePath: filePath}, nil
}

// Server starts the gRPC server.
func Serve() error {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        return status.Errorf(codes.Internal, "failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")

    grpcServer := grpc.NewServer()
    RegisterExcelGeneratorServiceServer(grpcServer, &ExcelGeneratorService{})
    return grpcServer.Serve(lis)
}

func main() {
    if err := Serve(); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Define the gRPC message for the request.
type GenerateExcelRequest struct {
    Data       [][]*CellData `protobuf:"bytes,1,rep,name=data,proto3"`
    OutputPath string        `protobuf:"bytes,2,opt,name=output_path,json=outputPath,proto3"`
    FileName   string        `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3"`
}

// Define the gRPC message for the response.
type GenerateExcelResponse struct {
    FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3"`
}

// Define the structure for cell data.
type CellData struct {
    Value string `protobuf:"bytes,1,opt,name=value,proto3"`
}

// RegisterExcelGeneratorServiceServer registers the ExcelGeneratorService server to the gRPC server.
func RegisterExcelGeneratorServiceServer(s *grpc.Server, srv ExcelGeneratorService) {
    ExcelGeneratorServiceServer = &excelGeneratorServiceServer{srv}
    s.RegisterService(&_ExcelGeneratorService_serviceDesc, ExcelGeneratorServiceServer)
}

// excelGeneratorServiceServer is the server API for ExcelGeneratorService service.
type excelGeneratorServiceServer struct {
    ExcelGeneratorService
}

// UnimplementedExcelGeneratorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExcelGeneratorServiceServer struct{}

// Unimplemented ExcelGeneratorServiceServer must be embedded to have forward compatible implementations.
func (*UnimplementedExcelGeneratorServiceServer) GenerateExcelSpreadsheet(context.Context, *GenerateExcelRequest) (*GenerateExcelResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method GenerateExcelSpreadsheet not implemented")
}

// excelGeneratorServiceServer implements ExcelGeneratorServiceServer interface.
func (s *excelGeneratorServiceServer) GenerateExcelSpreadsheet(ctx context.Context, req *GenerateExcelRequest) (*GenerateExcelResponse, error) {
    return s.ExcelGeneratorService.GenerateExcelSpreadsheet(ctx, req)
}

// Define the service descriptor.
var _ExcelGeneratorService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "ExcelGeneratorService",
    HandlerType: (*ExcelGeneratorService)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "GenerateExcelSpreadsheet",
            Handler: func (srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
                return srv.(ExcelGeneratorServiceServer).GenerateExcelSpreadsheet(ctx, dec(&GenerateExcelRequest{}))
            },
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "excel_generator_service.proto",
}
