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
	KeyShortLinkIdPrefix = "shortLink:shortId:"
	//长连接code前缀
	KeyShortLinkLongMd5Prefix = "shortLink:LongMd5:"
)

var (
	//长短连接缓存过期时间
	ExpireTime = 3600 //3600

	//短码id起始位置
	MinShortId int64 = 1000
)

func AddRandExpire(expireTime int) int {
	return rand.Intn(10) + expireTime
}
