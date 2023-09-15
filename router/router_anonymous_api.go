// Package router
// @author tabuyos
// @since 2023/8/11
// @description router
package router

import (
	platformApi "deepsea/module/platform/api"
	"github.com/gin-gonic/gin"
)

func setAnonymousApi(version *gin.RouterGroup) {
	platformModGroup := version.Group("/platform")
	{
		platAccountGroup := platformModGroup.Group("/plat-account")
		{
			platAccountHandler := platformApi.NewPlatAccountHandler()
			platAccountGroup.POST("/register", platAccountHandler.Register())
			platAccountGroup.POST("/login", platAccountHandler.Login())
			platAccountGroup.POST("/refresh", platAccountHandler.Refresh())
		}
	}
}
