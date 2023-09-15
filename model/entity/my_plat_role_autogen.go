package entity

import (
	"deepsea/model/iris"
	"time"
)

type MyPlatRole struct {
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

	evaluator *iris.Evaluator[MyPlatRole]
}

// NewMyPlatRole 初始化
func NewMyPlatRole() *MyPlatRole {
	return &MyPlatRole{}
}

// IDCol ID 列
func (e *MyPlatRole) IDCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("id", func(rec *MyPlatRole) any {
		return &rec.ID
	})
}

// PidCol Pid 列
func (e *MyPlatRole) PidCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("pid", func(rec *MyPlatRole) any {
		return &rec.Pid
	})
}

// PlatIdCol PlatId 列
func (e *MyPlatRole) PlatIdCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("plat_id", func(rec *MyPlatRole) any {
		return &rec.PlatId
	})
}

// TenantIdCol TenantId 列
func (e *MyPlatRole) TenantIdCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("tenant_id", func(rec *MyPlatRole) any {
		return &rec.TenantId
	})
}

// NameCol Name 列
func (e *MyPlatRole) NameCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("name", func(rec *MyPlatRole) any {
		return &rec.Name
	})
}

// StatusCol Status 列
func (e *MyPlatRole) StatusCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("status", func(rec *MyPlatRole) any {
		return &rec.Status
	})
}

// CreateByCol CreateBy 列
func (e *MyPlatRole) CreateByCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("create_by", func(rec *MyPlatRole) any {
		return &rec.CreateBy
	})
}

// CreateAtCol CreateAt 列
func (e *MyPlatRole) CreateAtCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("create_at", func(rec *MyPlatRole) any {
		return &rec.CreateAt
	})
}

// ModifyByCol ModifyBy 列
func (e *MyPlatRole) ModifyByCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("modify_by", func(rec *MyPlatRole) any {
		return &rec.ModifyBy
	})
}

// ModifyAtCol ModifyAt 列
func (e *MyPlatRole) ModifyAtCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("modify_at", func(rec *MyPlatRole) any {
		return &rec.ModifyAt
	})
}

// DeletedCol Deleted 列
func (e *MyPlatRole) DeletedCol() *iris.Column[MyPlatRole] {
	return iris.WithColumn("deleted", func(rec *MyPlatRole) any {
		return &rec.Deleted
	})
}

// Configure evaluator 配置
func (e *MyPlatRole) Configure(fn func(*iris.Evaluator[MyPlatRole])) {
	if e.evaluator == nil {
		e.evaluator = iris.WithLogicalEvaluator[MyPlatRole]()
	}
	fn(e.evaluator)
}

// ColumnAndValue 列值计算
func (e *MyPlatRole) ColumnAndValue(fns ...func(*iris.Column[MyPlatRole], any) bool) (selfishs []iris.Selfish, values []any) {
	fn := func(*iris.Column[MyPlatRole], any) bool {
		return true
	}
	if len(fns) > 0 {
		fn = fns[0]
	}

	if fn(e.IDCol(), e.ID) {
		selfishs = append(selfishs, e.IDCol())
		values = append(values, *e.ID)
	}
	if fn(e.PidCol(), e.Pid) {
		selfishs = append(selfishs, e.PidCol())
		values = append(values, *e.Pid)
	}
	if fn(e.PlatIdCol(), e.PlatId) {
		selfishs = append(selfishs, e.PlatIdCol())
		values = append(values, *e.PlatId)
	}
	if fn(e.TenantIdCol(), e.TenantId) {
		selfishs = append(selfishs, e.TenantIdCol())
		values = append(values, *e.TenantId)
	}
	if fn(e.NameCol(), e.Name) {
		selfishs = append(selfishs, e.NameCol())
		values = append(values, *e.Name)
	}
	if fn(e.StatusCol(), e.Status) {
		selfishs = append(selfishs, e.StatusCol())
		values = append(values, *e.Status)
	}
	if fn(e.CreateByCol(), e.CreateBy) {
		selfishs = append(selfishs, e.CreateByCol())
		values = append(values, *e.CreateBy)
	}
	if fn(e.CreateAtCol(), e.CreateAt) {
		selfishs = append(selfishs, e.CreateAtCol())
		values = append(values, *e.CreateAt)
	}
	if fn(e.ModifyByCol(), e.ModifyBy) {
		selfishs = append(selfishs, e.ModifyByCol())
		values = append(values, *e.ModifyBy)
	}
	if fn(e.ModifyAtCol(), e.ModifyAt) {
		selfishs = append(selfishs, e.ModifyAtCol())
		values = append(values, *e.ModifyAt)
	}
	if fn(e.DeletedCol(), e.Deleted) {
		selfishs = append(selfishs, e.DeletedCol())
		values = append(values, *e.Deleted)
	}
	return
}

// Asterisk 所有列
func (e *MyPlatRole) Asterisk(fns ...func(string) string) []*iris.Column[MyPlatRole] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris.Column[MyPlatRole]{
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

// PKey 主键
func (e *MyPlatRole) PKey() *iris.Column[MyPlatRole] {
	return e.IDCol()
}

// LogicDelKey 逻辑删除
func (e *MyPlatRole) LogicDelKey() *iris.Column[MyPlatRole] {
	return e.DeletedCol()
}

// Evaluator 计算器
func (e *MyPlatRole) Evaluator() *iris.Evaluator[MyPlatRole] {
	if e == nil {
		return nil
	}
	return e.evaluator
}

// Table 表
func (e *MyPlatRole) Table() *iris.RefTable {
	return iris.WithTable("my_plat_role")
}

// Self 原始信息
func (e *MyPlatRole) Self() *MyPlatRole {
	return e
}
