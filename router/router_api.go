// Package router
// @author tabuyos
// @since 2023/7/10
// @description router
package router

import (
	"github.com/gin-gonic/gin"
)

// getCurrentVersionRouter 获取当前的版本路由
func getCurrentVersionRouter(handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return getV1Router(handlers...)
}

// getV1Router 获取 V1 版本路由
func getV1Router(handlers ...gin.HandlerFunc) *gin.RouterGroup {
	v1Group := baseRouter.Group("/api/v1", handlers...)
	return v1Group
}

func setApiRouter() {
	// // 匿名(无需认证) api
	// setAnonymousApi(getCurrentVersionRouter())
	// // 公共 api
	// setCommonApi(getCurrentVersionRouter(middleware.CheckAuth(middleware.NewTokenAuth())))
	// // 平台 api
	// setPlatformApi(getCurrentVersionRouter(middleware.CheckAuth(middleware.NewTokenAuth())))
	// // 用户(实名) api
	// setUserApi(getCurrentVersionRouter(middleware.CheckAuth(middleware.NewTokenAuth())))
	// // 租户 api
	// setTenantApi(getCurrentVersionRouter(middleware.CheckAuth(middleware.NewTokenAuth())))
	// // 系统 api
	// setSystemApi(getCurrentVersionRouter(middleware.CheckAuth(middleware.NewTokenAuth())))
}
