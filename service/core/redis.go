/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package core

import (
	"crypto/md5"
	"errors"
	"fmt"
	"shortlink/common"
	"shortlink/config"
	"shortlink/models/db"
	"shortlink/models/redis"
	"shortlink/service/convert"
	"shortlink/service/generate"

	"github.com/sirupsen/logrus"
)

type redisWay struct{}

func newRedisWay() ICore {
	return &redisWay{}
}

func (s *redisWay) GetShortLinkByLongLink(longLink string) (string, error) {
	longLinkMd5 := fmt.Sprintf("%x", md5.Sum([]byte(longLink)))

	shortLink, err := redis.GetString(common.KeyShortLinkLongMd5Prefix + longLinkMd5)
	if err != nil {

		logrus.Debug("not in redis")

		//缓存过期或者不存在
		tBShortLink, err := db.GetTBShortLink(map[string]interface{}{"long_link_md5": longLinkMd5})
		if err == nil {
			shortLink = fmt.Sprintf("%s/%s", config.Config().ShortLink.ShortUrl, tBShortLink.ShortCode)
			if err := redis.SetEX(common.KeyShortLinkLongMd5Prefix+longLinkMd5, shortLink, common.AddRandExpire(common.ExpireTime)); err != nil {
				logrus.Errorf("save longLink and shortLink err:%s", err.Error())
				return "", err
			}
			goto END
		}

		//create short id
		shortId, err := generate.Create()
		if err != nil {
			logrus.Errorf("create short id err:%s", err.Error())
			return "", err
		}

		//short id convert to short code
		shortCode := convert.Encode(shortId)
		//short code make a shortlink
		shortLink = fmt.Sprintf("%s/%s", config.Config().ShortLink.ShortUrl, shortCode)

		//save to mysql
		if err := db.AddTBShortLink(&db.TBShortLink{
			ShortId:     uint64(shortId),
			ShortCode:   shortCode,
			LongUrl:     longLink,
			LongLinkMd5: longLinkMd5,
		}); err != nil {
			return "", err
		}

		if err := redis.SetEX(common.KeyShortLinkLongMd5Prefix+longLinkMd5, shortLink, common.ExpireTime); err != nil {
			logrus.Errorf("save longLink and shortLink err:%s", err.Error())
			return "", err
		}
	}

END:

	//update expire time
	if err := redis.Expire(common.KeyShortLinkLongMd5Prefix+longLinkMd5, common.AddRandExpire(common.ExpireTime)); err != nil {
		return "", err
	}

	return shortLink, nil
}

func (s *redisWay) GetLongLinkByShortLink(shortCode string) (string, error) {
	longLink, err := redis.GetString(common.KeyShortLinkShortCodePrefix + shortCode)
	if err != nil {
		tBShortLink, err := db.GetTBShortLink(map[string]interface{}{"short_code": shortCode})
		if err != nil {
			return "", err
		}

		if err := redis.SetEX(common.KeyShortLinkShortCodePrefix+shortCode, tBShortLink.LongUrl, common.AddRandExpire(common.ExpireTime)); err != nil {
			return "", err
		}

		longLink = tBShortLink.LongUrl
	}

	if longLink == "" {
		return "", errors.New("shortCode no longLink")
	}

	return longLink, nil
}

func (s *redisWay) CreateShortLinkByCustomizeShortCode(shortCode, longLink string) (string, error) {
	longLinkMd5 := fmt.Sprintf("%x", md5.Sum([]byte(longLink)))

	//short id convert to short code
	shortId := convert.Decode(shortCode).(int64)

	//save to mysql
	if err := db.AddTBShortLink(&db.TBShortLink{
		ShortId:     uint64(shortId),
		ShortCode:   shortCode,
		LongUrl:     longLink,
		LongLinkMd5: longLinkMd5,
	}); err != nil {
		return "", err
	}

	//short code make a shortlink
	shortLink := fmt.Sprintf("%s/%s", config.Config().ShortLink.ShortUrl, shortCode)

	return shortLink, nil
}
