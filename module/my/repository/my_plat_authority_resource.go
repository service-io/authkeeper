// Package repository
// @author tabuyos
// @since 2023/09/15
// @description my_plat_authority_resource
package repository

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatAuthorityResourceRtyPool 持久池
var myPlatAuthorityResourceRtyPool = &sync.Pool{New: func() interface{} {
	return new(myPlatAuthorityResourceRty)
}}

// MyPlatAuthorityResourceRty 持久层接口
type MyPlatAuthorityResourceRty interface {
	iMyPlatAuthorityResourceAutoGen
}
type myPlatAuthorityResourceRty struct {
	myPlatAuthorityResourceAutoGen
}

// NewMyPlatAuthorityResourceRty 从池中创建
func NewMyPlatAuthorityResourceRty(ctx *gin.Context) (MyPlatAuthorityResourceRty, func()) {
	rty := myPlatAuthorityResourceRtyPool.Get().(*myPlatAuthorityResourceRty)
	rty.ctx = ctx
	rel := func() {
		rty.ctx = nil
		myPlatAuthorityResourceRtyPool.Put(rty)
	}
	return rty, rel
}
