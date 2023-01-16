package routers

import (
	"TVHelper/common"
	"TVHelper/douban"

	"github.com/gin-gonic/gin"
)

func LoadDouBan(e *gin.Engine) {
	e.GET("/home", func(c *gin.Context) {
		dbId := c.Query("douban")

		if c.Query("wd") != "" || c.Query("play") != "" {
			c.String(404, "404 Not Found")
			return
		}

		if t, ext, pg := c.Query("t"), c.Query("ext"), c.Query("pg"); t != "" {
			c.JSON(200, douban.CateFilter(t, ext, pg, dbId))
			return
		}

		result := common.Result{
			Class:   douban.GetDbClass(),
			Filters: douban.GetDbFilter(),
			List:    douban.SubjectRealTimeHotest(),
		}

		if dbId == "" {
			result.Class = result.Class[1:]
			delete(result.Filters, "0interests")
		}

		c.JSON(200, result)
	})
}
