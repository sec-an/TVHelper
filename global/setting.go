package global

import (
	"TVHelper/pkg/setting"

	"github.com/go-redis/redis/v8"

	"github.com/imroc/req/v3"

	"go.uber.org/zap"
)

var (
	ServerSetting *setting.ServerSettingS
	LogSetting    *setting.LogSettingS
	RedisSetting  *setting.RedisSettingS
	SpiderSetting *setting.SpiderSettingS
	Logger        *zap.Logger
	DouBanClient  *req.Client
	ParserClient  *req.Client
	RedisClient   *redis.Client
)
