// 代码生成时间: 2025-08-10 15:35:49
package main

import (
    "context"
    "log"
    "net"
    "os"
    "path/filepath"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "archive/zip"
    "io"
    "io/ioutil"
    "archive/tar"
    "compress/gzip"
)

// FileCompressionService defines the methods for the file compression and decompression service.
type FileCompressionService struct{}

// DecompressZip decompresses a zip file to the specified directory.
func (s *FileCompressionService) DecompressZip(ctx context.Context, in *DecompressZipRequest) (*DecompressZipResponse, error) {
    // Check if the input is valid.
    if in == nil || in.ZipFilePath == "" || in.DestinationDirectory == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid input")
    }

    // Decompress the zip file.
    err := decompressZipFile(in.ZipFilePath, in.DestinationDirectory)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to decompress zip: %v", err)
    }

    // Return success response.
    return &DecompressZipResponse{Success: true}, nil
}

// DecompressTarGz decompresses a tar.gz file to the specified directory.
func (s *FileCompressionService) DecompressTarGz(ctx context.Context, in *DecompressTarGzRequest) (*DecompressTarGzResponse, error) {
    // Check if the input is valid.
    if in == nil || in.TarGzFilePath == "" || in.DestinationDirectory == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid input")
    }

    // Decompress the tar.gz file.
    err := decompressTarGzFile(in.TarGzFilePath, in.DestinationDirectory)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to decompress tar.gz: %v", err)
    }

    // Return success response.
    return &DecompressTarGzResponse{Success: true}, nil
}

// decompressZipFile decompresses a zip file.
func decompressZipFile(zipFilePath string, destinationDirectory string) error {
    reader, err := zip.OpenReader(zipFilePath)
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        rc, err := file.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        // Create directory structure.
        dirPath := filepath.Join(destinationDirectory, filepath.Dir(file.Name))
        if _, err := os.Stat(dirPath); os.IsNotExist(err) {
            os.MkdirAll(dirPath, 0755)
        }

        // Write file content to destination.
        filePath := filepath.Join(destinationDirectory, file.Name)
        fileWriter, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            return err
        }
        defer fileWriter.Close()
        _, err = io.Copy(fileWriter, rc)
        if err != nil {
            return err
        }
    }
    return nil
}

// decompressTarGzFile decompresses a tar.gz file.
func decompressTarGzFile(tarGzFilePath string, destinationDirectory string) error {
    file, err := os.Open(tarGzFilePath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    gzReader := gzip.NewReader(file)
    defer gzReader.Close()

    tarReader := tar.NewReader(gzReader)
    for {
        header, err := tarReader.Next()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        targetPath := filepath.Join(destinationDirectory, header.Name)
        switch header.Typeflag {
        case tar.TypeDir:
            if _, err := os.Stat(targetPath); err != nil {
                if err := os.MkdirAll(targetPath, 0755); err != nil {
                    return err
                }
            }
        case tar.TypeReg:
            outFile, err := os.Create(targetPath)
            if err != nil {
                return err
            }
            defer outFile.Close()
            if _, err := io.Copy(outFile, tarReader); err != nil {
                return err
            }
        }
    }
    return nil
}

// StartGRPCServer starts the gRPC server.
func StartGRPCServer() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    defer listener.Close()

    // Register the file compression service on the server.
    s := grpc.NewServer()
    pb.RegisterFileCompressionServiceServer(s, &FileCompressionService{})

    // Start the server in a separate goroutine.
    go func() {
        if err := s.Serve(listener); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()
}

// main function to start the gRPC server.
func main() {
    StartGRPCServer()
}

// DecompressZipRequest defines the request message for the DecompressZip method.
type DecompressZipRequest struct {
    ZipFilePath          string
    DestinationDirectory string
}

// DecompressZipResponse defines the response message for the DecompressZip method.
type DecompressZipResponse struct {
    Success bool
}

// DecompressTarGzRequest defines the request message for the DecompressTarGz method.
type DecompressTarGzRequest struct {
    TarGzFilePath        string
    DestinationDirectory string
}

// DecompressTarGzResponse defines the response message for the DecompressTarGz method.
type DecompressTarGzResponse struct {
    Success bool
}

// FileCompressionServiceServer defines the server-side interface for the file compression service.
type FileCompressionServiceServer interface {
    DecompressZip(context.Context, *DecompressZipRequest) (*DecompressZipResponse, error)
    DecompressTarGz(context.Context, *DecompressTarGzRequest) (*DecompressTarGzResponse, error)
}

// RegisterFileCompressionServiceServer registers the file compression service on the gRPC server.
func RegisterFileCompressionServiceServer(s *grpc.Server, srv FileCompressionServiceServer) {
    pb.RegisterFileCompressionServiceServer(s, srv)
}
