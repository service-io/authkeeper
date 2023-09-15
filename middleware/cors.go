// Package middleware
// @author tabuyos
// @since 2023/8/4
// @description middleware
package middleware

import (
	"deepsea/config/constant"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Host", "Referer", "Origin", "Connection", "Content-Length", "Content-Type", constant.HeaderLoginToken, constant.HeaderSignToken},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}
