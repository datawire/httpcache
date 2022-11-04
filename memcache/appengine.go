package memcache

import (
	"context"

	"google.golang.org/appengine/v2/log"
	"google.golang.org/appengine/v2/memcache"

	"github.com/datawire/httpcache"
)

// appengineCache is an implementation of httpcache.Cache that caches responses
// in App Engine's memcache.
type appengineCache struct {
	context.Context
}

// Get returns the response corresponding to key if present.
func (c *appengineCache) Get(key string) (resp []byte, ok bool) {
	item, err := memcache.Get(c.Context, cacheKey(key))
	if err != nil {
		if err != memcache.ErrCacheMiss {
			log.Errorf(c.Context, "error getting cached response: %v", err)
		}
		return nil, false
	}
	return item.Value, true
}

// Set saves a response to the cache as key.
func (c *appengineCache) Set(key string, resp []byte) {
	item := &memcache.Item{
		Key:   cacheKey(key),
		Value: resp,
	}
	if err := memcache.Set(c.Context, item); err != nil {
		log.Errorf(c.Context, "error caching response: %v", err)
	}
}

// Delete removes the response with key from the cache.
func (c *appengineCache) Delete(key string) {
	if err := memcache.Delete(c.Context, cacheKey(key)); err != nil {
		log.Errorf(c.Context, "error deleting cached response: %v", err)
	}
}

// New returns a new httpcache.Cache for the given context.
func NewWithAppEngine(ctx context.Context) httpcache.Cache {
	return &appengineCache{ctx}
}
