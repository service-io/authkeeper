package main

import (
	"context"
	"deepsea/config"
	"deepsea/config/env"
	"deepsea/docs"
	"deepsea/helper/database"
	"deepsea/helper/encodingx"
	"deepsea/helper/netx"
	"deepsea/helper/recorderx"
	"deepsea/helper/redisx"
	"deepsea/helper/snowflakeid"
	"deepsea/router"
	"errors"
	"fmt"
	"github.com/json-iterator/go/extra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//go:generate go mod tidy
//go:generate go mod download

func lowerCamelCase(f string) string {
	return strings.ToLower(f[:1]) + f[1:]
}

// @title           DEEPSEA PLATFORM
// @version         1.0
// @description     西太深海平台 api 文档.
// @termsOfService  http://www.deepseaqt.com

// @contact.name   API Support
// @contact.url    http://www.deepseaqt.com/support
// @contact.email  support@deepseaqt.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @basePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name DeepseaQt-Auth
func main() {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.InitEnv()
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	recorder := recorderx.DefaultRecorder()
	baseRouter := router.BaseRouter()

	address := netx.GetAddress()

	_, port := netx.ParseAddress(address)
	docs.SwaggerInfo.Host = netx.JoinPort(port)

	extra.RegisterFuzzyDecoders()
	extra.RegisterTimeAsInt64Codec(time.Millisecond)
	extra.SetNamingStrategy(lowerCamelCase)
	encodingx.RegisterInt64ToString()

	baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
		c.DocExpansion = "none"
	}))

	server := http.Server{
		Addr:    address,
		Handler: baseRouter.Handler(),
	}

	go func() {
		recorder.Info(fmt.Sprintf("服务监听-> %s\n", server.Addr))
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			recorder.Fatalf("服务监听错误: %s", err)
		}
	}()

	recorder.Info("服务启动成功")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 接收, 无信号时阻塞
	<-quit
	recorder.Info("服务即将关闭...")

	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channel()

	err := server.Shutdown(ctx)
	if err != nil {
		recorder.Fatal("服务关闭失败")
	}

	recorder.Info("服务退出...")

	// err := baseRouter.Run(address)
	// if err != nil {
	// 	return
	// }

}
