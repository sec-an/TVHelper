package parser

import (
	"time"

	"github.com/imroc/req/v3"
)

func NewReqClient(timeout time.Duration) *req.Client {
	return req.C().
		//DevMode().
		SetUserAgent("okhttp/5.0.0-alpha.10").
		SetTimeout(timeout)
}
