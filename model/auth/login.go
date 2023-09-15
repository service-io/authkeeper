// Package auth
// @author tabuyos
// @since 2023/8/6
// @description auth
package auth

type Login struct {
	// 账号名称
	Name *string `json:"name"`
	// 账号密码
	Password *string `json:"password"`
	// 租户ID
	TenantID *int64 `json:"tenantId"`
	// 验证码
	VerifyCode *string `json:"verifyCode"`
	// 手机号
	Tel *string `json:"tel"`
	// 邮箱
	Email *string `json:"email"`
}

type Refresh struct {
	// 刷新令牌
	Token *string `json:"token"`
	// 登录名称
	Name *string `json:"name"`
}
