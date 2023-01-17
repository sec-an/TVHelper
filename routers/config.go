package routers

import (
	"TVHelper/parser"

	"github.com/gin-gonic/gin"
)

func LoadConfig(e *gin.Engine) {
	conf := e.Group("/config")
	{
		conf.GET("/:filename", parser.ConfigHandler)
		conf.GET("/src/:path/:file", func(c *gin.Context) {
			path := c.Param("path")
			file := c.Param("file")
			c.File("source_config/" + path + "/" + file)
		})
	}
}
