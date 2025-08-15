// 代码生成时间: 2025-08-16 05:24:56
// image_resize_service.go
// This service provides functionality for batch resizing of images.

package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "os"
    "strings"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

import "path"

// ImageResizeService defines the gRPC service for resizing images
type ImageResizeService struct{}

// ResizeImage resizes an image to the specified dimensions
func (s *ImageResizeService) ResizeImage(ctx context.Context, req *ImageResizeRequest) (*emptypb.Empty, error) {
    // Check if the image path and resize dimensions are valid
    if req.GetImagePath() == "" || req.GetWidth() <= 0 || req.GetHeight() <= 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid image path or dimensions")
    }

    // Read the image from the file path
    imgData, err := ioutil.ReadFile(req.GetImagePath())
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to read image: %v", err)
    }

    // TODO: Add image processing logic here (e.g., using an image processing library)
    // For demonstration purposes, we simulate a successful resize
    fmt.Printf("Resizing image from %s to dimensions: %dx%d
", req.GetImagePath(), req.GetWidth(), req.GetHeight())

    // Write the resized image to a new file path
    // TODO: Implement writing logic here

    return &emptypb.Empty{}, nil
}

// ImageResizeRequest defines the request message for resizing an image
type ImageResizeRequest struct {
    ImagePath string `protobuf:"bytes,1,opt,name=image_path,json=indexPath,proto3" json:"imagePath,omitempty"`
    Width     int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
    Height    int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")

    s := grpc.NewServer()
    pb.RegisterImageResizeServiceServer(s, &ImageResizeService{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
