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
	"shortlink/service/convert"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	var (
		err    error
		result controllers.Result
	)

	shortLink := c.Query("shortLink")

	strList := strings.Split(shortLink, "/")
	shortCode := strList[len(strList)-1]
	if shortCode == "" {
		logrus.WithField("shortLink", shortLink).Error("shortCode is empty")
		result.Failure(c, controllers.ErrParams)
		return
	}

	shortId := convert.Decode(shortCode).(int64)

	tBShortLink, err := db.GetTBShortLink(map[string]interface{}{"short_id": shortId})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"shortId":   shortId,
			"shortCode": shortCode,
			"shortLink": shortLink,
		}).Error("tBShortLink is empty")
		result.Failure(c, controllers.ErrShortCodeNoLongLink)
		return
	}

	result.Success(c, tBShortLink)
}
