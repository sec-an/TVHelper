package routers

import (
	"TVHelper/global"
	"TVHelper/internal/common"
	"TVHelper/internal/douban"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func douBanHandler(c *gin.Context) {
	dbId := c.Query("douban")

	// 屏蔽搜素和播放请求
	if c.Query("wd") != "" || c.Query("play") != "" {
		c.PureJSON(403, gin.H{
			"error": "豆瓣主页不提供搜索及播放功能",
		})
		return
	}

	// 分类筛选
	if t, ext, pg := c.Query("t"), c.Query("ext"), c.Query("pg"); t != "" {
		if res, err := douban.CateFilter(t, ext, pg, dbId); err != nil {
			global.Logger.Error("豆瓣分类筛选", zap.Error(err))
			c.PureJSON(502, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
		} else {
			c.PureJSON(200, res)
		}
		return
	}

	// 实时热门，返回首页数据
	subjectRealTimeHotest, err := douban.SubjectRealTimeHotest()
	if err != nil {
		global.Logger.Error("subjectRealTimeHotest", zap.Error(err))
	}
	result := common.Result{
		Class:   douban.GetDbClass(),
		Filters: douban.GetDbFilter(),
		List:    subjectRealTimeHotest,
	}

	// 未提供豆瓣id，删去数据中“我的豆瓣”分类及筛选
	if dbId == "" {
		result.Class = result.Class[1:]
		delete(result.Filters, "0interests")
	}

	c.PureJSON(200, result)
}
