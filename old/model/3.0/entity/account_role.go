// Package entity
// @author tabuyos
// @since 2023/9/12
// @description entity
package entity

import (
	iris2 "metis/old/model/3.0/iris"
	"time"
)

type AccountRole struct {
	// 主键ID
	ID *int64 `json:"id"`
	// 账号ID
	AccountId *int64 `json:"accountId"`
	// 角色ID
	RoleId *int64 `json:"roleId"`
	// 状态 0-停用;1-启用
	Status *int `json:"status"`
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

	evaluator *iris2.Evaluator[AccountRole]
}

// NewAccountRole 初始化
func NewAccountRole() *AccountRole {
	return &AccountRole{}
}

func (e *AccountRole) IDCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("id", func(rec *AccountRole) any {
		return &rec.ID
	})
}
func (e *AccountRole) AccountIdCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("account_id", func(rec *AccountRole) any {
		return &rec.AccountId
	})
}
func (e *AccountRole) RoleIdCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("role_id", func(rec *AccountRole) any {
		return &rec.RoleId
	})
}
func (e *AccountRole) StatusCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("status", func(rec *AccountRole) any {
		return &rec.Status
	})
}
func (e *AccountRole) CreateByCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("create_by", func(rec *AccountRole) any {
		return &rec.CreateBy
	})
}
func (e *AccountRole) CreateAtCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("create_at", func(rec *AccountRole) any {
		return &rec.CreateAt
	})
}
func (e *AccountRole) ModifyByCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("modify_by", func(rec *AccountRole) any {
		return &rec.ModifyBy
	})
}
func (e *AccountRole) ModifyAtCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("modify_at", func(rec *AccountRole) any {
		return &rec.ModifyAt
	})
}
func (e *AccountRole) DeletedCol() *iris2.Column[AccountRole] {
	return iris2.WithColumn("deleted", func(rec *AccountRole) any {
		return &rec.Deleted
	})
}

func (e *AccountRole) Configure(fn func(*iris2.Evaluator[AccountRole])) {
	if e.evaluator == nil {
		e.evaluator = iris2.WithLogicalEvaluator[AccountRole]()
	}
	fn(e.evaluator)
}

func (e *AccountRole) Asterisk(fns ...func(string) string) []*iris2.Column[AccountRole] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris2.Column[AccountRole]{
		e.IDCol().Decorate(fn),
		e.AccountIdCol().Decorate(fn),
		e.RoleIdCol().Decorate(fn),
		e.StatusCol().Decorate(fn),
		e.CreateByCol().Decorate(fn),
		e.CreateAtCol().Decorate(fn),
		e.ModifyByCol().Decorate(fn),
		e.ModifyAtCol().Decorate(fn),
		e.DeletedCol().Decorate(fn),
	}
}

func (e *AccountRole) PKey() *iris2.Column[AccountRole] {
	return e.IDCol()
}

func (e *AccountRole) LogicDelKey() *iris2.Column[AccountRole] {
	return e.DeletedCol()
}

func (e *AccountRole) Evaluator() iris2.EvalInfoService[AccountRole] {
	return e.evaluator
}

func (e *AccountRole) Table() *iris2.RefTable {
	return iris2.WithTable("account_role")
}

func (e *AccountRole) Self() *AccountRole {
	return e
}
