// 代码生成时间: 2025-09-24 01:18:46
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// InventoryItem represents an item in the inventory
type InventoryItem struct {
    ID          string
    Name        string
    Quantity    int
}

// InventoryService provides methods to manage the inventory
type InventoryService struct {
    // Map to store inventory items
    items map[string]InventoryItem
}

// NewInventoryService creates a new instance of InventoryService
func NewInventoryService() *InventoryService {
    return &InventoryService{
        items: make(map[string]InventoryItem),
    }
}

// AddItem adds a new item to the inventory
func (s *InventoryService) AddItem(ctx context.Context, item *InventoryItem) error {
    if _, exists := s.items[item.ID]; exists {
        return status.Errorf(codes.AlreadyExists, "item with ID %s already exists", item.ID)
    }
    s.items[item.ID] = *item
    return nil
}

// RemoveItem removes an item from the inventory
func (s *InventoryService) RemoveItem(ctx context.Context, id string) error {
    if _, exists := s.items[id]; !exists {
        return status.Errorf(codes.NotFound, "item with ID %s not found", id)
    }
    delete(s.items, id)
    return nil
}

// UpdateItem updates an existing item in the inventory
func (s *InventoryService) UpdateItem(ctx context.Context, item *InventoryItem) error {
    if _, exists := s.items[item.ID]; !exists {
        return status.Errorf(codes.NotFound, "item with ID %s not found", item.ID)
    }
    s.items[item.ID] = *item
    return nil
}

// GetItem retrieves an item from the inventory by ID
func (s *InventoryService) GetItem(ctx context.Context, id string) (*InventoryItem, error) {
    item, exists := s.items[id]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "item with ID %s not found", id)
    }
    return &item, nil
}

// ListItems lists all items in the inventory
func (s *InventoryService) ListItems(ctx context.Context) ([]InventoryItem, error) {
    var items []InventoryItem
    for _, item := range s.items {
        items = append(items, item)
    }
    return items, nil
}

// RunServer starts the gRPC server
func RunServer(address string, service *InventoryService) {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on " + address)
    
    s := grpc.NewServer()
    // Register the InventoryService with the gRPC server
    // Assuming a proto file with the service definition is available and compiled
    // pb.RegisterInventoryServiceServer(s, service)
    
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    service := NewInventoryService()
    RunServer(":50051", service)
}
