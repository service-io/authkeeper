// Package router
// @author tabuyos
// @since 2023/8/11
// @description router
package router

import (
	"deepsea/middleware"
	commonApi "deepsea/module/common/api"
	"github.com/gin-gonic/gin"
)

func setCommonApi(version *gin.RouterGroup) {
	commonModGroup := version.Group("/common")
	{
		commonAreaGroup := commonModGroup.Group("/common-area")
		{
			commonAreaHandler := commonApi.NewCommonAreaHandler()
			commonAreaGroup.GET("/whoami", middleware.CheckRBAC(middleware.NewRole("COMMON")), commonAreaHandler.Whoami())
			commonAreaGroup.PUT("/add", middleware.Limiter(), commonAreaHandler.Add())
			commonAreaGroup.DELETE("/remove", middleware.Limiter(), commonAreaHandler.Remove())
			commonAreaGroup.POST("/base-modify", middleware.Limiter(), commonAreaHandler.Modify())
			commonAreaGroup.GET("/detail", commonAreaHandler.Detail())
			commonAreaGroup.POST("/list-page", commonAreaHandler.ListPage())
			commonAreaGroup.POST("/list-page-condition", commonAreaHandler.ListPageByCondition())
			commonAreaGroup.GET("/list", commonAreaHandler.FindList())
		}
		commonDictGroup := commonModGroup.Group("/common-dict")
		{
			commonDictHandler := commonApi.NewCommonDictHandler()
			commonDictGroup.GET("/whoami", commonDictHandler.Whoami())
			commonDictGroup.PUT("/add", middleware.Limiter(), commonDictHandler.Add())
			commonDictGroup.DELETE("/remove", middleware.Limiter(), commonDictHandler.Remove())
			commonDictGroup.POST("/base-modify", middleware.Limiter(), commonDictHandler.Modify())
			commonDictGroup.GET("/detail", commonDictHandler.Detail())
			commonDictGroup.POST("/list-page", commonDictHandler.ListPage())
			commonDictGroup.POST("/list-page-condition", commonDictHandler.ListPageByCondition())
		}
		commonDictDataGroup := commonModGroup.Group("/common-dict-data")
		{
			commonDictDataHandler := commonApi.NewCommonDictDataHandler()
			commonDictDataGroup.GET("/whoami", commonDictDataHandler.Whoami())
			commonDictDataGroup.PUT("/add", middleware.Limiter(), commonDictDataHandler.Add())
			commonDictDataGroup.DELETE("/remove", middleware.Limiter(), commonDictDataHandler.Remove())
			commonDictDataGroup.POST("/base-modify", middleware.Limiter(), commonDictDataHandler.Modify())
			commonDictDataGroup.GET("/detail", commonDictDataHandler.Detail())
			commonDictDataGroup.POST("/list-page", commonDictDataHandler.ListPage())
			commonDictDataGroup.POST("/list-page-condition", commonDictDataHandler.ListPageByCondition())

		}
	}

}
