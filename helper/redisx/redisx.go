// Package redisx
// @author tabuyos
// @since 2023/8/2
// @description redis
package redisx

import (
	"context"
	"deepsea/config"
	"deepsea/helper/recorderx"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var client redis.UniversalClient

type redisLogger struct {
	recorder recorderx.Recorder
}

func InitRedisX() {
	recorder := recorderx.DefaultRecorder()
	recorder.Info("初始化 Redis...")

	redisConfig := config.TomlConfig().Redis

	client = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    redisConfig.Addrs,
		Username: redisConfig.Username,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	ping := client.Ping(context.Background())
	result, err := ping.Result()
	if err != nil {
		recorder.Info("Redis 初始化失败...")
		panic(err)
	}
	recorder.Info(fmt.Sprintf("Ping 响应结果: %s", result))
	recorder.Info("Redis 初始化完成...")
}

func FetchRedisX() redis.UniversalClient {
	return client
}
