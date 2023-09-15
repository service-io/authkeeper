// Package security
// @author tabuyos
// @since 2023/8/7
// @description security
package security

import (
	"deepsea/config"
	"deepsea/config/constant"
	"deepsea/helper"
	"deepsea/helper/recorderx"
	"deepsea/helper/redisx"
	"deepsea/helper/runerror"
	"deepsea/helper/sonar"
	"deepsea/model/auth"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

// GenerateHash 生成 hash
func GenerateHash(pwd string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hash
}

// GeneratePasswordFromByte 从hash生成加密
func GeneratePasswordFromByte(hash []byte) string {
	return string(hash)
}

// GeneratePassword 直接从字符串生成加密
func GeneratePassword(pwd string) string {
	return string(GenerateHash(pwd))
}

// ComparePassword 比较密码
func ComparePassword(storePwd, userPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storePwd), []byte(userPwd))
	if err != nil {
		return false
	}
	return true
}

// GenerateToken 生成令牌
func GenerateToken(hash []byte) string {
	tokenHash := make([]byte, len(hash), len(hash))
	copy(tokenHash, hash)

	helper.Shuffle(tokenHash)

	nanosecond := time.Now().UnixNano()
	for i := 0; i < 8; i++ {
		tokenHash[i] = uint8(nanosecond >> (4 * (7 - i)))
	}

	token := base64.StdEncoding.EncodeToString(tokenHash)

	return token
}

// GenerateDefaultToken 生成默认令牌
func GenerateDefaultToken() string {
	return GenerateCustomToken(64)
}

// GenerateCustomToken 生成自定义令牌
func GenerateCustomToken(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[r.Intn(len(runes))]
	}
	return string(b)
}

// SetAccessToken 设置访问令牌
func SetAccessToken(ctx *gin.Context, token, value string, duration time.Duration) bool {
	return SetToken(ctx, constant.AccessTokenPrefix, token, value, duration)
}

// SetRefreshToken 设置刷新令牌
func SetRefreshToken(ctx *gin.Context, token, value string, duration time.Duration) bool {
	return SetToken(ctx, constant.RefreshTokenPrefix, token, value, duration)
}

// GetByAccessToken 获取访问令牌
func GetByAccessToken(ctx *gin.Context, token string) string {
	return GetByToken(ctx, constant.AccessTokenPrefix, token)
}

// GetByRefreshToken 获取刷新lp
func GetByRefreshToken(ctx *gin.Context, token string) string {
	return GetByToken(ctx, constant.RefreshTokenPrefix, token)
}

// SetToken 设置任意令牌
func SetToken(ctx *gin.Context, prefix, token, value string, duration time.Duration) bool {
	if len(token) == 0 || len(value) == 0 {
		return false
	}
	recorder := recorderx.FetchRecorder(ctx)
	emitter, release := redisx.NewRedisEmitter(recorder)
	defer release()
	ok := emitter.SetNX(emitter.BuildKey(prefix, token), value, duration)
	return ok
}

func GenTokenInfo(name string, hash []byte) *auth.TokenDetail {
	securityConfig := config.TomlConfig().Security
	accessToken := GenerateToken(hash)
	refreshToken := GenerateToken(hash)
	expireTime := time.Now().Add(securityConfig.AccessTTL)

	tokenCert := &auth.TokenDetail{
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
		Name:         &name,
		ExpireTime:   &expireTime,
	}
	return tokenCert
}

func StoreToken(ctx *gin.Context, info *auth.TokenDetail, aid, tid int64) bool {
	securityConfig := config.TomlConfig().Security
	recorder := recorderx.FetchRecorder(ctx)
	ok := SetAccessToken(ctx, *info.AccessToken, redisx.BuildVal(strconv.FormatInt(aid, 10), strconv.FormatInt(tid, 10)), securityConfig.AccessTTL)
	if !ok {
		recorder.Error("设置访问令牌失败")
		code := runerror.GetSysErp(runerror.ModAuth, runerror.TokenStoreError)
		panic(runerror.NewAll(code, runerror.NewError("设置令牌失败")))
	}
	ok = SetRefreshToken(ctx, *info.RefreshToken, redisx.BuildVal(strconv.FormatInt(aid, 10), strconv.FormatInt(tid, 10), *info.Name), securityConfig.RefreshTTL)
	if !ok {
		recorder.Error("设置访问令牌失败")
		code := runerror.GetSysErp(runerror.ModAuth, runerror.TokenStoreError)
		panic(runerror.NewAll(code, runerror.NewError("设置令牌失败")))
	}
	return true
}

func DiscardToken(ctx *gin.Context, accessToken, refreshToken string) bool {
	recorder := recorderx.FetchRecorder(ctx)
	emitter, release := redisx.NewRedisEmitter(recorder)
	defer release()

	var keys []string
	if accessToken != "" {
		keys = append(keys, emitter.BuildKey(constant.AccessTokenPrefix, accessToken))
	}
	if refreshToken != "" {
		keys = append(keys, emitter.BuildKey(constant.RefreshTokenPrefix, refreshToken))
	}

	return emitter.Del(keys...)
}

func RenewToken(ctx *gin.Context, token string) {
	go func() {
		cert := GetCert(ctx)
		if len(token) == 0 {
			return
		}
		if cert == nil {
			return
		}
		if cert.AccountID == nil {
			return
		}
		if cert.TenantID == nil {
			return
		}
		securityConfig := config.TomlConfig().Security
		recorder := recorderx.FetchRecorder(ctx)
		emitter, release := redisx.NewRedisEmitter(recorder)
		defer release()
		_ = emitter.SetXX(emitter.BuildKey(constant.KeyDelimiter, token), redisx.BuildVal(strconv.FormatInt(*cert.AccountID, 10), strconv.FormatInt(*cert.TenantID, 10)), securityConfig.AccessTTL)
	}()
}

// GetByToken 获取任意令牌
func GetByToken(ctx *gin.Context, prefix, token string) string {
	if len(token) == 0 {
		code := runerror.GetSysErp(runerror.ModAuth, runerror.LenError)
		panic(runerror.NewAll(code, runerror.NewError("令牌长度错误")))
	}
	recorder := recorderx.FetchRecorder(ctx)
	emitter, release := redisx.NewRedisEmitter(recorder)
	defer release()
	info := emitter.Get(emitter.BuildKey(prefix, token))
	if len(info) == 0 {
		code := runerror.GetUsrErp(runerror.ModAuth, runerror.NotLoginError)
		panic(runerror.NewAll(code, runerror.NewError("用户未登录")))
	}
	return info
}

// SetCert 设置认证
func SetCert(ctx *gin.Context, cert *auth.CertDetail) {
	if cert == nil {
		return
	}
	ctx.Set(constant.CtxCertKey, cert)
}

// GetCert 获取认证
func GetCert(ctx *gin.Context) *auth.CertDetail {
	value, exists := ctx.Get(constant.CtxCertKey)
	if !exists {
		code := runerror.GetUsrErp(runerror.ModAuth, runerror.NotFoundUserError)
		panic(runerror.New().WithCode(code).WithMessage("无法查找用户信息"))
	}
	cert, ok := value.(*auth.CertDetail)
	if ok {
		return cert
	}
	code := runerror.GetSysErp(runerror.ModAuth, runerror.ConvertError)
	panic(runerror.New().WithCode(code).WithMessage("获取到用户信息, 但是无法转换"))
}

// GetNillableCert 获取认证(可为 NIL)
func GetNillableCert(ctx *gin.Context) *auth.CertDetail {
	value, exists := ctx.Get(constant.CtxCertKey)
	if !exists {
		return nil
	}
	cert, ok := value.(*auth.CertDetail)
	if ok {
		return cert
	}
	code := runerror.GetSysErp(runerror.ModAuth, runerror.ConvertError)
	panic(runerror.New().WithCode(code).WithMessage("获取到用户信息, 但是无法转换"))
}

// GetAccountID 获取账号 ID
func GetAccountID(ctx *gin.Context) int64 {
	// cert := GetCert(ctx)
	cert := GetNillableCert(ctx)
	if cert == nil {
		// panic(runerror.NewError("无法获取认证"))
		return 0
	}
	if cert.AccountID == nil {
		// panic(runerror.NewError("无法获取账户"))
		return 0
	}
	return *cert.AccountID
}

// GetAccountIDString 获取账号 ID
func GetAccountIDString(ctx *gin.Context) string {
	// cert := GetCert(ctx)
	cert := GetNillableCert(ctx)
	if cert == nil {
		// panic(runerror.NewError("无法获取认证"))
		return "0"
	}
	if cert.AccountID == nil {
		// panic(runerror.NewError("无法获取账户"))
		return "0"
	}
	return strconv.FormatInt(*cert.AccountID, 10)
}

// GetTenantID 获取租户 ID
func GetTenantID(ctx *gin.Context) int64 {
	// cert := GetCert(ctx)
	cert := GetNillableCert(ctx)
	if cert == nil {
		// panic(runerror.NewError("无法获取认证"))
		return 111000
	}
	if cert.TenantID == nil {
		// panic(runerror.NewError("无法获取租户"))
		return 111000
	}
	return *cert.TenantID
}

// GetIP 获取 IP
func GetIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if len(ip) == 0 {
		return "127.0.0.1"
	}
	return ip
}

type aclOpr interface {
	FindValidNameByAccountId(aid int64) []string
}

// LookupRole 初始化角色
func LookupRole(ctx *gin.Context, acl aclOpr) {
	roles := LoadingACL(ctx, acl, constant.RolePrefix)

	if roles != nil {
		SetRoleIntoCtx(ctx, *roles...)
	}
}

// LookupPerm 初始化权限
func LookupPerm(ctx *gin.Context, acl aclOpr) {

	perms := LoadingACL(ctx, acl, constant.PermPrefix)

	if perms != nil {
		SetPermIntoCtx(ctx, *perms...)
	}
}

func DiscardACL(ctx *gin.Context, aid int64) bool {
	recorder := recorderx.FetchRecorder(ctx)
	emitter, release := redisx.NewRedisEmitter(recorder)
	defer release()
	id := strconv.FormatInt(aid, 10)
	return emitter.Del(emitter.BuildKey(constant.RolePrefix, id), emitter.BuildKey(constant.PermPrefix, id))
}

func LoadingACL(ctx *gin.Context, acl aclOpr, prefix string) *[]string {
	cert := GetCert(ctx)
	if cert == nil {
		return nil
	}
	if cert.AccountID == nil {
		return nil
	}
	recorder := recorderx.FetchRecorder(ctx)
	emitter, release := redisx.NewRedisEmitter(recorder)
	aid := strconv.FormatInt(*cert.AccountID, 10)
	defer release()

	key := emitter.BuildKey(prefix, aid)
	readFromCache := func(id *int64) *[]string {
		members := emitter.SMembers(key)
		if len(members) == 0 {
			return nil
		}
		return &members
	}

	writeToCache := func(id *int64, members *[]string) {
		emitter.SAdd(key, helper.ToAnyItems(*members)...)
		emitter.Expire(key, 30*time.Minute)
	}

	readFromDB := func(id *int64) *[]string {
		members := acl.FindValidNameByAccountId(*cert.AccountID)
		if len(members) == 0 {
			return nil
		}
		return &members
	}

	return sonar.Lookup[int64, []string](readFromCache, cert.AccountID).Backing(readFromDB, writeToCache).Get()
}

// SetRoleIntoCtx 设置角色
func SetRoleIntoCtx(ctx *gin.Context, roles ...string) {
	if len(roles) == 0 {
		return
	}
	ctx.Set(constant.CtxRoleKey, roles)
}

// GetRoleFromCtx 获取角色
func GetRoleFromCtx(ctx *gin.Context) []string {
	value, exists := ctx.Get(constant.CtxRoleKey)
	if exists {
		return value.([]string)
	}
	return []string{}
}

// SetPermIntoCtx 设置权限
func SetPermIntoCtx(ctx *gin.Context, perms ...string) {
	if len(perms) == 0 {
		return
	}
	ctx.Set(constant.CtxPermKey, perms)
}

// GetPermFromCtx 获取权限
func GetPermFromCtx(ctx *gin.Context) []string {
	value, exists := ctx.Get(constant.CtxPermKey)
	if exists {
		return value.([]string)
	}
	return []string{}
}
