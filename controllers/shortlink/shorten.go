/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package shortlink

import (
	"shortlink/controllers"
	"shortlink/service/core"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ChangeReq struct {
	LongLink string `json:"longLink"`
}

func Shorten(c *gin.Context) {
	var (
		err    error
		req    ChangeReq
		result controllers.Result
	)

	if err := c.BindJSON(&req); err != nil {
		logrus.Errorf("bind json err:%s", err.Error())
		result.Failure(c, controllers.ErrParams)
		return
	}

	if req.LongLink == "" {
		logrus.Errorf("longLink empty")
		result.Failure(c, controllers.ErrParams)
		return
	}

	//shortLink in cache?
	shortLink, err := core.GetShortLinkByLongLink(req.LongLink)
	if err != nil {
		logrus.Errorf("get longLink err:%s", err.Error())
		result.Failure(c, controllers.ErrCreateShortLink)
		return
	}

	logrus.WithFields(logrus.Fields{
		"shortLink": shortLink,
	}).Debug("change longLink to shortLink")

	result.Success(c, shortLink)
}
