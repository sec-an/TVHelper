package parser

import (
	"TVHelper/common"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DisposaBoy/JsonConfigReader"

	"github.com/gin-gonic/gin"
)

func ConfigHandler(c *gin.Context) {
	id := c.Param("id")
	configFile, err := os.Open("config/" + id + ".json")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
		return
	}
	defer configFile.Close()
	var parser Parser
	configWithOutComment := JsonConfigReader.New(configFile)
	err = json.NewDecoder(configWithOutComment).Decode(&parser)
	if err != nil {
		fmt.Printf("失败 [Err:%s]\n", err.Error())
		return
	}
	subscribe := &common.Config{}
	for _, itemSubscribe := range parser.Subscribe {
		tmpSubscribe := &common.Config{}
		if len(subscribe.Sites) != 0 && !itemSubscribe.AlwaysOn { // 已有数据且非默认展示
			continue
		}
		if data := getJson(itemSubscribe.Url); data != "" {
			err = json.Unmarshal([]byte(data), tmpSubscribe)
			if err != nil {
				fmt.Printf("失败 [Err:%s]\n", err.Error())
				continue
			}
			isMultiJar := itemSubscribe.MultiJar
			if len(subscribe.Sites) == 0 {
				isMultiJar = false
			}
			if isMultiJar {
				for iy, itemSite := range tmpSubscribe.Sites {
					if itemSite.Jar == "" {
						if itemSubscribe.Jar != "" {
							tmpSubscribe.Sites[iy].Jar = itemSubscribe.Jar
						} else {
							tmpSubscribe.Sites[iy].Jar = tmpSubscribe.Spider
						}
					}
				}
			}
			if len(subscribe.Sites) == 0 {
				subscribe = tmpSubscribe
			} else {
				subscribe.Ads = append(subscribe.Ads, tmpSubscribe.Ads...)
				subscribe.Flags = append(subscribe.Flags, tmpSubscribe.Flags...)
				subscribe.Parses = append(subscribe.Parses, tmpSubscribe.Parses...)
				subscribe.Sites = append(subscribe.Sites, tmpSubscribe.Sites...)
			}
		}
	}
	if parser.Spider != "" {
		subscribe.Spider = parser.Spider
	}
	if parser.Wallpaper != "" {
		subscribe.Wallpaper = parser.Wallpaper
	}
	if len(parser.Lives) != 0 {
		subscribe.Lives = parser.Lives
	}
	subscribe.Ads = duplicateRemoving(append(subscribe.Ads, parser.MixAds...))
	subscribe.Flags = duplicateRemoving(append(subscribe.Flags, parser.MixFlags...))
	subscribe.Parses = duplicateRemoving(append(subscribe.Parses, parser.MixParses...))
	subscribe.Sites = append(parser.SitesPrepend, subscribe.Sites...)
	subscribe.Sites = duplicateRemoving(append(subscribe.Sites, parser.SitesAppend...))
	if len(parser.SitesBlacklist) != 0 {
		sitesFinal := make([]common.Site, 0, len(subscribe.Sites))
		for _, value := range subscribe.Sites {
			if !find(parser.SitesBlacklist, value.Name) {
				sitesFinal = append(sitesFinal, value)
			}
		}
		subscribe.Sites = sitesFinal
	}
	c.JSON(200, subscribe)
}
