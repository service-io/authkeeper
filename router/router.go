// Package router
// @author tabuyos
// @since 2023/6/30
// @description router
package router

import (
	"deepsea/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var baseRouter *gin.Engine

func InitRouter() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		// 是否打印路由信息
	}

	baseRouter = gin.New()
	baseRouter.Use(middleware.CORS(), middleware.FillTrace(), middleware.RegisterLogger(), middleware.Visitor(), middleware.Recovery())
	baseRouter.NoRoute(func(ctx *gin.Context) {
		if ctx.Error(errors.New("未找到指定 API")) != nil {
		}
		ctx.String(http.StatusNotFound, "404 NOT FOUND!")
	})

	setApiRouter()
}

func BaseRouter() *gin.Engine {
	return baseRouter
}
