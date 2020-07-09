/**
 *
 * @author liangjf
 * @create on 2020/7/9
 * @version 1.0
 */
package redis

import "github.com/chasex/redis-go-cluster"

func SAdd(key string, member interface{}) (int, error) {
	return redis.Int(_cluster.Do("SADD", key, member))
}

func Scard(key string) (int, error) {
	return redis.Int(_cluster.Do("SCARD", key))
}

func SIsMember(key string, member interface{}) (bool, error) {
	return redis.Bool(_cluster.Do("SISMEMBER", key, member))
}

func SMembers(key string) ([]string, error) {
	return redis.Strings(_cluster.Do("SMEMBERS", key))
}

func SPopString(key string) (string, error) {
	return redis.String(_cluster.Do("SPOP", key, 1))
}

func SPopNString(key string, count int) ([]string, error) {
	return redis.Strings(_cluster.Do("SPOP", key, count))
}

func SPopInt(key string) (int, error) {
	return redis.Int(_cluster.Do("SPOP", key, 1))
}

func SPopNInt(key string, count int) ([]int, error) {
	return redis.Ints(_cluster.Do("SPOP", key, count))
}

func SRandMemberString(key string) (string, error) {
	return redis.String(_cluster.Do("SRANDMEMBER", key))
}

func SRandMemberInt(key string) (int, error) {
	return redis.Int(_cluster.Do("SRANDMEMBER", key))
}

func SRandMemberNString(key string, count int) ([]string, error) {
	return redis.Strings(_cluster.Do("SRANDMEMBER", key, count))
}

func SRandMemberNInt(key string, count int) ([]int, error) {
	return redis.Ints(_cluster.Do("SRANDMEMBER", key, count))
}

func SRem(key string, member interface{}) (int, error) {
	return redis.Int(_cluster.Do("SREM", key, member))
}
