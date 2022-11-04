package memcachecache

import (
	"net"
	"testing"

	"github.com/datawire/httpcache/test"
)

const testServer = "localhost:11211"

func TestMemCache(t *testing.T) {
	conn, err := net.Dial("tcp", testServer)
	if err != nil {
		t.Fatal(err)
	}
	conn.Write([]byte("flush_all\r\n")) // flush memcache
	conn.Close()

	test.Cache(t, New(testServer))
}
