/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package redis

import "github.com/chasex/redis-go-cluster"

func Set(key string, value interface{}) error {
	if _, err := _cluster.Do("SET", key, value); err != nil {
		return err
	}
	return nil
}

func SetEX(key string, value interface{}, expireSecond int) error {
	if _, err := _cluster.Do("SET", key, value, "EX", expireSecond); err != nil {
		return err
	}
	return nil
}

func SetPX(key string, value interface{}, expireMS int) error {
	if _, err := _cluster.Do("SET", key, value, "PX", expireMS); err != nil {
		return err
	}
	return nil
}

func SetNX(key string, value interface{}) error {
	if _, err := _cluster.Do("SET", key, value, "NX"); err != nil {
		return err
	}
	return nil
}

func SetXX(key string, value interface{}) error {
	if _, err := _cluster.Do("SET", key, value, "XX"); err != nil {
		return err
	}
	return nil
}

func Get(key string) (interface{}, error) {
	reply, err := _cluster.Do("GET", key)
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func GetString(key string) (string, error) {
	return redis.String(_cluster.Do("GET", key))
}

func GetInt(key string) (int, error) {
	return redis.Int(_cluster.Do("GET", key))
}

func GetSet(key string, newValue interface{}) (interface{}, error) {
	return _cluster.Do("GETSET", key, newValue)
}

func Incr(key string) (int64, error) {
	return redis.Int64(_cluster.Do("INCR", key))
}

func IncrBy(key string, increment int64) (int64, error) {
	return redis.Int64(_cluster.Do("INCRBY", key, increment))
}

func Decr(key string) (int64, error) {
	reply, err := redis.Int64(_cluster.Do("DECR", key))
	if err != nil {
		return 0, err
	}

	return reply, nil
}

func DecrBy(key string, decrement int64) (int64, error) {
	return redis.Int64(_cluster.Do("DECRBY", key, decrement))
}

func StrLen(key string) (int, error) {
	return redis.Int(_cluster.Do("STRLEN", key))
}
