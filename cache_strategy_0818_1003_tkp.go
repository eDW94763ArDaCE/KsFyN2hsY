// 代码生成时间: 2025-08-18 10:03:26
package main
# FIXME: 处理边界情况

import (
    "fmt"
    "log"
    "sync"
    "time"
)

// CacheItem represents an item in the cache with its value and expiration time.
type CacheItem struct {
    Value    interface{}
    Expiry   time.Time
}

// Cache is the main cache structure that holds the cached items.
type Cache struct {
    items map[string]CacheItem
    mu    sync.RWMutex
}

// NewCache creates a new cache instance.
func NewCache() *Cache {
    return &Cache{items: make(map[string]CacheItem)}
}

// Set sets the value for a given key with an expiration time.
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    expiry := time.Now().Add(duration)
    c.items[key] = CacheItem{Value: value, Expiry: expiry}
}

// Get retrieves a value by its key and returns the item if it exists and is not expired.
func (c *Cache) Get(key string) (interface{}, bool) {
# 改进用户体验
    c.mu.RLock()
    defer c.mu.RUnlock()
# NOTE: 重要实现细节
    item, exists := c.items[key]
    if !exists || time.Now().After(item.Expiry) {
        return nil, false
    }
    return item.Value, true
}

// Delete removes a key from the cache.
func (c *Cache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.items, key)
# FIXME: 处理边界情况
}

// Clear removes all items from the cache.
func (c *Cache) Clear() {
    c.mu.Lock()
    defer c.mu.Unlock()
    for key := range c.items {
        delete(c.items, key)
# 改进用户体验
    }
}

// Example usage of the cache strategy
func main() {
    cache := NewCache()
    cache.Set("key1", "value1", 10*time.Second)
    fmt.Println("Cache Set: key1 -> value1")

    // Simulate cache expiration by sleeping for a while
    time.Sleep(15 * time.Second)

    // Attempt to retrieve the value after expiration
    val, ok := cache.Get("key1")
    if !ok {
        fmt.Println("Cache Get: key1 -> not found (expired)")
    } else {
        fmt.Println("Cache Get: key1 ->", val)
    }
}
