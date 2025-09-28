// 代码生成时间: 2025-09-29 00:01:43
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"compress/gzip"
	"github.com/golang/protobuf/proto"
)

const (
	port = ":50051"
)

// FileCompressionService defines the service methods for file compression and decompression.
type FileCompressionService struct{}

// Compress a file by gzip algorithm.
func (s *FileCompressionService) Compress(c context.Context, in *CompressRequest) (*CompressResponse, error) {
	if in.File == nil || len(in.File.Data) == 0 {
		return nil, errors.New("file data is empty")
	}
	
	compressedFile, err := gzipCompress(in.File.Data)
	if err != nil {
		return nil, err
	}
	
	return &CompressResponse{File: &File{Data: compressedFile}}, nil
}

// Decompress a file that was compressed by gzip algorithm.
func (s *FileCompressionService) Decompress(c context.Context, in *DecompressRequest) (*DecompressResponse, error) {
	if in.File == nil || len(in.File.Data) == 0 {
		return nil, errors.New("file data is empty")
	}
	
	decompressedFile, err := gzipDecompress(in.File.Data)
	if err != nil {
		return nil, err
	}
	
	return &DecompressResponse{File: &File{Data: decompressedFile}}, nil
}

// gzipCompress compresses data to gzip format.
func gzipCompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// gzipDecompress decompresses gzip formatted data to original format.
func gzipDecompress(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	decompressedData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	r.Close()
	return decompressedData, nil
}

// startService starts the GRPC server with reflection.
func startService() {
	ls, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	registerFileCompressionService(grpcServer)
	
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// registerFileCompressionService registers the file compression service to the GRPC server.
func registerFileCompressionService(s *grpc.Server) {
	RegisterFileCompressionServiceServer(s, &FileCompressionService{})
}

func main() {
	startService()
}

// The following are the GRPC service definitions and message types in proto files.
// You would need to define these in a .proto file and generate the corresponding Go code.

// message File {
//     bytes data = 1;
// }

// message CompressRequest {
//     File file = 1;
// }

// message CompressResponse {
//     File file = 1;
// }

// message DecompressRequest {
//     File file = 1;
// }

// message DecompressResponse {
//     File file = 1;
// }

// service FileCompressionService {
//     rpc Compress(CompressRequest) returns (CompressResponse);
//     rpc Decompress(DecompressRequest) returns (DecompressResponse);
// }
