// Package recorderx
// @author tabuyos
// @since 2023/8/24
// @description recorderx
package recorderx

import "github.com/gin-gonic/gin"

type PersistService interface {
	Persistence()
}

type SignService func(ctx *gin.Context) string

type deliver struct {
	ginCtx       *gin.Context
	signSupplier SignService
}
