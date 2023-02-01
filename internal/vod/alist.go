package vod

import (
	"TVHelper/global"
	"time"

	"github.com/imroc/req/v3"
)

type RequestBody struct {
	Path string `json:"path"`
}

func NewReqClient(timeout time.Duration) *req.Client {
	return req.C().
		//DevMode().
		SetBaseURL(global.AListSetting.Host+"/api").
		SetCommonHeader("authorization", global.AListSetting.Token).
		SetTimeout(timeout)
}
