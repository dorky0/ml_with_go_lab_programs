package main
import (
"fmt"
"time"
)
// Cache is a simple in-memory cache implementation

type Cache struct {
data map[string]interface{}
expiration map[string]time.Time
}
// Set adds a value to the cache with the specified key and expiration duration
func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
c.data[key] = value
c.expiration[key] = time.Now().Add(expiration)
}
// Get retrieves a value from the cache based on the specified key
func (c *Cache) Get(key string) (interface{}, bool) {
value, found := c.data[key]
if !found {
return nil, false
}
expTime, expFound := c.expiration[key]
if !expFound || time.Now().After(expTime) {
// Value expired or expiration not found, delete it
delete(c.data, key)
delete(c.expiration, key)
return nil, false
}
return value, true
}
func main() {
// Create a new cache
cache := &Cache{
data: make(map[string]interface{}),
expiration: make(map[string]time.Time),
}
// Set a value in the cache with a 5-second expiration time
cache.Set("key1", "value1", 5*time.Second)
// Retrieve the value from the cache
value, found := cache.Get("key1")
if found {
fmt.Println("Value:", value)
} else {
fmt.Println("Value not found")
}

// Wait for 6 seconds to let the value expire
time.Sleep(6 * time.Second)
// Try to retrieve the expired value
value, found = cache.Get("key1")
if found {
fmt.Println("Value:", value)
} else {
fmt.Println("Value not found")
}
}