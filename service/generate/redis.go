/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package generate

import (
	"shortlink/common"
	"shortlink/models/redis"

	"github.com/sirupsen/logrus"
)

/**
基于redis原子递增实现短码id的生成
*/
type redisGenerate struct {
	opts Options
}

func NewRedisGenerate() IGenerate {
	return &redisGenerate{}
}

//生成短码id
func (g *redisGenerate) Create(opts ...Option) (int64, error) {
	var (
		err     error
		shortId int64
	)

	g.opts = defaultOptions
	for _, o := range opts {
		o(&g.opts)
	}

	//TODO lua script
	ok, err := redis.Exists(common.KeyShortLinkCreateId)
	if !ok || err != nil {
		shortId, err = redis.IncrBy(common.KeyShortLinkCreateId, common.MinShortId)
		if err != nil {
			logrus.Errorf("incr key err:%s", err.Error())
			return -1, err
		}
	}

	shortId, err = redis.Incr(common.KeyShortLinkCreateId)
	if err != nil {
		logrus.Errorf("incr key err:%s", err.Error())
		return -1, err
	}

	return shortId, nil
}
