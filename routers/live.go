package routers

import (
	"github.com/gin-gonic/gin"
)

func LoadLive(e *gin.Engine) {
	e.GET("/live/:file", func(c *gin.Context) {
		file := c.Param("file")
		c.File("live/" + file)
	})
}
