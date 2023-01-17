package douban

import (
	"time"

	"github.com/imroc/req/v3"
)

var dbClient = NewReqClient()

func NewReqClient() *req.Client {
	return req.C().
		//DevMode()
		SetBaseURL("https://frodo.douban.com/api/v2").
		SetCommonHeaders(map[string]string{
			"Referer":      "https://servicewechat.com/wx2f9b06c1de1ccfca/84/page-frame.html",
			"content-type": "application/json",
			"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, " +
				"like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9." +
				"501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat",
		}).
		SetCommonQueryParam("apikey", "0ac44ae016490db2204ce0a042db2916").
		SetTimeout(5 * time.Second)
}
