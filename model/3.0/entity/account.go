// Package entity
// @author tabuyos
// @since 2023/9/12
// @description entity
package entity

import (
	"metis/model/3.0/iris"
	"time"
)

type Account struct {
	// 主键ID
	ID *int64 `json:"id"`
	// 租户ID
	TenantId *int64 `json:"tenantId"`
	// 账户名
	Name *string `json:"name"`
	// 密码
	Pwd *string `json:"pwd"`
	// 绑定手机号
	Mobile *string `json:"mobile"`
	// 绑定邮箱
	Email *string `json:"email"`
	// 绑定用户
	BindUser *int64 `json:"bindUser"`
	// 所属部门
	DeptId *int64 `json:"deptId"`
	// 岗位
	PostIds *string `json:"postIds"`
	// 头像
	Avatar *string `json:"avatar"`
	// 登录IP
	LoginIp *string `json:"loginIp"`
	// 最后登录时间
	LoginTime *time.Time `json:"loginTime"`
	// 登录次数
	Logins *int `json:"logins"`
	// 登录失败次数
	LoginErrors *int `json:"loginErrors"`
	// 在线数
	Onlines *int `json:"onlines"`
	// 状态 0-停用;1-启用
	Status *int8 `json:"status"`
	// 创建者
	CreateBy *int64 `json:"createBy"`
	// 创建时间
	CreateAt *time.Time `json:"createAt"`
	// 更新人
	ModifyBy *int64 `json:"modifyBy"`
	// 更新时间
	ModifyAt *time.Time `json:"modifyAt"`
	// 逻辑删除 0-否 1-是
	Deleted *int8 `json:"deleted"`

	evaluator *iris.Evaluator[Account]
}

// NewAccount 初始化
func NewAccount() *Account {
	return &Account{}
}

func (e *Account) IDCol() *iris.Column[Account] {
	return iris.WithColumn("id", func(rec *Account) any {
		return &rec.ID
	})
}
func (e *Account) TenantIdCol() *iris.Column[Account] {
	return iris.WithColumn("tenant_id", func(rec *Account) any {
		return &rec.TenantId
	})
}
func (e *Account) NameCol() *iris.Column[Account] {
	return iris.WithColumn("name", func(rec *Account) any {
		return &rec.Name
	})
}
func (e *Account) PwdCol() *iris.Column[Account] {
	return iris.WithColumn("pwd", func(rec *Account) any {
		return &rec.Pwd
	})
}
func (e *Account) MobileCol() *iris.Column[Account] {
	return iris.WithColumn("mobile", func(rec *Account) any {
		return &rec.Mobile
	})
}
func (e *Account) EmailCol() *iris.Column[Account] {
	return iris.WithColumn("email", func(rec *Account) any {
		return &rec.Email
	})
}
func (e *Account) BindUserCol() *iris.Column[Account] {
	return iris.WithColumn("bind_user", func(rec *Account) any {
		return &rec.BindUser
	})
}
func (e *Account) DeptIdCol() *iris.Column[Account] {
	return iris.WithColumn("dept_id", func(rec *Account) any {
		return &rec.DeptId
	})
}
func (e *Account) PostIdsCol() *iris.Column[Account] {
	return iris.WithColumn("post_ids", func(rec *Account) any {
		return &rec.PostIds
	})
}
func (e *Account) AvatarCol() *iris.Column[Account] {
	return iris.WithColumn("avatar", func(rec *Account) any {
		return &rec.Avatar
	})
}
func (e *Account) LoginIpCol() *iris.Column[Account] {
	return iris.WithColumn("login_ip", func(rec *Account) any {
		return &rec.LoginIp
	})
}
func (e *Account) LoginTimeCol() *iris.Column[Account] {
	return iris.WithColumn("login_time", func(rec *Account) any {
		return &rec.LoginTime
	})
}
func (e *Account) LoginsCol() *iris.Column[Account] {
	return iris.WithColumn("logins", func(rec *Account) any {
		return &rec.Logins
	})
}
func (e *Account) LoginErrorsCol() *iris.Column[Account] {
	return iris.WithColumn("login_errors", func(rec *Account) any {
		return &rec.LoginErrors
	})
}
func (e *Account) OnlinesCol() *iris.Column[Account] {
	return iris.WithColumn("onlines", func(rec *Account) any {
		return &rec.Onlines
	})
}
func (e *Account) StatusCol() *iris.Column[Account] {
	return iris.WithColumn("status", func(rec *Account) any {
		return &rec.Status
	})
}
func (e *Account) CreateByCol() *iris.Column[Account] {
	return iris.WithColumn("create_by", func(rec *Account) any {
		return &rec.CreateBy
	})
}
func (e *Account) CreateAtCol() *iris.Column[Account] {
	return iris.WithColumn("create_at", func(rec *Account) any {
		return &rec.CreateAt
	})
}
func (e *Account) ModifyByCol() *iris.Column[Account] {
	return iris.WithColumn("modify_by", func(rec *Account) any {
		return &rec.ModifyBy
	})
}
func (e *Account) ModifyAtCol() *iris.Column[Account] {
	return iris.WithColumn("modify_at", func(rec *Account) any {
		return &rec.ModifyAt
	})
}
func (e *Account) DeletedCol() *iris.Column[Account] {
	return iris.WithColumn("deleted", func(rec *Account) any {
		return &rec.Deleted
	})
}

func (e *Account) Configure(fn func(*iris.Evaluator[Account])) {
	if e.evaluator == nil {
		e.evaluator = iris.WithLogicalEvaluator[Account]()
	}
	fn(e.evaluator)
}

func (e *Account) ColumnAndValue(fns ...func(*iris.Column[Account], any) bool) (selfishs []iris.Selfish, values []any) {
	fn := func(*iris.Column[Account], any) bool {
		return true
	}
	if len(fns) > 0 {
		fn = fns[0]
	}
	if fn(e.IDCol(), e.ID) {
		selfishs = append(selfishs, e.IDCol())
		values = append(values, *e.ID)
	}
	return
}

func (e *Account) Asterisk(fns ...func(string) string) []*iris.Column[Account] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris.Column[Account]{
		e.IDCol().Decorate(fn),
		e.TenantIdCol().Decorate(fn),
		e.NameCol().Decorate(fn),
		e.PwdCol().Decorate(fn),
		e.MobileCol().Decorate(fn),
		e.EmailCol().Decorate(fn),
		e.BindUserCol().Decorate(fn),
		e.DeptIdCol().Decorate(fn),
		e.PostIdsCol().Decorate(fn),
		e.AvatarCol().Decorate(fn),
		e.LoginIpCol().Decorate(fn),
		e.LoginTimeCol().Decorate(fn),
		e.LoginsCol().Decorate(fn),
		e.LoginErrorsCol().Decorate(fn),
		e.OnlinesCol().Decorate(fn),
		e.StatusCol().Decorate(fn),
		e.CreateByCol().Decorate(fn),
		e.CreateAtCol().Decorate(fn),
		e.ModifyByCol().Decorate(fn),
		e.ModifyAtCol().Decorate(fn),
		e.DeletedCol().Decorate(fn),
	}
}

func (e *Account) PKey() *iris.Column[Account] {
	return e.IDCol()
}

func (e *Account) LogicDelKey() *iris.Column[Account] {
	return e.DeletedCol()
}

func (e *Account) Evaluator() iris.EvalInfoService[Account] {
	if e == nil {
		return nil
	}
	return e.evaluator
}

func (e *Account) Table() *iris.RefTable {
	return iris.WithTable("account")
}

func (e *Account) Self() *Account {
	return e
}
