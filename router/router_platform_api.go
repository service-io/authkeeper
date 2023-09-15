// Package router
// @author tabuyos
// @since 2023/8/11
// @description router
package router

import (
	"deepsea/middleware"
	platformApi "deepsea/module/platform/api"
	"github.com/gin-gonic/gin"
)

func setPlatformApi(version *gin.RouterGroup) {
	platformModGroup := version.Group("/platform")
	{
		platGroup := platformModGroup.Group("/plat")
		{
			platHandler := platformApi.NewPlatHandler()
			platGroup.GET("/whoami", platHandler.Whoami())
			platGroup.PUT("/add", middleware.Limiter(), platHandler.Add())
			platGroup.DELETE("/remove", middleware.Limiter(), platHandler.Remove())
			platGroup.POST("/base-modify", middleware.Limiter(), platHandler.Modify())
			platGroup.GET("/detail", platHandler.Detail())
			platGroup.POST("/list-page", platHandler.ListPage())
			platGroup.POST("/list-page-condition", platHandler.ListPageByCondition())
			platGroup.GET("/list", platHandler.FindList())
		}
		platAccountGroup := platformModGroup.Group("/plat-account")
		{
			platAccountHandler := platformApi.NewPlatAccountHandler()
			platAccountGroup.PUT("/add", middleware.Limiter(), platAccountHandler.Add())
			platAccountGroup.DELETE("/remove", middleware.Limiter(), platAccountHandler.Remove())
			platAccountGroup.POST("/base-modify", middleware.Limiter(), platAccountHandler.Modify())
			platAccountGroup.GET("/detail", platAccountHandler.Detail())
			platAccountGroup.POST("/list-page", platAccountHandler.ListPage())
			platAccountGroup.POST("/list-page-condition", platAccountHandler.ListPageByCondition())
			platAccountGroup.POST("/logout", platAccountHandler.Logout())
			platAccountGroup.GET("/profile", platAccountHandler.Profile())
			platAccountGroup.POST("/refresh-acl", platAccountHandler.RefreshACL())
			platAccountGroup.GET("/owner", platAccountHandler.FindOwnerResource())
		}
		platAccountRoleGroup := platformModGroup.Group("/plat-account-role")
		{
			platAccountRoleHandler := platformApi.NewPlatAccountRoleHandler()
			platAccountRoleGroup.GET("/whoami", platAccountRoleHandler.Whoami())
			platAccountRoleGroup.PUT("/add", middleware.Limiter(), platAccountRoleHandler.Add())
			platAccountRoleGroup.DELETE("/remove", middleware.Limiter(), platAccountRoleHandler.Remove())
			platAccountRoleGroup.POST("/base-modify", middleware.Limiter(), platAccountRoleHandler.Modify())
			platAccountRoleGroup.GET("/detail", platAccountRoleHandler.Detail())
			platAccountRoleGroup.POST("/list-page", platAccountRoleHandler.ListPage())
			platAccountRoleGroup.POST("/list-page-condition", platAccountRoleHandler.ListPageByCondition())
			platAccountRoleGroup.PUT("/batch-add", middleware.Limiter(), platAccountRoleHandler.BatchAdd())
		}
		platAccountTokenGroup := platformModGroup.Group("/plat-account-token")
		{
			platAccountTokenHandler := platformApi.NewPlatAccountTokenHandler()
			platAccountTokenGroup.GET("/whoami", platAccountTokenHandler.Whoami())
			platAccountTokenGroup.PUT("/add", middleware.Limiter(), platAccountTokenHandler.Add())
			platAccountTokenGroup.DELETE("/remove", middleware.Limiter(), platAccountTokenHandler.Remove())
			platAccountTokenGroup.POST("/base-modify", middleware.Limiter(), platAccountTokenHandler.Modify())
			platAccountTokenGroup.GET("/detail", platAccountTokenHandler.Detail())
			platAccountTokenGroup.POST("/list-page", platAccountTokenHandler.ListPage())
			platAccountTokenGroup.POST("/list-page-condition", platAccountTokenHandler.ListPageByCondition())
		}
		platAuthorityGroup := platformModGroup.Group("/plat-authority")
		{
			platAuthorityHandler := platformApi.NewPlatAuthorityHandler()
			platAuthorityGroup.GET("/whoami", platAuthorityHandler.Whoami())
			platAuthorityGroup.PUT("/add", middleware.Limiter(), platAuthorityHandler.Add())
			platAuthorityGroup.DELETE("/remove", middleware.Limiter(), platAuthorityHandler.Remove())
			platAuthorityGroup.POST("/base-modify", middleware.Limiter(), platAuthorityHandler.Modify())
			platAuthorityGroup.GET("/detail", platAuthorityHandler.Detail())
			platAuthorityGroup.POST("/list-page", platAuthorityHandler.ListPage())
			platAuthorityGroup.POST("/list-page-condition", platAuthorityHandler.ListPageByCondition())
			platAuthorityGroup.GET("/list", platAuthorityHandler.FindList())
		}
		platAuthorityResourceGroup := platformModGroup.Group("/plat-authority-resource")
		{
			platAuthorityResourceHandler := platformApi.NewPlatAuthorityResourceHandler()
			platAuthorityResourceGroup.GET("/whoami", platAuthorityResourceHandler.Whoami())
			platAuthorityResourceGroup.PUT("/add", middleware.Limiter(), platAuthorityResourceHandler.Add())
			platAuthorityResourceGroup.DELETE("/remove", middleware.Limiter(), platAuthorityResourceHandler.Remove())
			platAuthorityResourceGroup.POST("/base-modify", middleware.Limiter(), platAuthorityResourceHandler.Modify())
			platAuthorityResourceGroup.GET("/detail", platAuthorityResourceHandler.Detail())
			platAuthorityResourceGroup.POST("/list-page", platAuthorityResourceHandler.ListPage())
			platAuthorityResourceGroup.POST("/list-page-condition", platAuthorityResourceHandler.ListPageByCondition())
			platAuthorityResourceGroup.PUT("/batch-add", middleware.Limiter(), platAuthorityResourceHandler.BatchAdd())
		}
		platResourceGroup := platformModGroup.Group("/plat-resource")
		{
			platResourceHandler := platformApi.NewPlatResourceHandler()
			platResourceGroup.GET("/whoami", platResourceHandler.Whoami())
			platResourceGroup.PUT("/add", middleware.Limiter(), platResourceHandler.Add())
			platResourceGroup.DELETE("/remove", middleware.Limiter(), platResourceHandler.Remove())
			platResourceGroup.POST("/base-modify", middleware.Limiter(), platResourceHandler.Modify())
			platResourceGroup.GET("/detail", platResourceHandler.Detail())
			platResourceGroup.POST("/list-page", platResourceHandler.ListPage())
			platResourceGroup.POST("/list-page-condition", platResourceHandler.ListPageByCondition())
			platResourceGroup.GET("/plat", platResourceHandler.FindResource())
			platResourceGroup.GET("/list", platResourceHandler.FindList())
		}
		platRoleGroup := platformModGroup.Group("/plat-role")
		{
			platRoleHandler := platformApi.NewPlatRoleHandler()
			platRoleGroup.GET("/whoami", platRoleHandler.Whoami())
			platRoleGroup.PUT("/add", middleware.Limiter(), platRoleHandler.Add())
			platRoleGroup.DELETE("/remove", middleware.Limiter(), platRoleHandler.Remove())
			platRoleGroup.POST("/base-modify", middleware.Limiter(), platRoleHandler.Modify())
			platRoleGroup.GET("/detail", platRoleHandler.Detail())
			platRoleGroup.POST("/list-page", platRoleHandler.ListPage())
			platRoleGroup.POST("/list-page-condition", platRoleHandler.ListPageByCondition())
			platRoleGroup.GET("/list", platRoleHandler.FindList())
			platRoleGroup.GET("/tree", platRoleHandler.FindTree())
		}
		platRoleAuthorityGroup := platformModGroup.Group("/plat-role-authority")
		{
			platRoleAuthorityHandler := platformApi.NewPlatRoleAuthorityHandler()
			platRoleAuthorityGroup.GET("/whoami", platRoleAuthorityHandler.Whoami())
			platRoleAuthorityGroup.PUT("/add", middleware.Limiter(), platRoleAuthorityHandler.Add())
			platRoleAuthorityGroup.DELETE("/remove", middleware.Limiter(), platRoleAuthorityHandler.Remove())
			platRoleAuthorityGroup.POST("/base-modify", middleware.Limiter(), platRoleAuthorityHandler.Modify())
			platRoleAuthorityGroup.GET("/detail", platRoleAuthorityHandler.Detail())
			platRoleAuthorityGroup.POST("/list-page", platRoleAuthorityHandler.ListPage())
			platRoleAuthorityGroup.POST("/list-page-condition", platRoleAuthorityHandler.ListPageByCondition())
			platRoleAuthorityGroup.PUT("/batch-add", middleware.Limiter(), platRoleAuthorityHandler.BatchAdd())
		}
		platTenantGroup := platformModGroup.Group("/plat-tenant")
		{
			platTenantHandler := platformApi.NewPlatTenantHandler()
			platTenantGroup.GET("/whoami", platTenantHandler.Whoami())
			platTenantGroup.PUT("/add", middleware.Limiter(), platTenantHandler.Add())
			platTenantGroup.DELETE("/remove", middleware.Limiter(), platTenantHandler.Remove())
			platTenantGroup.POST("/base-modify", middleware.Limiter(), platTenantHandler.Modify())
			platTenantGroup.GET("/detail", platTenantHandler.Detail())
			platTenantGroup.POST("/list-page", platTenantHandler.ListPage())
			platTenantGroup.POST("/list-page-condition", platTenantHandler.ListPageByCondition())
			platTenantGroup.PUT("/batch-add", middleware.Limiter(), platTenantHandler.BatchAdd())
		}
	}
}
