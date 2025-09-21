// 代码生成时间: 2025-09-21 23:05:00
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/timestamppb"
    "github.com/xuri/excelize/v2"
)

// ExcelGeneratorService defines the gRPC service for generating Excel files.
type ExcelGeneratorService struct{}

// GenerateExcel is a gRPC method that generates an Excel file based on the given request.
func (s *ExcelGeneratorService) GenerateExcel(ctx context.Context, req *GenerateExcelRequest) (*GenerateExcelResponse, error) {
    // Validate the request
    if req == nil || req.Data == nil || len(req.Data) == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "empty request")
    }

    // Create a new Excel file
    f := excelize.NewFile()
    sheetName := req.SheetName
    if sheetName == "" {
        sheetName = "Sheet1"
    }

    // Add data to the Excel file
    for _, row := range req.Data {
        for _, cell := range row {
            f.SetCellValue(sheetName, fmt.Sprintf("A%d", len(row)+1), cell)
        }
    }

    // Save the file to the specified path
    filePath := req.FilePath
    if filePath == "" {
        filePath = "output.xlsx"
    }
    err := f.SaveAs(filePath)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to save Excel file: %v", err)
    }

    // Return a success response
    return &GenerateExcelResponse{
        FilePath: proto.String(filePath),
        Success: proto.Bool(true),
    }, nil
}

// GenerateExcelRequest defines the request message for the GenerateExcel method.
type GenerateExcelRequest struct {
    // Data is a 2D slice of strings representing the Excel data.
    Data [][]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
    // SheetName is the name of the sheet to be created in the Excel file.
    SheetName string `protobuf:"bytes,2,opt,name=sheet_name,json=sheetName,proto3" json:"sheetName,omitempty"`
    // FilePath is the file path where the Excel file will be saved.
    FilePath string `protobuf:"bytes,3,opt,name=file_path,json=file_path,proto3" json:"filePath,omitempty"`
}

// GenerateExcelResponse defines the response message for the GenerateExcel method.
type GenerateExcelResponse struct {
    // FilePath is the file path where the Excel file was saved.
    FilePath *string `protobuf:"bytes,1,opt,name=file_path,json=file_path,proto3" json:"filePath,omitempty"`
    // Success indicates whether the Excel file was generated successfully.
    Success *bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func main() {
   lis, err := net.Listen("tcp", ":50051")
   if err != nil {
       log.Fatalf("failed to listen: %v", err)
   }

   // Create a new gRPC server
   srv := grpc.NewServer()

   // Register the ExcelGeneratorService with the server
   pb.RegisterExcelGeneratorServiceServer(srv, &ExcelGeneratorService{})

   // Register reflection service on gRPC server.
   reflection.Register(srv)

   // Start the server
   if err := srv.Serve(lis); err != nil {
       log.Fatalf("failed to serve: %v", err)
   }
}
