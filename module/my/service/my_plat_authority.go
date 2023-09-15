// Package service
// @author tabuyos
// @since 2023/09/15
// @description my_plat_authority
package service

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatAuthoritySvcPool 服务池
var myPlatAuthoritySvcPool = &sync.Pool{New: func() interface{} {
	return new(myPlatAuthoritySvc)
}}

// MyPlatAuthoritySvc 服务接口
type MyPlatAuthoritySvc interface {
	iMyPlatAuthorityAutoGen
}

// myPlatAuthoritySvc 服务结构体
type myPlatAuthoritySvc struct {
	myPlatAuthorityAutoGen
}

// NewMyPlatAuthoritySvc 从池中创建服务
func NewMyPlatAuthoritySvc(ctx *gin.Context) (MyPlatAuthoritySvc, func()) {
	svc := myPlatAuthoritySvcPool.Get().(*myPlatAuthoritySvc)
	svc.ctx = ctx
	rel := func() {
		svc.ctx = nil
		myPlatAuthoritySvcPool.Put(svc)
	}
	return svc, rel
}
