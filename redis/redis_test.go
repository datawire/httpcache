package redis

import (
	"testing"

	"github.com/datawire/httpcache/test"
	"github.com/gomodule/redigo/redis"
)

func TestRedisCache(t *testing.T) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	conn.Do("FLUSHALL")

	test.Cache(t, NewWithClient(conn))
}
