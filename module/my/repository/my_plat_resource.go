// Package repository
// @author tabuyos
// @since 2023/09/15
// @description my_plat_resource
package repository

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// myPlatResourceRtyPool 持久池
var myPlatResourceRtyPool = &sync.Pool{New: func() interface{} {
	return new(myPlatResourceRty)
}}

// MyPlatResourceRty 持久层接口
type MyPlatResourceRty interface {
	iMyPlatResourceAutoGen
}
type myPlatResourceRty struct {
	myPlatResourceAutoGen
}

// NewMyPlatResourceRty 从池中创建
func NewMyPlatResourceRty(ctx *gin.Context) (MyPlatResourceRty, func()) {
	rty := myPlatResourceRtyPool.Get().(*myPlatResourceRty)
	rty.ctx = ctx
	rel := func() {
		rty.ctx = nil
		myPlatResourceRtyPool.Put(rty)
	}
	return rty, rel
}
