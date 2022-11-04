// Package memcache provides an implementation of httpcache.Cache that stores
// responses in an in-memory map.
package memorycache

import (
	"sync"
)

// cache is an implemtation of Cache that stores responses in an in-memory map.
type cache struct {
	mu    sync.RWMutex
	items map[string][]byte
}

// Get returns the []byte representation of the response and true if present, false if not
func (c *cache) Get(key string) (resp []byte, ok bool) {
	c.mu.RLock()
	resp, ok = c.items[key]
	c.mu.RUnlock()
	return resp, ok
}

// Contains returns whether the cache contains a cached response.
func (c *cache) Contains(key string) bool {
	c.mu.RLock()
	_, ok := c.items[key]
	c.mu.RUnlock()
	return ok
}

// Set saves response resp to the cache with key
func (c *cache) Set(key string, resp []byte) {
	c.mu.Lock()
	c.items[key] = resp
	c.mu.Unlock()
}

// Delete removes key from the cache
func (c *cache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}

// New returns a new httpcache.Cache that will store items in an in-memory map
func New() *cache {
	c := &cache{items: map[string][]byte{}}
	return c
}
