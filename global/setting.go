package global

import (
	"TVHelper/pkg/setting"

	"go.uber.org/zap"
)

var (
	ServerSetting *setting.ServerSettingS
	LogSetting    *setting.LogSettingS
	Logger        *zap.Logger
)
