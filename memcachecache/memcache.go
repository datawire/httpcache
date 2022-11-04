// Package memcachecache provides an implementation of httpcache.Cache
// that uses memcache to store cached responses.
//
// New() and NewWithClient() connect to a normal memcache server;
// NewWithAppEngine connects to Google App Engine memcache instance.
package memcachecache

import (
	"github.com/bradfitz/gomemcache/memcache"

	"github.com/datawire/httpcache"
)

// memcacheCache is an implementation of httpcache.Cache that caches responses
// in a memcache server.
type memcacheCache struct {
	*memcache.Client
}

// cacheKey modifies an httpcache key for use in memcache.  Specifically, it
// prefixes keys to avoid collision with other data stored in memcache.
func cacheKey(key string) string {
	return "httpcache:" + key
}

// Get returns the response corresponding to key if present.
func (c *memcacheCache) Get(key string) (resp []byte, ok bool) {
	item, err := c.Client.Get(cacheKey(key))
	if err != nil {
		return nil, false
	}
	return item.Value, true
}

// Set saves a response to the cache as key.
func (c *memcacheCache) Set(key string, resp []byte) {
	item := &memcache.Item{
		Key:   cacheKey(key),
		Value: resp,
	}
	c.Client.Set(item)
}

// Delete removes the response with key from the cache.
func (c *memcacheCache) Delete(key string) {
	c.Client.Delete(cacheKey(key))
}

// New returns a new httpcache.Cache using the provided memcache server(s) with
// equal weight. If a server is listed multiple times, it gets a proportional
// amount of weight.
func New(server ...string) httpcache.Cache {
	return NewWithClient(memcache.New(server...))
}

// NewWithClient returns a new httpcache.Cache with the given memcache client.
func NewWithClient(client *memcache.Client) httpcache.Cache {
	return &memcacheCache{client}
}
