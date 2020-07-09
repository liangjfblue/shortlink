/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package sd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
