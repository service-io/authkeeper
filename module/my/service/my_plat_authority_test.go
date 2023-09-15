// Package service
// @author tabuyos
// @since 2023/9/15
// @description service
package service

import (
	"deepsea/config"
	"deepsea/config/constant"
	"deepsea/config/env"
	"deepsea/helper"
	"deepsea/helper/database"
	"deepsea/helper/encodingx"
	"deepsea/helper/recorderx"
	"deepsea/helper/redisx"
	"deepsea/helper/snowflakeid"
	"deepsea/model/auth"
	"deepsea/model/dto"
	"deepsea/model/page"
	"deepsea/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func TestMyPlatAuthorityAutoGen_Add(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthoritySvc(ctx)
	defer rel()
	authority := dto.NewMyPlatAuthority()
	authority.Name = helper.ToPtr("测试")
	id := svc.Add(authority)

	fmt.Println(id)

	time.Sleep(time.Second)
}

func TestMyPlatAuthorityAutoGen_Modify(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthoritySvc(ctx)
	defer rel()
	authority := dto.NewMyPlatAuthority()
	authority.ID = helper.ToPtr[int64](1702559289006100480)
	authority.Name = helper.ToPtr("测试123")
	id := svc.Modify(authority)

	fmt.Println(id)

	time.Sleep(time.Second)
}

func TestMyPlatAuthorityAutoGen_Query(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthoritySvc(ctx)
	defer rel()
	authority := dto.NewMyPlatAuthority()
	authority.ID = helper.ToPtr[int64](1702559289006100480)
	authority.Name = helper.ToPtr("测试123")
	id := svc.Find(1702559289006100480)

	fmt.Println(id)
	fmt.Println(*id.Name)

	time.Sleep(time.Second)
}

func TestMyPlatAuthorityAutoGen_QueryPage(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthoritySvc(ctx)
	defer rel()
	query := page.NewQuery()
	query.Page = 1
	query.Size = 10
	id := svc.FindWithPage(*query)

	fmt.Println(id)

	time.Sleep(time.Second)
}

func TestMyPlatAuthorityAutoGen_Delete(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthoritySvc(ctx)
	defer rel()
	query := page.NewQuery()
	query.Page = 1
	query.Size = 10
	id := svc.Remove(1702559289006100480)

	fmt.Println(id)

	time.Sleep(time.Second)
}

func TestMyPlatAuthorityAutoGen_ShipAdd(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthorityResourceSvc(ctx)
	// svc, rel := NewMyPlatResourceSvc(ctx)
	defer rel()
	ship := dto.NewMyPlatAuthorityResource()
	// ship := dto.NewMyPlatResource()
	// ship.Name = helper.ToPtr("你好")
	// 1702560316145012736
	ship.AuthorityId = helper.ToPtr[int64](1702559289006100480)
	ship.ResourceId = helper.ToPtr[int64](1702569573737304064)
	id := svc.Add(ship)

	fmt.Println(id)

	time.Sleep(time.Second)
}

func TestMyPlatAuthorityAutoGen_ShipQueryLeft(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthorityResourceSvc(ctx)
	defer rel()
	id := svc.FindAuthorityByResourceID(1702569573737304064)
	fmt.Println(*id[0].Name)
	time.Sleep(time.Second)
}

func TestMyPlatAuthorityAutoGen_ShipQueryRight(t *testing.T) {
	// 初始化编码环境
	encodingx.InitEncodingX()
	// 初始化环境
	env.SpecialEnv("dev")
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	recorderx.InitRecorder()
	// 初始化数据库(单机)
	database.InitDB()
	// 初始化 Redis
	redisx.InitRedisX()
	// 初始化路由
	router.InitRouter()
	// 初始化雪花算法
	snowflakeid.InitSnowflake()

	ctx := &gin.Context{}
	var aid int64 = 111000
	var tid int64 = 111001
	ctx.Set(constant.CtxCertKey, auth.New(&aid, &tid))
	svc, rel := NewMyPlatAuthorityResourceSvc(ctx)
	defer rel()
	id := svc.FindResourceByAuthorityID(1702559289006100480)
	fmt.Println(*id[0].Name)
	time.Sleep(time.Second)
}
