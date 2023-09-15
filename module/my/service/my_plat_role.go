// Package service
// @author tabuyos
// @since 2023/09/15
// @description my_plat_role
package service

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatRoleSvcPool 服务池
var myPlatRoleSvcPool = &sync.Pool{New: func() interface{} {
	return new(myPlatRoleSvc)
}}

// MyPlatRoleSvc 服务接口
type MyPlatRoleSvc interface {
	iMyPlatRoleAutoGen
}

// myPlatRoleSvc 服务结构体
type myPlatRoleSvc struct {
	myPlatRoleAutoGen
}

// NewMyPlatRoleSvc 从池中创建服务
func NewMyPlatRoleSvc(ctx *gin.Context) (MyPlatRoleSvc, func()) {
	svc := myPlatRoleSvcPool.Get().(*myPlatRoleSvc)
	svc.ctx = ctx
	rel := func() {
		svc.ctx = nil
		myPlatRoleSvcPool.Put(svc)
	}
	return svc, rel
}
