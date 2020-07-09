/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package router

import (
	"shortlink/config"
	"shortlink/controllers/sd"
	"shortlink/controllers/shortlink"
	"shortlink/router/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(config.Config().Server.RunMode)
	g := gin.Default()

	g.GET("/sl/:shortCode", shortlink.Get)

	sdV1 := g.Group("/v1/sd")
	sdV1.Use()
	{
		sdV1.GET("/health", sd.Health)
	}

	shortLinkV1 := g.Group("/v1/shorten")
	shortLinkV1.Use(middleware.LimitMiddleware())
	{
		shortLinkV1.POST("", shortlink.Shorten)
		shortLinkV1.GET("", shortlink.Info)
	}

	return g
}
