package test_test

import (
	"testing"

	"github.com/datawire/httpcache/memorycache"
	"github.com/datawire/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, memorycache.New())
}
