package routers

import (
	"TVHelper/global"
	"TVHelper/internal/middleware/zaplog"
	"TVHelper/internal/parser"

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

	r.GET("/vod", douBanHandler) // for t4 public
	r.GET("/home", douBanHandler)

	r.GET("/live/:file", func(c *gin.Context) {
		file := c.Param("file")
		c.File("configs/live/" + file)
	})

	return r
}
