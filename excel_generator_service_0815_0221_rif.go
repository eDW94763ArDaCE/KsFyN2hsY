// 代码生成时间: 2025-08-15 02:21:50
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
    "log"
    "path/filepath"
    "excelize"
)

// ExcelGeneratorService is the server API for ExcelGenerator service.
type ExcelGeneratorService struct{}

// GenerateExcel is a RPC method to generate an Excel file based on the input parameters.
func (s *ExcelGeneratorService) GenerateExcel(ctx context.Context, req *GenerateExcelRequest) (*GenerateExcelResponse, error) {
    var response GenerateExcelResponse
    defer func() {
        if r := recover(); r != nil {
            response.Error = fmt.Sprintf("%v", r)
            return &response, fmt.Errorf("panic occurred")
        }
    }()

    // Open the file.
    f := excelize.NewFile()
    defer f.Close()

    // Add a sheet.
    index := f.NewSheet("Sheet1")

    // Set the active sheet of the workbook to index.
    f.SetActiveSheet(index)

    // Write data to the file.
    for i, dataRow := range req.Data {
        for j, cellData := range dataRow {
            if err := f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), cellData); err != nil {
                return nil, err
            }
        }
    }

    // Save the file.
    dir, _ := filepath.Abs(filepath.Dir(req.FileName))
    file, err := os.Create(filepath.Join(dir, req.FileName))
    if err != nil {
        return nil, err
    }
    defer file.Close()

    if err := f.Write(file); err != nil {
        return nil, err
    }

    response.FilePath = req.FileName
    return &response, nil
}

// GenerateExcelRequest is the request message for the GenerateExcel RPC.
type GenerateExcelRequest struct {
    // The name of the file to generate.
    FileName string "protobuf:\"s,opt,name=fileName,proto3\"
    // The data to populate in the Excel file.
    Data [][]string "protobuf:\"rep,name=data,proto3\"
}

// GenerateExcelResponse is the response message for the GenerateExcel RPC.
type GenerateExcelResponse struct {
    // The path to the generated file.
    FilePath string "protobuf:\"s,opt,name=filePath,proto3\"
    // Any error that occurred during file generation.
    Error string "protobuf:\"s,opt,name=error,proto3\"
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port :50051")

    s := grpc.NewServer()
    pb.RegisterExcelGeneratorServiceServer(s, &ExcelGeneratorService{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
