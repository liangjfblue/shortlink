/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package redis

import (
	"testing"

	"github.com/chasex/redis-go-cluster"
)

func TestRedisPool(t *testing.T) {
	err := Init([]string{"172.16.7.16:8001", "172.16.7.16:8002", "172.16.7.16:8003"})
	if err != nil {
		t.Fatal(err)
	}

	if err = Set("name", "aaa"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(redis.String(Get("name")))
	}
}
