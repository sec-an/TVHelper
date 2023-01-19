package global

import (
	"TVHelper/pkg/setting"

	"github.com/imroc/req/v3"

	"go.uber.org/zap"
)

var (
	ServerSetting *setting.ServerSettingS
	LogSetting    *setting.LogSettingS
	SpiderSetting *setting.SpiderSettingS
	Logger        *zap.Logger
	DouBanClient  *req.Client
	ParserClient  *req.Client
)
