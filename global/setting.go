package global

import (
	"TVHelper/pkg/setting"
	"database/sql"

	"github.com/go-redis/redis/v8"

	"github.com/imroc/req/v3"

	"go.uber.org/zap"
)

var (
	ServerSetting *setting.ServerSettingS
	LogSetting    *setting.LogSettingS
	MysqlSetting  *setting.MysqlSettingS
	RedisSetting  *setting.RedisSettingS
	SpiderSetting *setting.SpiderSettingS
	AListSetting  *setting.AListSettingS
	Logger        *zap.Logger
	DouBanClient  *req.Client
	AListClient   *req.Client
	ParserClient  *req.Client
	RedisClient   *redis.Client
	Mysql         *sql.DB
)
