// 代码生成时间: 2025-08-04 10:17:25
package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

// CacheItem defines the structure of an item in the cache.
type CacheItem struct {
    Value      interface{}
    Expiration time.Time
}

// Cache defines the structure of the cache.
type Cache struct {
    sync.RWMutex
    items map[string]CacheItem
    ttl  time.Duration
}

// NewCache creates a new cache with a given TTL.
func NewCache(ttl time.Duration) *Cache {
    return &Cache{
        items: make(map[string]CacheItem),
        ttl:  ttl,
    }
}

// Set sets a value in the cache with the given key.
func (c *Cache) Set(key string, value interface{}) {
    c.Lock()
    defer c.Unlock()
    c.items[key] = CacheItem{
        Value: value,
        Expiration: time.Now().Add(c.ttl),
    }
}

// Get retrieves a value from the cache by its key.
func (c *Cache) Get(key string) (interface{}, bool) {
    c.RLock()
    defer c.RUnlock()
    item, exists := c.items[key]
    if !exists || time.Now().After(item.Expiration) {
        return nil, false
    }
    return item.Value, true
}

// Remove removes a value from the cache by its key.
func (c *Cache) Remove(key string) {
    c.Lock()
    defer c.Unlock()
    delete(c.items, key)
}

// Clean removes all expired items from the cache.
func (c *Cache) Clean() {
    c.Lock()
    defer c.Unlock()
    for key, item := range c.items {
        if time.Now().After(item.Expiration) {
            delete(c.items, key)
        }
    }
}

func main() {
    cache := NewCache(5 * time.Minute)
    cache.Set("key", "value")
    value, found := cache.Get("key")
    if found {
        fmt.Println("Found in cache:", value)
    } else {
        fmt.Println("Not found in cache")
    }
    cache.Clean()
    // Test expiration by sleeping for more than TTL
    time.Sleep(6 * time.Minute)
    value, found = cache.Get("key\)
    if !found {
        fmt.Println("Cache expired")
    }
}
