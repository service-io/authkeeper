// Package router
// @author tabuyos
// @since 2023/8/11
// @description router
package router

import (
	"deepsea/middleware"
	userApi "deepsea/module/user/api"
	"github.com/gin-gonic/gin"
)

func setUserApi(version *gin.RouterGroup) {
	userModGroup := version.Group("/user")
	{
		userGroup := userModGroup.Group("/user")
		{
			userHandler := userApi.NewUserHandler()
			userGroup.GET("/whoami", userHandler.Whoami())
			userGroup.PUT("/add", middleware.Limiter(), userHandler.Add())
			userGroup.DELETE("/remove", middleware.Limiter(), userHandler.Remove())
			userGroup.POST("/base-modify", middleware.Limiter(), userHandler.Modify())
			userGroup.GET("/detail", userHandler.Detail())
			userGroup.POST("/list-page", userHandler.ListPage())
			userGroup.POST("/list-page-condition", userHandler.ListPageByCondition())
		}
		userSignGroup := userModGroup.Group("/user-sign")
		{
			userSignHandler := userApi.NewUserSignHandler()
			userSignGroup.GET("/whoami", userSignHandler.Whoami())
			userSignGroup.PUT("/add", middleware.Limiter(), userSignHandler.Add())
			userSignGroup.DELETE("/remove", middleware.Limiter(), userSignHandler.Remove())
			userSignGroup.POST("/base-modify", middleware.Limiter(), userSignHandler.Modify())
			userSignGroup.GET("/detail", userSignHandler.Detail())
			userSignGroup.POST("/list-page", userSignHandler.ListPage())
			userSignGroup.POST("/list-page-condition", userSignHandler.ListPageByCondition())
		}
		userSignLogGroup := userModGroup.Group("/user-sign-log")
		{
			userSignLogHandler := userApi.NewUserSignLogHandler()
			userSignLogGroup.GET("/whoami", userSignLogHandler.Whoami())
			userSignLogGroup.PUT("/add", middleware.Limiter(), userSignLogHandler.Add())
			userSignLogGroup.DELETE("/remove", middleware.Limiter(), userSignLogHandler.Remove())
			userSignLogGroup.POST("/base-modify", middleware.Limiter(), userSignLogHandler.Modify())
			userSignLogGroup.GET("/detail", userSignLogHandler.Detail())
			userSignLogGroup.POST("/list-page", userSignLogHandler.ListPage())
		}
	}
}
