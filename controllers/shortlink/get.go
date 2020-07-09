/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package shortlink

import (
	"net/http"
	"shortlink/controllers"
	"shortlink/service/core"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Get(c *gin.Context) {
	var (
		err    error
		result controllers.Result
	)

	shortCode := c.Param("shortCode")

	logrus.WithField("shortCode", shortCode).Debug("shortlink short code")

	longLink, err := core.GetLongLinkByShortLink(shortCode)
	if err != nil {
		logrus.Errorf("shortCode no LongLink err:%s", err.Error())
		result.Failure(c, controllers.ErrShortCodeNoLongLink)
		return
	}

	logrus.WithFields(logrus.Fields{
		"shortCode": shortCode,
		"longLink":  longLink,
	}).Debug("shortCode get longLink, and redirect to longLink")

	c.Redirect(http.StatusFound, longLink)
}
