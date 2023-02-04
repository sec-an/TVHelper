package parser

import (
	"TVHelper/global"
	"TVHelper/internal/common"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"

	"github.com/DisposaBoy/JsonConfigReader"

	"github.com/gin-gonic/gin"
)

func ConfigHandler(c *gin.Context) {
	filename := c.Param("filename")

	// 读取配置文件，保存着config目录下
	configFile, err := os.Open("configs/config/" + filename + ".json")
	if err != nil {
		global.Logger.Error(filename, zap.Error(err))
		c.PureJSON(404, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			global.Logger.Error(filename, zap.Error(err))
		}
	}(configFile)

	// 解析配置
	var parser Parser
	configWithOutComment := JsonConfigReader.New(configFile) // 过滤注释
	err = json.NewDecoder(configWithOutComment).Decode(&parser)
	if err != nil {
		global.Logger.Error(filename+":解析配置出错", zap.Error(err))
		c.PureJSON(500, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	var subscribe common.Config

	if global.RedisSetting.Running {
		keyName := strings.Join([]string{"config", filename}, "_")
		cacheStr, err := global.RedisClient.Get(c, keyName).Result()
		if err == redis.Nil {
			global.Logger.Debug(keyName, zap.Error(err))
			subscribe = getScribe(parser)
			data, _ := json.Marshal(subscribe)
			_ = global.RedisClient.Set(c, keyName, data, global.RedisSetting.SubCacheTime).Err()
		} else if err != nil {
			global.Logger.Error(keyName, zap.Error(err))
		} else {
			_ = json.Unmarshal([]byte(cacheStr), &subscribe)
		}
	} else {
		subscribe = getScribe(parser)
	}

	c.PureJSON(200, subscribe)
}

func getScribe(parser Parser) (subscribe common.Config) {
	for _, itemSubscribe := range parser.Subscribe {
		tmpSubscribe := &common.Config{}
		// 已存在有效订阅，且当前订阅非强制展示
		if len(subscribe.Sites) != 0 && !itemSubscribe.AlwaysOn {
			continue
		}
		if data := getJson(itemSubscribe.Url); data != "" {
			err := json.Unmarshal([]byte(data), tmpSubscribe)
			if err != nil {
				// 订阅失效，更换下一订阅
				global.Logger.Error(itemSubscribe.Url+":订阅失效", zap.Error(err))
				continue
			}
			// 该订阅是否使用自定义jar
			isMultiJar := itemSubscribe.MultiJar
			// 若该订阅为第一个有效订阅
			//if len(subscribe.Sites) == 0 {
			//	isMultiJar = false
			//}
			// 判断是否需要修改订阅站点：是否多jar、自定义点播源前缀、点播源白名单、黑名单
			sitesNeedModify := isMultiJar || itemSubscribe.SitesPrefix != "" || len(itemSubscribe.SitesWhitelist) != 0 || len(itemSubscribe.SitesBlacklist) != 0
			if sitesNeedModify {
				sitesFinal := make([]common.Site, 0, len(subscribe.Sites))
				for iy, itemSite := range tmpSubscribe.Sites {
					// 该点播源未配置多jar
					if isMultiJar && itemSite.Jar == "" {
						if itemSubscribe.Jar != "" {
							// 配置文件中指定了该订阅的jar
							tmpSubscribe.Sites[iy].Jar = itemSubscribe.Jar
						} else {
							// 使用该订阅中的spider作为jar
							tmpSubscribe.Sites[iy].Jar = tmpSubscribe.Spider
						}
					}
					if itemSubscribe.SitesPrefix != "" {
						tmpSubscribe.Sites[iy].Name = strings.Join([]string{itemSubscribe.SitesPrefix, itemSite.Name}, "")
					}
					// 白名单非空且存在，添加点播源且忽略黑名单
					if len(itemSubscribe.SitesWhitelist) != 0 {
						if find(itemSubscribe.SitesWhitelist, itemSite.Key) {
							sitesFinal = append(sitesFinal, tmpSubscribe.Sites[iy])
						}
					} else if len(itemSubscribe.SitesBlacklist) != 0 {
						if !find(itemSubscribe.SitesBlacklist, itemSite.Key) {
							sitesFinal = append(sitesFinal, tmpSubscribe.Sites[iy])
						}
					} else {
						sitesFinal = append(sitesFinal, tmpSubscribe.Sites[iy])
					}
				}
				tmpSubscribe.Sites = sitesFinal
			}
			if len(subscribe.Sites) == 0 {
				// 第一个有效订阅
				subscribe = *tmpSubscribe
			} else {
				// 合并订阅
				subscribe.Ads = append(subscribe.Ads, tmpSubscribe.Ads...)
				subscribe.Flags = append(subscribe.Flags, tmpSubscribe.Flags...)
				subscribe.Parses = append(subscribe.Parses, tmpSubscribe.Parses...)
				subscribe.Sites = append(subscribe.Sites, tmpSubscribe.Sites...)
			}
		}
	}
	// 主jar替换
	if parser.Spider != "" {
		subscribe.Spider = parser.Spider
	}
	// 壁纸替换
	if parser.Wallpaper != "" {
		subscribe.Wallpaper = parser.Wallpaper
	}
	// 直播替换
	if len(parser.Lives) != 0 {
		subscribe.Lives = parser.Lives
	}
	// 去重
	subscribe.Ads = duplicateRemoving(append(subscribe.Ads, parser.MixAds...))
	subscribe.Flags = duplicateRemoving(append(subscribe.Flags, parser.MixFlags...))
	subscribe.Parses = duplicateRemoving(append(subscribe.Parses, parser.MixParses...))
	subscribe.Sites = append(parser.SitesPrepend, subscribe.Sites...)
	subscribe.Sites = duplicateRemoving(append(subscribe.Sites, parser.SitesAppend...))
	return
}

func Decoder(c *gin.Context) {
	url := c.DefaultQuery("url", "")
	if url == "" {
		c.PureJSON(200, gin.H{
			"error": "请检查url准确性",
		})
		return
	}
	data := getJson(url)
	if data == "" {
		c.PureJSON(200, gin.H{
			"error": "未解析到配置",
		})
		return
	}
	c.String(200, data)
}
