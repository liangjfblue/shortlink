/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package redis

import "github.com/chasex/redis-go-cluster"

func Del(key string) error {
	if _, err := _cluster.Do("DEL", key); err != nil {
		return err
	}
	return nil
}

func MDel(keys ...string) (int, error) {
	return redis.Int(_cluster.Do("DEL", keys))
}

func Exists(key string) (bool, error) {
	return redis.Bool(_cluster.Do("EXISTS", key))
}

func MExists(keys string) (int, error) {
	return redis.Int(_cluster.Do("EXISTS", keys))
}

func Expire(key string, seconds int) error {
	if _, err := _cluster.Do("EXPIRE", key, seconds); err != nil {
		return err
	}
	return nil
}

func RenameNX(key, nKey string) error {
	if _, err := _cluster.Do("RENAMENX", key, nKey); err != nil {
		return err
	}
	return nil
}

func TTL(key string) (int, error) {
	return redis.Int(_cluster.Do("TTL", key))
}

func Type(key string) (string, error) {
	return redis.String(_cluster.Do("TYPE", key))
}
