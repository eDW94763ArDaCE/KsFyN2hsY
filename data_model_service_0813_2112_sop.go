// 代码生成时间: 2025-08-13 21:12:58
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

// DataModel is the structure that will be used for our gRPC service.
type DataModel struct {
    Id        string                `json:"id"`
    Name      string                `json:"name"`
    CreatedAt *timestamppb.Timestamp `json:"createdAt"`
    UpdatedAt *timestamppb.Timestamp `json:"updatedAt"`
}

// Server is the receiver type for our service methods.
type Server struct {
    // This could be a database connection or any other state.
}

// GetDataModel is a method that will be called by the client to get data model details.
func (s *Server) GetDataModel(ctx context.Context, req *GetDataModelRequest) (*DataModel, error) {
    // Here you would typically interact with a database or other data source.
    // For simplicity, we'll just create a new DataModel and return it.
    dm := &DataModel{
        Id:        req.Id,
        Name:      "Sample Name",
        CreatedAt: timestamppb.Now(),
        UpdatedAt: timestamppb.Now(),
    }
    return dm, nil
}

// GetDataModelRequest is the request message for the GetDataModel RPC.
type GetDataModelRequest struct {
    Id string `json:"id"`
}

// RegisterServer registers the server with the gRPC server.
func RegisterServer(s *grpc.Server, server *Server) {
    // Register the server with the gRPC server.
    RegisterDataModelServiceServer(s, server)
}

// main is the entry point for the application.
func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    // Create a new gRPC server.
    s := grpc.NewServer()

    // Create a new Server instance.
    server := &Server{}

    // Register the server with the gRPC server.
    RegisterServer(s, server)

    // Start serving.
    if err := s.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// DataModelServiceServer is the server API for DataModelService service.
type DataModelServiceServer interface {
    GetDataModel(context.Context, *GetDataModelRequest) (*DataModel, error)
}

// RegisterDataModelServiceServer registers the DataModelServiceServer service with the gRPC server.
func RegisterDataModelServiceServer(s *grpc.Server, srv DataModelServiceServer) {
    s.RegisterService(&_DataModelService_serviceDesc, srv)
}

// _DataModelService_serviceDesc is the service descriptor for DataModelService.
var _DataModelService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "DataModelService",
    HandlerType: (*DataModelServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "GetDataModel",
            Handler: _DataModelService_GetDataModel_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "data_model.proto",
}

// _DataModelService_GetDataModel_Handler is the handler for the GetDataModel method.
func _DataModelService_GetDataModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(GetDataModelRequest)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(DataModelServiceServer).GetDataModel(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server: srv,
        FullMethod: "/DataModelService/GetDataModel",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(DataModelServiceServer).GetDataModel(ctx, req.(*GetDataModelRequest))
    }
    return interceptor(ctx, in, info, handler)
}

// NOTE: The actual protocol buffer definitions for DataModel, GetDataModelRequest,
// and the service definition should be in a .proto file, which is not included here.
// The .proto file should be compiled using the protoc compiler to generate the necessary
// Go code for the service and data models.
