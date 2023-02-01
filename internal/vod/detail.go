package vod

import (
	"TVHelper/global"
	"TVHelper/internal/common"
	"fmt"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func GetDetailContent(c *gin.Context) {
	var res common.Vod
	if id := c.Query("id"); id != "" {
		rows, err := global.Mysql.Query("select vod_id,version,episode,sp_path,cme_material_id,"+
			"cme_signature,size,width,height,fps from vod where douban_id=? order by version,episode", id)
		if err != nil {
			global.Logger.Error(id, zap.Error(err))
			c.PureJSON(502, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}
		var dataList = make([]Vod, 0)
		for rows.Next() {
			var data Vod
			err := rows.Scan(&data.VodId, &data.Version, &data.Episode, &data.SpPath, &data.CmeMaterialId,
				&data.CmeSignature, &data.Size, &data.Width, &data.Height, &data.Fps)
			if err != nil {
				global.Logger.Error(id, zap.Error(err))
			}
			dataList = append(dataList, data)
		}
		var cmeOriginUrls, cmeM3U8Urls, spOriginUrls []string
		for _, data := range dataList {
			fileName := data.SpPath[strings.LastIndex(data.SpPath, "/")+1:]
			cmeOriginUrls = append(cmeOriginUrls, strings.Join([]string{strconv.Itoa(data.Width), "×", strconv.Itoa(data.Height), "@",
				strconv.Itoa(data.Fps), "fps ", humanize.Bytes(data.Size), "$", data.CmeMaterialId, "@",
				data.CmeSignature}, ""))
			cmeM3U8Urls = append(cmeM3U8Urls, strings.Join([]string{"可选分辨率 ", fileName, "$", data.CmeMaterialId, "@",
				data.CmeSignature}, ""))
			spOriginUrls = append(spOriginUrls, strings.Join([]string{strconv.Itoa(data.Width), "×", strconv.Itoa(data.Height), "@",
				strconv.Itoa(data.Fps), "fps ", humanize.Bytes(data.Size), "$", strconv.Itoa(data.VodId)}, ""))
		}
		cmeOrigin := strings.Join(cmeOriginUrls, "#")
		cmeM3U8 := strings.Join(cmeM3U8Urls, "#")
		spOrigin := strings.Join(spOriginUrls, "#")
		res.TypeName = "tv"
		if strings.HasPrefix(dataList[0].SpPath, "/movie") {
			res.TypeName = "movie"
		}
		res.VodPlayFrom = "CME原画$$$CMEm3u8$$$SP原画"
		res.VodPlayUrl = strings.Join([]string{cmeOrigin, cmeM3U8, spOrigin}, "$$$")
	}

	c.PureJSON(200, res)
}
