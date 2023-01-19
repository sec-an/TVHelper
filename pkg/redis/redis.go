package redis

import (
	"TVHelper/global"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

func Init() *redis.Client {
	// Redis连接格式拼接
	redisAddr := fmt.Sprintf("%s:%d", global.RedisSetting.Host, global.RedisSetting.Port)
	// Redis 连接对象: NewClient将客户端返回到由选项指定的Redis服务器。
	redisClient := redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		Password:    global.RedisSetting.Auth,
		DB:          global.RedisSetting.Database,
		IdleTimeout: global.RedisSetting.IdleTimeout,
		PoolSize:    global.RedisSetting.PoolSize,
	})
	global.Logger.Info("Connecting Redis", zap.String("addr", redisAddr))

	// go-redis库v8版本相关命令都需要传递context.Context参数,Background 返回一个非空的Context,它永远不会被取消，没有值，也没有期限。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 验证是否连接到redis服务端
	res, err := redisClient.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("Redis Connect Failed!", zap.Error(err))
	}

	// 输出连接成功标识
	global.Logger.Info("Connect Successful!", zap.String("ping", res))
	return redisClient
}
