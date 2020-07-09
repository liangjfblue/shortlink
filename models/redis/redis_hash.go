/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package redis

import "github.com/chasex/redis-go-cluster"

func HGet(key, field string) (interface{}, error) {
	return _cluster.Do("HGET", key, field)
}

func HGetString(key, field string) (string, error) {
	return redis.String(_cluster.Do("HGET", key, field))
}

func HGetAll(key string) (map[string]string, error) {
	return redis.StringMap(_cluster.Do("HGETALL", key))
}

func HKeys(key string) ([]string, error) {
	return redis.Strings(_cluster.Do("HKEYS", key))
}

func HLen(key string) (int, error) {
	return redis.Int(_cluster.Do("HLEN", key))
}

func HMGet(key string, fields ...string) ([]string, error) {
	reply, err := redis.Strings(_cluster.Do("HMGET", key))
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func HSet(key, field string, value interface{}) error {
	if _, err := _cluster.Do("HSET", key, field, value); err != nil {
		return err
	}
	return nil
}

func HMSet(key string, pairs ...string) error {
	if _, err := _cluster.Do("HSET", key, pairs); err != nil {
		return err
	}
	return nil
}

func HSetNX(key, field string, value interface{}) (bool, error) {
	return redis.Bool(_cluster.Do("HSETNX", key, field, value))
}

func HIncrBy(key, field string, increment int) (int64, error) {
	return redis.Int64(_cluster.Do("HINCRBY", key, field, increment))
}

func HDel(key string, fields ...string) (int, error) {
	return redis.Int(_cluster.Do("HDEL", key, fields))
}

func HExist(key, field string) (bool, error) {
	return redis.Bool(_cluster.Do("HEXISTS", key, field))
}

func HStrLen(key, field string) (int, error) {
	return redis.Int(_cluster.Do("HSTRLEN", key, field))
}
