// Package repository
// @author tabuyos
// @since 2023/09/15
// @description my_plat_role
package repository

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatRoleRtyPool 持久池
var myPlatRoleRtyPool = &sync.Pool{New: func() interface{} {
	return new(myPlatRoleRty)
}}

// MyPlatRoleRty 持久层接口
type MyPlatRoleRty interface {
	iMyPlatRoleAutoGen
}
type myPlatRoleRty struct {
	myPlatRoleAutoGen
}

// NewMyPlatRoleRty 从池中创建
func NewMyPlatRoleRty(ctx *gin.Context) (MyPlatRoleRty, func()) {
	rty := myPlatRoleRtyPool.Get().(*myPlatRoleRty)
	rty.ctx = ctx
	rel := func() {
		rty.ctx = nil
		myPlatRoleRtyPool.Put(rty)
	}
	return rty, rel
}
