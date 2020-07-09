/**
 *
 * @author liangjf
 * @create on 2020/7/9
 * @version 1.0
 */
package middleware

import (
	"errors"
	"shortlink/config"
	"shortlink/controllers"
	"shortlink/models/redis"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func LimitRate(ip string) (bool, error) {
	current, err := redis.LLEN(ip)
	if err != nil {
		return false, err
	}

	if current >= config.Config().Server.OpenIpLimit {
		return true, errors.New("over limit")
	} else {
		if isExist, err := redis.Exists(ip); err == nil && !isExist {
			if _, err = redis.RPUSH(ip, 0); err != nil {
				return false, err
			}
			if err = redis.Expire(ip, 1); err != nil {
				return false, err
			}
		} else {
			if _, err = redis.RPUSHX(ip, 0); err != nil {
				return false, err
			}
		}
	}
	return false, nil
}

func LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Config().Server.OpenIpLimit > 0 {
			ip := strings.Split(c.Request.RemoteAddr, ":")[0]
			if over, err := LimitRate(ip); over {
				logrus.WithField("ip", ip).Errorf("over request limit, err:%s", err.Error())
				var result controllers.Result
				result.Failure(c, controllers.ErrLimitMiddleware)
				c.Abort()
			}
		}
		c.Next()
	}
}
