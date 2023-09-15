// Package auth
// @author tabuyos
// @since 2023/8/7
// @description auth
package auth

import "fmt"

type Register struct {
	// 头像
	Avatar *string `json:"avatar"`
	// 绑定邮箱
	Email *string `json:"email"`
	// 绑定手机号
	Mobile *string `json:"mobile"`
	// 账户名
	Name *string `json:"name"`
	// 密码
	Pwd *string `json:"pwd"`
	// 租户ID
	TenantId *int64 `json:"tenantId"`
}

func replaceToStar(pwd string) string {
	bytes := make([]byte, len(pwd))
	for i := 0; i < len(pwd); i++ {
		bytes[i] = '*'
	}
	return string(bytes)
}

func (rec *Register) String() string {
	return fmt.Sprintf("Register(Avatar: %v, Email: %v, Mobile: %v, Name: %v, Pwd: %v, TenantId: %v)", *rec.Avatar, *rec.Email, *rec.Mobile, *rec.Name, replaceToStar(*rec.Pwd), *rec.TenantId)
}
