// 代码生成时间: 2025-09-15 21:26:35
package main

import (
    "fmt"
    "time"
)

// CacheItem represents a cached item with its value and expiration time
type CacheItem struct {
    Value       string    `json:"value"`
    Expiration time.Time `json:"expiration"`
}

// CacheService represents a simple cache service with a map of cached items
type CacheService struct {
    cache map[string]*CacheItem
}

// NewCacheService creates a new instance of CacheService
func NewCacheService() *CacheService {
    return &CacheService{
        cache: make(map[string]*CacheItem),
    }
}

// Set caches a new item with a given key and value, and optional expiration duration
func (cs *CacheService) Set(key string, value string, duration time.Duration) error {
    if duration <= 0 {
        return fmt.Errorf("invalid duration: %v", duration)
    }

    expiration := time.Now().Add(duration)
    cs.cache[key] = &CacheItem{
        Value:       value,
        Expiration: expiration,
    }
    return nil
}

// Get retrieves a cached item by key, returning an error if the item is expired or not found
func (cs *CacheService) Get(key string) (string, error) {
    item, exists := cs.cache[key]
    if !exists {
        return "", fmt.Errorf("item not found")
    }

    if time.Now().After(item.Expiration) {
        delete(cs.cache, key) // Remove expired item from cache
        return "", fmt.Errorf("item expired")
    }

    return item.Value, nil
}

// Clear removes all items from the cache
func (cs *CacheService) Clear() {
    cs.cache = make(map[string]*CacheItem)
}

// Example usage
func main() {
    cacheService := NewCacheService()
    err := cacheService.Set("key1", "value1", 5*time.Minute)
    if err != nil {
        fmt.Println("Error setting cache item: ", err)
        return
    }

    value, err := cacheService.Get("key1")
    if err != nil {
        fmt.Println("Error getting cache item: ", err)
        return
    }

    fmt.Println("Cached value: ", value)
    cacheService.Clear()
}
