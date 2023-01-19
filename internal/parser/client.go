package parser

import (
	"time"

	"github.com/imroc/req/v3"
)

var parserClient = NewReqClient()

func NewReqClient() *req.Client {
	return req.C().
		//DevMode().
		SetUserAgent("okhttp/5.0.0-alpha.10").
		SetTimeout(5 * time.Second)
}
