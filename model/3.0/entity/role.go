// Package entity
// @author tabuyos
// @since 2023/9/13
// @description entity
package entity

import (
	"metis/model/3.0/iris"
	"time"
)

type Role struct {
	// 主键ID
	ID *int64 `json:"id"`
	// 角色父ID
	Pid *int64 `json:"pid"`
	// 平台ID
	PlatId *int64 `json:"platId"`
	// 租户ID
	TenantId *int64 `json:"tenantId"`
	// 角色名
	Name *string `json:"name"`
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

	evaluator *iris.Evaluator[Role]
}

// NewRole 初始化
func NewRole() *Role {
	return &Role{}
}

func (e *Role) IDCol() *iris.Column[Role] {
	return iris.WithColumn("id", func(rec *Role) any {
		return &rec.ID
	})
}
func (e *Role) PidCol() *iris.Column[Role] {
	return iris.WithColumn("pid", func(rec *Role) any {
		return &rec.Pid
	})
}
func (e *Role) PlatIdCol() *iris.Column[Role] {
	return iris.WithColumn("plat_id", func(rec *Role) any {
		return &rec.PlatId
	})
}
func (e *Role) TenantIdCol() *iris.Column[Role] {
	return iris.WithColumn("tenant_id", func(rec *Role) any {
		return &rec.TenantId
	})
}
func (e *Role) NameCol() *iris.Column[Role] {
	return iris.WithColumn("name", func(rec *Role) any {
		return &rec.Name
	})
}
func (e *Role) StatusCol() *iris.Column[Role] {
	return iris.WithColumn("status", func(rec *Role) any {
		return &rec.Status
	})
}
func (e *Role) CreateByCol() *iris.Column[Role] {
	return iris.WithColumn("create_by", func(rec *Role) any {
		return &rec.CreateBy
	})
}
func (e *Role) CreateAtCol() *iris.Column[Role] {
	return iris.WithColumn("create_at", func(rec *Role) any {
		return &rec.CreateAt
	})
}
func (e *Role) ModifyByCol() *iris.Column[Role] {
	return iris.WithColumn("modify_by", func(rec *Role) any {
		return &rec.ModifyBy
	})
}
func (e *Role) ModifyAtCol() *iris.Column[Role] {
	return iris.WithColumn("modify_at", func(rec *Role) any {
		return &rec.ModifyAt
	})
}
func (e *Role) DeletedCol() *iris.Column[Role] {
	return iris.WithColumn("deleted", func(rec *Role) any {
		return &rec.Deleted
	})
}

func (e *Role) Configure(fn func(*iris.Evaluator[Role])) {
	if e.evaluator == nil {
		e.evaluator = iris.WithLogicalEvaluator[Role]()
	}
	fn(e.evaluator)
}

func (e *Role) ColumnAndValue(fns ...func(*iris.Column[Role], any) bool) (selfishs []iris.Selfish, values []any) {
	fn := func(*iris.Column[Role], any) bool {
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

func (e *Role) Asterisk(fns ...func(string) string) []*iris.Column[Role] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris.Column[Role]{
		e.IDCol().Decorate(fn),
		e.PidCol().Decorate(fn),
		e.PlatIdCol().Decorate(fn),
		e.TenantIdCol().Decorate(fn),
		e.NameCol().Decorate(fn),
		e.StatusCol().Decorate(fn),
		e.CreateByCol().Decorate(fn),
		e.CreateAtCol().Decorate(fn),
		e.ModifyByCol().Decorate(fn),
		e.ModifyAtCol().Decorate(fn),
		e.DeletedCol().Decorate(fn),
	}
}

func (e *Role) PKey() *iris.Column[Role] {
	return e.IDCol()
}

func (e *Role) LogicDelKey() *iris.Column[Role] {
	return e.DeletedCol()
}

func (e *Role) Evaluator() iris.EvalInfoService[Role] {
	return e.evaluator
}

func (e *Role) Table() *iris.RefTable {
	return iris.WithTable("role")
}

func (e *Role) Self() *Role {
	return e
}
