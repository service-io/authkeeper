// Package router
// @author tabuyos
// @since 2023/8/11
// @description router
package router

import (
	systemApi "deepsea/module/system/api"
	"github.com/gin-gonic/gin"
)

func setSystemApi(version *gin.RouterGroup) {
	systemModGroup := version.Group("/system")
	{
		sysLogGroup := systemModGroup.Group("/sys-log")
		{
			sysLogHandler := systemApi.NewSysLogHandler()
			sysLogGroup.GET("/detail", sysLogHandler.Detail())
			sysLogGroup.POST("/list-page", sysLogHandler.ListPage())
			sysLogGroup.POST("/list-page-condition", sysLogHandler.ListPageByCondition())
		}
		sysOSSGroup := systemModGroup.Group("/sys-oss")
		{
			sysOSSHandler := systemApi.NewSysOssHandler()
			sysOSSGroup.POST("/upload", sysOSSHandler.Upload())
			sysOSSGroup.POST("/direct-upload", sysOSSHandler.DirectUpload())
			sysOSSGroup.GET("/sign-url/*key", sysOSSHandler.SignURLForKey())
			sysOSSGroup.GET("/get/*key", sysOSSHandler.GetForKey())
		}
	}
}
