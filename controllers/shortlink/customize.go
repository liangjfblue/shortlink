/**
 *
 * @author liangjf
 * @create on 2020/7/9
 * @version 1.0
 */
package shortlink

import (
	"shortlink/controllers"
	"shortlink/models/db"
	"shortlink/service/core"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomizeReq struct {
	ShortCode string `json:"shortCode"`
	LongLink  string `json:"longLink"`
}

func Customize(c *gin.Context) {
	var (
		err    error
		req    CustomizeReq
		result controllers.Result
	)

	if err = c.BindJSON(&req); err != nil {
		logrus.Errorf("bind json err:%s", err.Error())
		result.Failure(c, controllers.ErrParams)
		return
	}

	if req.ShortCode == "" || req.LongLink == "" {
		logrus.WithFields(logrus.Fields{
			"shortCode": req.ShortCode,
			"longLink":  req.LongLink,
		}).Error("param error")
		result.Failure(c, controllers.ErrParams)
		return
	}

	if _, err := db.GetTBShortLink(map[string]interface{}{"short_code": req.ShortCode}); err == nil {
		result.Failure(c, controllers.ErrShortCodeHadLongLink)
		return
	}

	shortLink, err := core.CreateShortLinkByCustomizeShortCode(req.ShortCode, req.LongLink)
	if err != nil {
		logrus.Errorf("get longLink err:%s", err.Error())
		result.Failure(c, controllers.ErrCreateShortLink)
		return
	}

	result.Success(c, shortLink)
}
