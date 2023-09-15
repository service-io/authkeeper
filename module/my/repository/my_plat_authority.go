// Package repository
// @author tabuyos
// @since 2023/09/15
// @description my_plat_authority
package repository

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatAuthorityRtyPool 持久池
var myPlatAuthorityRtyPool = &sync.Pool{New: func() interface{} {
	return new(myPlatAuthorityRty)
}}

// MyPlatAuthorityRty 持久层接口
type MyPlatAuthorityRty interface {
	iMyPlatAuthorityAutoGen
}
type myPlatAuthorityRty struct {
	myPlatAuthorityAutoGen
}

// NewMyPlatAuthorityRty 从池中创建
func NewMyPlatAuthorityRty(ctx *gin.Context) (MyPlatAuthorityRty, func()) {
	rty := myPlatAuthorityRtyPool.Get().(*myPlatAuthorityRty)
	rty.ctx = ctx
	rel := func() {
		rty.ctx = nil
		myPlatAuthorityRtyPool.Put(rty)
	}
	return rty, rel
}
