/**
 *
 * @author liangjf
 * @create on 2020/7/9
 * @version 1.0
 */
package redis

import "github.com/chasex/redis-go-cluster"

func LLEN(key string) (int, error) {
	return redis.Int(_cluster.Do("LLEN", key))
}

func RPUSH(key string, value interface{}) (int, error) {
	return redis.Int(_cluster.Do("RPUSH", key, value))
}

func RPUSHX(key string, value interface{}) (int, error) {
	return redis.Int(_cluster.Do("RPUSHX", key, value))
}
