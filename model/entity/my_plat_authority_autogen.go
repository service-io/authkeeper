package entity

import (
	"deepsea/model/iris"
	"time"
)

type MyPlatAuthority struct {
	// 主键ID
	ID *int64 `json:"id"`
	// 租户ID
	TenantId *int64 `json:"tenantId"`
	// 平台ID
	PlatId *int64 `json:"platId"`
	// 权限名
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

	evaluator *iris.Evaluator[MyPlatAuthority]
}

// NewMyPlatAuthority 初始化
func NewMyPlatAuthority() *MyPlatAuthority {
	return &MyPlatAuthority{}
}

// IDCol ID 列
func (e *MyPlatAuthority) IDCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("id", func(rec *MyPlatAuthority) any {
		return &rec.ID
	})
}

// TenantIdCol TenantId 列
func (e *MyPlatAuthority) TenantIdCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("tenant_id", func(rec *MyPlatAuthority) any {
		return &rec.TenantId
	})
}

// PlatIdCol PlatId 列
func (e *MyPlatAuthority) PlatIdCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("plat_id", func(rec *MyPlatAuthority) any {
		return &rec.PlatId
	})
}

// NameCol Name 列
func (e *MyPlatAuthority) NameCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("name", func(rec *MyPlatAuthority) any {
		return &rec.Name
	})
}

// StatusCol Status 列
func (e *MyPlatAuthority) StatusCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("status", func(rec *MyPlatAuthority) any {
		return &rec.Status
	})
}

// CreateByCol CreateBy 列
func (e *MyPlatAuthority) CreateByCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("create_by", func(rec *MyPlatAuthority) any {
		return &rec.CreateBy
	})
}

// CreateAtCol CreateAt 列
func (e *MyPlatAuthority) CreateAtCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("create_at", func(rec *MyPlatAuthority) any {
		return &rec.CreateAt
	})
}

// ModifyByCol ModifyBy 列
func (e *MyPlatAuthority) ModifyByCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("modify_by", func(rec *MyPlatAuthority) any {
		return &rec.ModifyBy
	})
}

// ModifyAtCol ModifyAt 列
func (e *MyPlatAuthority) ModifyAtCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("modify_at", func(rec *MyPlatAuthority) any {
		return &rec.ModifyAt
	})
}

// DeletedCol Deleted 列
func (e *MyPlatAuthority) DeletedCol() *iris.Column[MyPlatAuthority] {
	return iris.WithColumn("deleted", func(rec *MyPlatAuthority) any {
		return &rec.Deleted
	})
}

// Configure evaluator 配置
func (e *MyPlatAuthority) Configure(fn func(*iris.Evaluator[MyPlatAuthority])) {
	if e.evaluator == nil {
		e.evaluator = iris.WithLogicalEvaluator[MyPlatAuthority]()
	}
	fn(e.evaluator)
}

// ColumnAndValue 列值计算
func (e *MyPlatAuthority) ColumnAndValue(fns ...func(*iris.Column[MyPlatAuthority], any) bool) (selfishs []iris.Selfish, values []any) {
	fn := func(*iris.Column[MyPlatAuthority], any) bool {
		return true
	}
	if len(fns) > 0 {
		fn = fns[0]
	}

	if fn(e.IDCol(), e.ID) {
		selfishs = append(selfishs, e.IDCol())
		values = append(values, *e.ID)
	}
	if fn(e.TenantIdCol(), e.TenantId) {
		selfishs = append(selfishs, e.TenantIdCol())
		values = append(values, *e.TenantId)
	}
	if fn(e.PlatIdCol(), e.PlatId) {
		selfishs = append(selfishs, e.PlatIdCol())
		values = append(values, *e.PlatId)
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
func (e *MyPlatAuthority) Asterisk(fns ...func(string) string) []*iris.Column[MyPlatAuthority] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris.Column[MyPlatAuthority]{
		e.IDCol().Decorate(fn),
		e.TenantIdCol().Decorate(fn),
		e.PlatIdCol().Decorate(fn),
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
func (e *MyPlatAuthority) PKey() *iris.Column[MyPlatAuthority] {
	return e.IDCol()
}

// LogicDelKey 逻辑删除
func (e *MyPlatAuthority) LogicDelKey() *iris.Column[MyPlatAuthority] {
	return e.DeletedCol()
}

// Evaluator 计算器
func (e *MyPlatAuthority) Evaluator() *iris.Evaluator[MyPlatAuthority] {
	if e == nil {
		return nil
	}
	return e.evaluator
}

// Table 表
func (e *MyPlatAuthority) Table() *iris.RefTable {
	return iris.WithTable("my_plat_authority")
}

// Self 原始信息
func (e *MyPlatAuthority) Self() *MyPlatAuthority {
	return e
}
