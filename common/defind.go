/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package common

import "math/rand"

const (
	//短连接id自增key
	KeyShortLinkCreateId = "shortLink:creatId"
	//短码id前缀
	KeyShortLinkShortCodePrefix = "shortLink:shortCode:"
	//长连接code前缀
	KeyShortLinkLongMd5Prefix = "shortLink:LongMd5:"

	//etcd
	//发号器的目录
	EtcdCreateIdDir = "/shortLink/"

	//分布式锁
	DistributedLock = "/shortLink/distributed_lock"
)

var (
	//长短连接缓存过期时间
	ExpireTime = 10 //3600

	//短码id起始位置
	MinShortId int64 = 1000

	//短码id递增区间
	ShortIdIncrStep int64 = 10
)

func AddRandExpire(expireTime int) int {
	return rand.Intn(10) + expireTime
}
