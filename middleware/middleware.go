// Package middleware
// @author tabuyos
// @since 2023/6/30
// @description middleware
package middleware

import (
	"deepsea/config/constant"
	"deepsea/helper"
	"deepsea/helper/recorderx"
	"deepsea/helper/runerror"
	"deepsea/helper/security"
	"deepsea/model/reply"
	platformService "deepsea/module/platform/service"
	systemService "deepsea/module/system/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

type Handler interface {
	Handle(ctx *gin.Context) bool
}

func CheckAuth(chain ...AuthHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ok = false
		for _, handler := range chain {
			hs := handler.Handle(ctx)
			if hs {
				ok = true
				break
			}
		}
		if !ok {
			code := runerror.GetUsrErp(runerror.ModAuth, runerror.NotLoginError)
			panic(runerror.NewAll(code, errors.New("用户未登录")))
		}
	}
}

func CheckRBAC(chain ...RbacHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recorder := recorderx.FetchRecorder(ctx)
		id := security.GetAccountID(ctx)
		recorder.Info(fmt.Sprintf("检测 ID: %d 的相关 RBAC 信息", id))

		roleSvc, roleRel := platformService.NewPlatRoleSvc(ctx)
		defer roleRel()

		permSvc, permRel := platformService.NewPlatAuthoritySvc(ctx)
		defer permRel()

		security.LookupRole(ctx, roleSvc)
		security.LookupPerm(ctx, permSvc)

		var fail = false
		for _, handler := range chain {
			hs := handler.Handle(ctx)
			if !hs {
				fail = true
				break
			}
		}
		if fail {
			code := runerror.GetUsrErp(runerror.ModAuth, runerror.NoPermissionError)
			panic(runerror.NewAll(code, errors.New("无访问许可")))
		}
	}
}

func Limiter() gin.HandlerFunc {
	var apiLimiter = NewDefaultUserModifyApiLimiter()
	return func(ctx *gin.Context) {
		limiter := apiLimiter.Get(ctx)
		if !limiter.Allow() {
			code := runerror.GetUsrErp(runerror.ModNil, runerror.TooManyRequests)
			detail := reply.Failed().WithCode(code).WithMessage("请求过于频繁, 请稍后重试")
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, detail)
			return
		}
		ctx.Next()
	}
}

func RegisterLogger() gin.HandlerFunc {
	svc, _ := systemService.NewSysLogSvc(nil)
	svc.Persistence()

	signSupplier := func(ctx *gin.Context) string {
		ip := ctx.ClientIP()
		return helper.RecorderSign(security.GetAccountIDString(ctx), ip)
	}

	return func(ctx *gin.Context) {
		recorderx.WithGinContext(ctx, signSupplier)
	}
}

func FillTrace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := uuid.New().String()
		ctx.Set(constant.TraceIdKey, traceID)
		ctx.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				recorder := recorderx.FetchRecorder(ctx).WithOptions(recorderx.AddCallerSkip(4))
				switch v := err.(type) {
				case runerror.IRunError:
					code := v.Code()
					topMessage := v.TopMessage()
					message := v.Error()
					recorder.Errorf("fetch error -> code: %d message: %s", code, message)

					res := reply.Failed()
					// 大于默认异常的直接返回
					if code >= runerror.DftError {
						res.WithCode(code).WithMessage(topMessage)
					}
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
				case string:
					recorder.Errorf("fetch error -> top message: %s", v)
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, reply.Failed())
				case error:
					recorder.Errorf("fetch error -> top message: %s", v.Error())
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, reply.Failed())
				default:
					recorder.Errorf("fetch error -> top message: %+v", v)
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, reply.Failed())
				}
				return
			}
		}()
		ctx.Next()
	}
}

func Visitor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		visitor := recorderx.FetchVisitor(ctx)
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		// hand over to the next handler
		ctx.Next()
		go func() {
			latency := time.Since(start)
			clientIP := ctx.ClientIP()
			method := ctx.Request.Method
			statusCode := ctx.Writer.Status()
			errorMessage := strings.TrimRight(ctx.Errors.ByType(gin.ErrorTypePrivate).String(), "\n")
			msg := fmt.Sprintf("%v -> %v(%v)[%v] -> %v -> |%v|%v", clientIP, path, raw, method, latency, statusCode, errorMessage)
			visitor.Info(msg)
		}()
	}
}
