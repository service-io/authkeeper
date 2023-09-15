// Package router
// @author tabuyos
// @since 2023/8/11
// @description router
package router

import (
	"deepsea/middleware"
	tenantApi "deepsea/module/tenant/api"
	"github.com/gin-gonic/gin"
)

func setTenantApi(version *gin.RouterGroup) {
	tenantModGroup := version.Group("/tenant")
	{
		tenantGroup := tenantModGroup.Group("/tenant")
		{
			tenantHandler := tenantApi.NewTenantHandler()
			tenantGroup.GET("/whoami", tenantHandler.Whoami())
			tenantGroup.PUT("/add", middleware.Limiter(), tenantHandler.Add())
			tenantGroup.DELETE("/remove", middleware.Limiter(), tenantHandler.Remove())
			tenantGroup.POST("/base-modify", middleware.Limiter(), tenantHandler.Modify())
			tenantGroup.GET("/detail", tenantHandler.Detail())
			tenantGroup.POST("/list-page", tenantHandler.ListPage())
			tenantGroup.POST("/list-page-condition", tenantHandler.ListPageByCondition())
		}
		tenantDeptGroup := tenantModGroup.Group("/tenant-dept")
		{
			tenantDeptHandler := tenantApi.NewTenantDeptHandler()
			tenantDeptGroup.GET("/whoami", tenantDeptHandler.Whoami())
			tenantDeptGroup.PUT("/add", middleware.Limiter(), tenantDeptHandler.Add())
			tenantDeptGroup.DELETE("/remove", middleware.Limiter(), tenantDeptHandler.Remove())
			tenantDeptGroup.POST("/base-modify", middleware.Limiter(), tenantDeptHandler.Modify())
			tenantDeptGroup.GET("/detail", tenantDeptHandler.Detail())
			tenantDeptGroup.POST("/list-page", tenantDeptHandler.ListPage())
			tenantDeptGroup.POST("/list-page-condition", tenantDeptHandler.ListPageByCondition())
			tenantDeptGroup.GET("/list", tenantDeptHandler.FindList())

		}
		tenantPostGroup := tenantModGroup.Group("/tenant-post")
		{
			tenantPostHandler := tenantApi.NewTenantPostHandler()
			tenantPostGroup.GET("/whoami", tenantPostHandler.Whoami())
			tenantPostGroup.PUT("/add", middleware.Limiter(), tenantPostHandler.Add())
			tenantPostGroup.DELETE("/remove", middleware.Limiter(), tenantPostHandler.Remove())
			tenantPostGroup.POST("/base-modify", middleware.Limiter(), tenantPostHandler.Modify())
			tenantPostGroup.GET("/detail", tenantPostHandler.Detail())
			tenantPostGroup.POST("/list-page", tenantPostHandler.ListPage())
			tenantPostGroup.POST("/list-page-condition", tenantPostHandler.ListPageByCondition())
			tenantPostGroup.GET("/list", tenantPostHandler.FindList())

		}
	}
}
