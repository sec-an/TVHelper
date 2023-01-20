package main

import (
	"TVHelper/global"
	"TVHelper/internal/douban"
	"TVHelper/internal/parser"
	"TVHelper/internal/routers"
	"TVHelper/pkg/logging"
	"TVHelper/pkg/redis"
	"TVHelper/pkg/setting"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var dir string

func init() {
	// 获取当前路径
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&dir, "d", currDir, "TVHelper根目录")
	flag.Parse()
	// 切换至TVHelper根目录
	err = os.Chdir(dir)
	if err != nil {
		log.Fatal(err)
	}
	// 配置初始化
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// zap初始化
	logging.Init()
	// redis初始化
	if global.RedisSetting.Running {
		global.RedisClient = redis.Init()
		// 每次运行先清空缓存
		ctx := context.Background()
		res, err := global.RedisClient.FlushDB(ctx).Result()
		if err != nil {
			global.Logger.Panic("Failed to flush redis", zap.Error(err))
		}
		global.Logger.Info("FlushDB: " + res)
	}
	// 豆瓣、配置解析客户端初始化
	global.DouBanClient = douban.NewReqClient(global.SpiderSetting.DouBanClientTimeout)
	global.ParserClient = parser.NewReqClient(global.SpiderSetting.ParserClientTimeout)
}

func main() {
	global.Logger.Info("TVHelper starting...",
		zap.String("port", global.ServerSetting.HttpPort),
		zap.String("dir", dir))

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		global.Logger.Fatal("startup service failed...", zap.Error(err))
	}
}

func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Log", &global.LogSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Spider", &global.SpiderSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.RedisSetting.IdleTimeout *= time.Second
	global.RedisSetting.SubCacheTime *= time.Minute
	global.SpiderSetting.DouBanClientTimeout *= time.Millisecond
	global.SpiderSetting.ParserClientTimeout *= time.Millisecond
	return nil
}
