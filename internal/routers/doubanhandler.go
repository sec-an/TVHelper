package routers

import (
	"TVHelper/global"
	"TVHelper/internal/common"
	"TVHelper/internal/douban"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

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
		if global.RedisSetting.Running {
			keyName := strings.Join([]string{t, ext, pg}, "_")
			res := &common.Result{}
			if t == "0interests" {
				keyName = strings.Join([]string{keyName, dbId}, "_")
			}
			cacheStr, err := global.RedisClient.Get(c, keyName).Result()
			if err == redis.Nil {
				// 未命中缓存
				global.Logger.Debug(keyName, zap.Error(err))
				*res, err = douban.CateFilter(t, ext, pg, dbId)
				if err != nil {
					// 请求豆瓣异常
					global.Logger.Error("Error DouBan Api", zap.Error(err))
					c.PureJSON(502, gin.H{
						"error": fmt.Sprintf("%v", err),
					})
					return
				} else {
					data, _ := json.Marshal(*res)
					if t == "0interests" {
						_ = global.RedisClient.Set(c, keyName, data, 15*time.Minute).Err()
					} else {
						_ = global.RedisClient.Set(c, keyName, data, 2*time.Hour).Err()
					}
				}

			} else if err != nil {
				global.Logger.Error(keyName, zap.Error(err))
			} else {
				_ = json.Unmarshal([]byte(cacheStr), res)
			}
			c.PureJSON(200, res)
		} else {
			if res, err := douban.CateFilter(t, ext, pg, dbId); err != nil {
				global.Logger.Error("douBanCateFilter", zap.Error(err))
				c.PureJSON(502, gin.H{
					"error": fmt.Sprintf("%v", err),
				})
			} else {
				c.PureJSON(200, res)
			}
		}
		return
	}

	var subjectRealTimeHotest []common.Vod

	// 实时热门，返回首页数据
	if global.RedisSetting.Running {
		realTimeHotestStr, err := global.RedisClient.Get(c, "real_time_hotest").Result()
		if err == redis.Nil {
			global.Logger.Debug("subjectRealTimeHotest", zap.Error(err))
			subjectRealTimeHotest = getSubjectRealTimeHotest()
			data, _ := json.Marshal(subjectRealTimeHotest)
			_ = global.RedisClient.Set(c, "real_time_hotest", data, 30*time.Minute).Err()
		} else if err != nil {
			global.Logger.Error("subjectRealTimeHotest", zap.Error(err))
		} else {
			_ = json.Unmarshal([]byte(realTimeHotestStr), &subjectRealTimeHotest)
		}
	} else {
		subjectRealTimeHotest = getSubjectRealTimeHotest()
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

func getSubjectRealTimeHotest() (subjectRealTimeHotest []common.Vod) {
	subjectRealTimeHotest, err := douban.SubjectRealTimeHotest()
	if err != nil {
		global.Logger.Error("subjectRealTimeHotest", zap.Error(err))
	}
	return
}
