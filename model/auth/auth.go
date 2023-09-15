// Package auth
// @author tabuyos
// @since 2023/8/5
// @description auth
package auth

import (
	"deepsea/config/constant"
	"deepsea/helper/runerror"
	"errors"
	"strconv"
	"strings"
	"time"
)

type CertDetail struct {
	AccountID  *int64
	TenantID   *int64
	Name       *string
	Extra      *string
	Token      *string
	SuperAdmin *bool
	Anonymous  *bool
}

type TokenDetail struct {
	// 访问令牌
	AccessToken *string `json:"accessToken"`
	// 刷新令牌
	RefreshToken *string `json:"refreshToken"`
	// 登录名称
	Name *string `json:"name"`
	// 过期时间
	ExpireTime *time.Time `json:"expireTime"`
}

func New(aid, tid *int64) *CertDetail {
	return &CertDetail{
		AccountID: aid,
		TenantID:  tid,
	}
}

// Parse 解析认证信息
func Parse(info string) *CertDetail {
	fields := strings.Split(info, constant.ValDelimiter)
	lg := len(fields)
	if lg == 0 {
		code := runerror.GetSysErp(runerror.ModAuth, runerror.ParseError)
		panic(runerror.New().WithCode(code).WithMessage("无法解析用户信息(无信息)"))
	}

	if lg == 1 {
		id := parseID(fields[0])
		return &CertDetail{AccountID: id}
	}

	if lg == 2 {
		aid := parseID(fields[0])
		tid := parseID(fields[1])
		return &CertDetail{
			AccountID: aid,
			TenantID:  tid,
		}
	}

	if lg == 3 {
		aid := parseID(fields[0])
		tid := parseID(fields[1])
		name := fields[2]
		return &CertDetail{
			AccountID: aid,
			TenantID:  tid,
			Name:      &name,
		}
	}

	code := runerror.GetSysErp(runerror.ModAuth, runerror.LenError)
	panic(runerror.NewAll(code, errors.Join(errors.New("无法解析用户信息(存在多个信息)"))))
}

func parseID(s string) *int64 {
	id, err := strconv.ParseInt(s, 10, 64)
	code := runerror.GetSysErp(runerror.ModAuth, runerror.ParseError)
	if err != nil {
		panic(runerror.NewAll(code, errors.Join(err, errors.New("无法解析用户信息(解析 ID 失败"))))
	}
	return &id
}
