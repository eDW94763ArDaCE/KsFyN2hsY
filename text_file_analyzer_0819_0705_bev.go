// 代码生成时间: 2025-08-19 07:05:02
// text_file_analyzer.go
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/types/known/emptypb"
    "path/filepath"
    "strings"
)

// Define the structure for the analyzer service.
type TextFileAnalyzerService struct{}

// Define the request message for the AnalyzeFile method.
type AnalyzeFileRequest struct {
    FilePath string
}

// Define the response message for the AnalyzeFile method.
type AnalyzeFileResponse struct {
    Content string
    Lines   int
    Words   int
    Chars   int
}

// AnalyzeFile analyzes the content of a text file and returns statistics.
func (s *TextFileAnalyzerService) AnalyzeFile(ctx context.Context, req *AnalyzeFileRequest) (*AnalyzeFileResponse, error) {
    fileContent, err := ioutil.ReadFile(req.FilePath)
    if err != nil {
        return nil, err
    }
    
    lines := strings.Count(fileContent, "
")
    words := strings.Count(string(fileContent), " ") + 1 // +1 for the last word
    chars := len(fileContent)
    
    return &AnalyzeFileResponse{
        Content: string(fileContent),
        Lines:   lines,
        Words:   words,
        Chars:   chars,
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    grpcServer := grpc.NewServer()
    
    // Register the service with the gRPC server.
    RegisterTextFileAnalyzerServiceServer(grpcServer, &TextFileAnalyzerService{})
    
    // Register reflection service on gRPC server.
    reflection.Register(grpcServer)
    
    fmt.Println("Server listening on port 50051")
    
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
