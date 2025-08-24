// 代码生成时间: 2025-08-25 01:52:57
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "sync"
    "golang.org/x/sync/singleflight"
)

// CacheService is the struct that handles caching.
type CacheService struct {
    cache    map[string]interface{}
    lock     *sync.Mutex
    group    singleflight.Group
    expiry   time.Duration
}

// NewCacheService creates a new CacheService with a specific expiry duration.
func NewCacheService(expiry time.Duration) *CacheService {
    return &CacheService{
        cache:  make(map[string]interface{}),
        lock:   &sync.Mutex{},
        expiry: expiry,
    }
}

// Get retrieves an item from the cache. If the item is not present or expired,
// it fetches the item from the source and updates the cache.
func (cs *CacheService) Get(ctx context.Context, key string, fetchFunc func() (interface{}, error)) (interface{}, error) {
    cs.lock.Lock()
    defer cs.lock.Unlock()

    // Check if the item is in the cache.
    if item, exists := cs.cache[key]; exists {
        return item, nil
    }

    // Use singleflight to manage concurrent calls for the same key.
    result, err, shared := cs.group.Do(key, func() (interface{}, error) {
        // Fetch the item from the source.
        fetchedItem, fetchErr := fetchFunc()
        if fetchErr != nil {
            return nil, fetchErr
        }

        // Store the fetched item in the cache with an expiry.
        cs.cache[key] = fetchedItem
        return fetchedItem, nil
    })

    if err != nil {
        log.Printf("Error fetching item from source: %v", err)
        return nil, err
    }

    // If shared, the result is cached and will be used by other goroutines.
    return result, nil
}

// Expire clears the cache for a specific key.
func (cs *CacheService) Expire(key string) {
    cs.lock.Lock()
    defer cs.lock.Unlock()
    delete(cs.cache, key)
}

// Clear flushes all items from the cache.
func (cs *CacheService) Clear() {
    cs.lock.Lock()
    defer cs.lock.Unlock()
    cs.cache = make(map[string]interface{})
}

// Example usage of CacheService.
func main() {
    // Create a new cache service with a 5-minute expiry.
    cacheService := NewCacheService(5 * time.Minute)

    // Define a fetch function that simulates fetching data.
    fetchFunc := func() (interface{}, error) {
        time.Sleep(2 * time.Second) // Simulate latency.
        return "Fetched data", nil
    }

    // Retrieve data from the cache or fetch if not present.
    data, err := cacheService.Get(context.Background(), "example_key", fetchFunc)
    if err != nil {
        log.Fatalf("Failed to get data: %v", err)
    }
    fmt.Println("Data:", data)

    // Expire the cache for a specific key.
    cacheService.Expire("example_key")

    // Clear all items from the cache.
    cacheService.Clear()
}
