// Package service
// @author tabuyos
// @since 2023/09/15
// @description my_plat_authority_resource
package service

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatAuthorityResourceSvcPool 服务池
var myPlatAuthorityResourceSvcPool = &sync.Pool{New: func() interface{} {
	return new(myPlatAuthorityResourceSvc)
}}

// MyPlatAuthorityResourceSvc 服务接口
type MyPlatAuthorityResourceSvc interface {
	iMyPlatAuthorityResourceAutoGen
}

// myPlatAuthorityResourceSvc 服务结构体
type myPlatAuthorityResourceSvc struct {
	myPlatAuthorityResourceAutoGen
}

// NewMyPlatAuthorityResourceSvc 从池中创建服务
func NewMyPlatAuthorityResourceSvc(ctx *gin.Context) (MyPlatAuthorityResourceSvc, func()) {
	svc := myPlatAuthorityResourceSvcPool.Get().(*myPlatAuthorityResourceSvc)
	svc.ctx = ctx
	rel := func() {
		svc.ctx = nil
		myPlatAuthorityResourceSvcPool.Put(svc)
	}
	return svc, rel
}
