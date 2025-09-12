// 代码生成时间: 2025-09-12 08:30:50
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// InventoryService defines the gRPC service for inventory management.
type InventoryService struct{}

// CheckInventory checks the stock of an item and returns the available quantity.
func (i *InventoryService) CheckInventory(ctx context.Context, req *InventoryRequest) (*InventoryResponse, error) {
    // Here you would add your logic to check the inventory, for now, we'll just simulate some data.
    stock := make(map[string]int)
    stock["item1"] = 10
    stock[req.ItemId] = stock[req.ItemId] - req.Quantity
    if stock[req.ItemId] < 0 {
        return nil, fmt.Errorf("not enough stock for item: %s", req.ItemId)
    }
    return &InventoryResponse{Quantity: stock[req.ItemId]}, nil
}

// UpdateInventory updates the stock of an item based on the provided request.
func (i *InventoryService) UpdateInventory(ctx context.Context, req *InventoryUpdateRequest) (*emptypb.Empty, error) {
    // Here you would add your logic to update the inventory, for now, we'll just simulate some data.
    stock := make(map[string]int)
    stock["item1"] = 10
    stock[req.ItemId] = stock[req.ItemId] + req.Delta
    return &emptypb.Empty{}, nil
}

// InventoryRequest is the request message for checking inventory.
type InventoryRequest struct {
    ItemId  string
    Quantity int32
}

// InventoryResponse is the response message for inventory check.
type InventoryResponse struct {
    Quantity int32
}

// InventoryUpdateRequest is the request message for updating inventory.
type InventoryUpdateRequest struct {
    ItemId string
    Delta  int32
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")

    grpcServer := grpc.NewServer()
    RegisterInventoryServiceServer(grpcServer, &InventoryService{})
    reflection.Register(grpcServer)
    grpcServer.Serve(lis)
}

// RegisterInventoryServiceServer registers the InventoryService with the gRPC server.
func RegisterInventoryServiceServer(s *grpc.Server, srv *InventoryService) {
    s.RegisterService(&_InventoryService_serviceDesc, srv)
}

// This is just a placeholder service descriptor for generating the gRPC service.
// In a real application, you would generate this from a .proto file.
var _InventoryService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "InventoryService",
    HandlerType: (*InventoryService)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "CheckInventory",
            Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
                in := new(InventoryRequest)
                if err := dec(in); err != nil {
                    return nil, err
                }
                return srv.(*InventoryService).CheckInventory(ctx, in)
            },
        },
        {
            MethodName: "UpdateInventory",
            Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
                in := new(InventoryUpdateRequest)
                if err := dec(in); err != nil {
                    return nil, err
                }
                return srv.(*InventoryService).UpdateInventory(ctx, in)
            },
        },
    },
    Streams: []grpc.StreamDesc{},
   Metadata: "inventory_management.pb",
}
