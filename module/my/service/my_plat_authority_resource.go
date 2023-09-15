// Package service
// @author tabuyos
// @since 2023/09/15
// @description my_plat_authority_resource
package service

import (
	"deepsea/helper/recorderx"
	"deepsea/model/dto"
	"deepsea/module/my/repository"
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
	FindAuthorityByResourceID(id int64) []*dto.MyPlatAuthority
	FindResourceByAuthorityID(id int64) []*dto.MyPlatResource
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

func (svc *myPlatAuthorityResourceSvc) FindAuthorityByResourceID(id int64) []*dto.MyPlatAuthority {
	recorder := recorderx.FetchRecorder(svc.ctx)
	recorder.Infof("查询 Resource ID: %+v 的数据", id)
	rty, release := repository.NewMyPlatAuthorityResourceRty(svc.ctx)
	defer release()
	ets := rty.SelectAuthorityByResourceID(id)
	if ets == nil {
		return nil
	}
	values := make([]*dto.MyPlatAuthority, len(ets))
	for i, eto := range ets {
		value := dto.NewMyPlatAuthority()
		value.From(eto)
		values[i] = value
	}
	return values
}

func (svc *myPlatAuthorityResourceSvc) FindResourceByAuthorityID(id int64) []*dto.MyPlatResource {
	recorder := recorderx.FetchRecorder(svc.ctx)
	recorder.Infof("查询 Resource ID: %+v 的数据", id)
	rty, release := repository.NewMyPlatAuthorityResourceRty(svc.ctx)
	defer release()
	ets := rty.SelectResourceByAuthorityID(id)
	if ets == nil {
		return nil
	}
	values := make([]*dto.MyPlatResource, len(ets))
	for i, eto := range ets {
		value := dto.NewMyPlatResource()
		value.From(eto)
		values[i] = value
	}
	return values
}
