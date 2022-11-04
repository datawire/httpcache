package test_test

import (
	"testing"

	"github.com/datawire/httpcache"
	"github.com/datawire/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
