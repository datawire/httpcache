//go:build appengine
// +build appengine

package memcache

import (
	"testing"

	"google.golang.org/appengine/v2/aetest"

	"github.com/datawire/httpcache/test"
)

func TestAppEngine(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	test.Cache(t, New(ctx))
}
