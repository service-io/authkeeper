// Package middleware
// @author tabuyos
// @since 2023/8/5
// @description middleware
package middleware

import (
	"deepsea/config/constant"
	"deepsea/helper/recorderx"
	"deepsea/helper/security"
	"deepsea/model/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler Handler

type cert struct {
}

type anonymousCert cert

type tokenCert cert

type signatureCert cert

func NewTokenAuth() AuthHandler {
	return &tokenCert{}
}

func NewAnonymousAuth() AuthHandler {
	return &anonymousCert{}
}

func NewSignatureAuth() AuthHandler {
	return &signatureCert{}
}

// Handle 使用 Token 进行认证
func (rec *tokenCert) Handle(ctx *gin.Context) bool {
	recorder := recorderx.FetchRecorder(ctx)
	header := ctx.GetHeader(constant.HeaderLoginToken)
	if len(header) == 0 {
		recorder.Info("无 token 信息, 将交由下一个处理链")
		return false
	}

	info := security.GetByAccessToken(ctx, header)
	certDetail := auth.Parse(info)
	certDetail.Token = &header
	security.SetCert(ctx, certDetail)
	security.RenewToken(ctx, header)
	return true
}

// Handle 使用匿名认证
func (rec *anonymousCert) Handle(ctx *gin.Context) bool {
	recorder := recorderx.FetchRecorder(ctx)
	recorder.Info("匿名认证...")

	name := "anonymous"
	aid := int64(0)
	anonymous := true
	certDetail := &auth.CertDetail{Name: &name, AccountID: &aid, Anonymous: &anonymous}
	security.SetCert(ctx, certDetail)
	return true
}

// Handle 使用签名进行认证
func (rec *signatureCert) Handle(ctx *gin.Context) bool {
	return true
}
