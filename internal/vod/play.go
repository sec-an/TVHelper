package vod

import (
	"TVHelper/global"
	"net/url"

	"github.com/tidwall/gjson"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetPlayContent(c *gin.Context) {
	var signedUrl string
	if id := c.Query("id"); id != "" {
		row := global.Mysql.QueryRow("select sp_path from vod where vod_id=?", id)
		var data Vod
		err := row.Scan(&data.SpPath)
		if err != nil {
			global.Logger.Error(id, zap.Error(err))
		}
		resp, err := global.AListClient.R().SetBody(&RequestBody{Path: data.SpPath}).Post("/fs/get")
		if err != nil {
			global.Logger.Error(id, zap.Error(err))
		}
		respStr := resp.String()
		sign := gjson.Get(respStr, "data.sign").String()
		u, _ := url.Parse("https://cdn.sec-an.cn/d" + data.SpPath)
		q := u.Query()
		q.Set("sign", sign)
		u.RawQuery = q.Encode()
		signedUrl = u.String()
	}
	c.PureJSON(200, gin.H{
		"url": signedUrl,
	})
}
