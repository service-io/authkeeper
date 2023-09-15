// Code generated by tabuyos. DO NOT EDIT.

// Package api
// @author tabuyos
// @since 2023/09/15
// @description my_plat_resource
package api

import (
	"deepsea/helper/recorderx"
	"deepsea/model/dto"
	"deepsea/model/page"
	"deepsea/model/reply"
	"deepsea/module/my/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// MyPlatResourceHandler API 处理器
type MyPlatResourceHandler struct{}

// NewMyPlatResourceHandler 创建 API 处理器
func NewMyPlatResourceHandler() *MyPlatResourceHandler {
	return &MyPlatResourceHandler{}
}

// Add 新增数据
// @Summary      新增数据
// @Description  新增数据
// @Tags         my,MyPlatResource
// @Accept       json
// @Produce      json
// @Param        req_info    body     dto.MyPlatResource  true  "待新增的数据对象"
// @Success      200  {object}  reply.Reply  "操作结果"
// @Security     ApiKeyAuth
// @Router       /my/resource/add [put]
func (*MyPlatResourceHandler) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recorder := recorderx.FetchRecorder(ctx)
		newMyPlatResource := dto.NewMyPlatResource()
		err := ctx.ShouldBindJSON(newMyPlatResource)
		recorder.MaybePanic(err)
		svc, release := service.NewMyPlatResourceSvc(ctx)
		defer release()
		id := svc.Add(newMyPlatResource)
		if id != 0 {
			ctx.JSON(http.StatusOK, reply.OkPayload(id))
		} else {
			ctx.JSON(http.StatusOK, reply.FailedMessage("新增失败"))
		}
	}
}

// Remove 根据 ID 删除数据
// @Summary      删除数据
// @Description  根据 ID 删除数据
// @Tags         my,MyPlatResource
// @Accept       json
// @Produce      json
// @Param        id    query     integer  true  "待删除 ID"
// @Success      200  {object}  reply.Reply  "操作结果"
// @Security     ApiKeyAuth
// @Router       /my/resource/remove [delete]
func (*MyPlatResourceHandler) Remove() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recorder := recorderx.FetchRecorder(ctx)
		id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
		recorder.MaybePanic(err)
		if id <= 0 {
			ctx.JSON(http.StatusOK, reply.FailedMessage("请传递正确的 ID"))
			return
		}
		svc, release := service.NewMyPlatResourceSvc(ctx)
		defer release()
		op := svc.Remove(id)
		if op {
			ctx.JSON(http.StatusOK, reply.Ok().WithState(op).WithMessage("操作成功"))
			return
		}
		ctx.JSON(http.StatusOK, reply.Failed().WithState(op).WithMessage("操作失败"))
	}
}

// Modify 根据 ID 修改数据
// @Summary      修改数据
// @Description  根据 ID 修改数据
// @Tags         my,MyPlatResource
// @Accept       json
// @Produce      json
// @Param        req_info    body     dto.MyPlatResource  true  "待修改的数据对象"
// @Success      200  {object}  reply.Reply  "操作结果"
// @Security     ApiKeyAuth
// @Router       /my/resource/base-modify [post]
func (*MyPlatResourceHandler) Modify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recorder := recorderx.FetchRecorder(ctx)
		newMyPlatResource := dto.NewMyPlatResource()
		err := ctx.ShouldBindJSON(newMyPlatResource)
		recorder.MaybePanic(err)
		if newMyPlatResource.ID == nil {
			ctx.JSON(http.StatusOK, reply.FailedMessage("请传递正确的 ID, 以及需要修改的信息"))
			return
		}
		svc, release := service.NewMyPlatResourceSvc(ctx)
		defer release()
		op := svc.Modify(newMyPlatResource)
		if op {
			ctx.JSON(http.StatusOK, reply.OkMessage("操作成功"))
			return
		}
		ctx.JSON(http.StatusOK, reply.FailedMessage("操作失败"))
	}
}

// Detail 根据 ID 获取详情
// @Summary      获取详情
// @Description  根据 ID 获取详情
// @Tags         my,MyPlatResource
// @Accept       json
// @Produce      json
// @Param        id    query     integer  true  "查询 ID"
// @Success      200  {object}  reply.Reply{payload=dto.MyPlatResource}  "查询详情"
// @Security     ApiKeyAuth
// @Router       /my/resource/detail [get]
func (*MyPlatResourceHandler) Detail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recorder := recorderx.FetchRecorder(ctx)
		id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
		recorder.MaybePanic(err)
		if id <= 0 {
			ctx.JSON(http.StatusOK, reply.FailedMessage("请传递正确的 ID"))
			return
		}
		svc, release := service.NewMyPlatResourceSvc(ctx)
		defer release()
		myPlatResource := svc.Find(id)
		if myPlatResource != nil {
			ctx.JSON(http.StatusOK, reply.OkPayload(myPlatResource))
			return
		}
		ctx.JSON(http.StatusOK, reply.Failed().WithMessage("无对应数据"))
	}
}

// ListPage 分页列表
// @Summary      分页列表
// @Description  获取分页列表
// @Tags         my,MyPlatResource
// @Accept       json
// @Produce      json
// @Param        req_info    body     page.Query  false  "分页信息"
// @Success      200  {object}  reply.Reply{payload=[]page.Result}  "分页列表"
// @Security     ApiKeyAuth
// @Router       /my/resource/list-page [post]
func (*MyPlatResourceHandler) ListPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recorder := recorderx.FetchRecorder(ctx)
		query := page.NewQuery()
		query.Page = 1
		query.Size = 20
		err := ctx.ShouldBindJSON(query)
		recorder.MaybePanic(err)
		svc, release := service.NewMyPlatResourceSvc(ctx)
		defer release()
		ctx.JSON(http.StatusOK, reply.OkPayload(svc.FindWithPage(*query)))
	}
}
