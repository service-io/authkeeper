// Package service
// @author tabuyos
// @since 2023/09/15
// @description my_plat_resource
package service

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatResourceSvcPool 服务池
var myPlatResourceSvcPool = &sync.Pool{New: func() interface{} {
	return new(myPlatResourceSvc)
}}

// MyPlatResourceSvc 服务接口
type MyPlatResourceSvc interface {
	iMyPlatResourceAutoGen
}

// myPlatResourceSvc 服务结构体
type myPlatResourceSvc struct {
	myPlatResourceAutoGen
}

// NewMyPlatResourceSvc 从池中创建服务
func NewMyPlatResourceSvc(ctx *gin.Context) (MyPlatResourceSvc, func()) {
	svc := myPlatResourceSvcPool.Get().(*myPlatResourceSvc)
	svc.ctx = ctx
	rel := func() {
		svc.ctx = nil
		myPlatResourceSvcPool.Put(svc)
	}
	return svc, rel
}
