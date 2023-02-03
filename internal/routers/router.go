package routers

import (
	"TVHelper/global"
	"TVHelper/internal/live"
	"TVHelper/internal/middleware/zaplog"
	"TVHelper/internal/parser"
	"TVHelper/internal/vod"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(zaplog.GinLogger(global.Logger), zaplog.GinRecovery(global.Logger, true))

	conf := r.Group("/config")
	{
		conf.GET("/:filename", parser.ConfigHandler)
		conf.GET("/src/:path/:file", func(c *gin.Context) {
			path := c.Param("path")
			file := c.Param("file")
			c.File("configs/source_config/" + path + "/" + file)
		})
	}

	v := r.Group("/vod")
	{
		v.GET("/", douBanHandler) // for t4 public
		v.GET("/detail", vod.GetDetailContent)
		v.GET("/play", vod.GetPlayContent)
	}
	r.GET("/home", douBanHandler)

	l := r.Group("/live")
	{
		l.GET("/bestv/:channel/:bitrate/mnf.m3u8", live.BesTvHandler)
		l.GET("/:file", func(c *gin.Context) {
			file := c.Param("file")
			c.File("configs/live/" + file)
		})
	}

	return r
}
