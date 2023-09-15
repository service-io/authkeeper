package entity

import (
	"deepsea/model/iris"
	"time"
)

type MyPlatAuthorityResource struct {
	// 主键ID
	ID *int64 `json:"id"`
	// 权限ID
	AuthorityId *int64 `json:"authorityId"`
	// 资源ID
	ResourceId *int64 `json:"resourceId"`
	// 平台ID
	PlatId *int64 `json:"platId"`
	// 租户ID
	TenantId *int64 `json:"tenantId"`
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

	evaluator *iris.Evaluator[MyPlatAuthorityResource]
}

// NewMyPlatAuthorityResource 初始化
func NewMyPlatAuthorityResource() *MyPlatAuthorityResource {
	return &MyPlatAuthorityResource{}
}

// IDCol ID 列
func (e *MyPlatAuthorityResource) IDCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("id", func(rec *MyPlatAuthorityResource) any {
		return &rec.ID
	})
}

// AuthorityIdCol AuthorityId 列
func (e *MyPlatAuthorityResource) AuthorityIdCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("authority_id", func(rec *MyPlatAuthorityResource) any {
		return &rec.AuthorityId
	})
}

// ResourceIdCol ResourceId 列
func (e *MyPlatAuthorityResource) ResourceIdCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("resource_id", func(rec *MyPlatAuthorityResource) any {
		return &rec.ResourceId
	})
}

// PlatIdCol PlatId 列
func (e *MyPlatAuthorityResource) PlatIdCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("plat_id", func(rec *MyPlatAuthorityResource) any {
		return &rec.PlatId
	})
}

// TenantIdCol TenantId 列
func (e *MyPlatAuthorityResource) TenantIdCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("tenant_id", func(rec *MyPlatAuthorityResource) any {
		return &rec.TenantId
	})
}

// CreateByCol CreateBy 列
func (e *MyPlatAuthorityResource) CreateByCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("create_by", func(rec *MyPlatAuthorityResource) any {
		return &rec.CreateBy
	})
}

// CreateAtCol CreateAt 列
func (e *MyPlatAuthorityResource) CreateAtCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("create_at", func(rec *MyPlatAuthorityResource) any {
		return &rec.CreateAt
	})
}

// ModifyByCol ModifyBy 列
func (e *MyPlatAuthorityResource) ModifyByCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("modify_by", func(rec *MyPlatAuthorityResource) any {
		return &rec.ModifyBy
	})
}

// ModifyAtCol ModifyAt 列
func (e *MyPlatAuthorityResource) ModifyAtCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("modify_at", func(rec *MyPlatAuthorityResource) any {
		return &rec.ModifyAt
	})
}

// DeletedCol Deleted 列
func (e *MyPlatAuthorityResource) DeletedCol() *iris.Column[MyPlatAuthorityResource] {
	return iris.WithColumn("deleted", func(rec *MyPlatAuthorityResource) any {
		return &rec.Deleted
	})
}

// Configure evaluator 配置
func (e *MyPlatAuthorityResource) Configure(fn func(*iris.Evaluator[MyPlatAuthorityResource])) {
	if e.evaluator == nil {
		e.evaluator = iris.WithLogicalEvaluator[MyPlatAuthorityResource]()
	}
	fn(e.evaluator)
}

// ColumnAndValue 列值计算
func (e *MyPlatAuthorityResource) ColumnAndValue(fns ...func(*iris.Column[MyPlatAuthorityResource], any) bool) (selfishs []iris.Selfish, values []any) {
	fn := func(*iris.Column[MyPlatAuthorityResource], any) bool {
		return true
	}
	if len(fns) > 0 {
		fn = fns[0]
	}

	if fn(e.IDCol(), e.ID) {
		selfishs = append(selfishs, e.IDCol())
		values = append(values, *e.ID)
	}
	if fn(e.AuthorityIdCol(), e.AuthorityId) {
		selfishs = append(selfishs, e.AuthorityIdCol())
		values = append(values, *e.AuthorityId)
	}
	if fn(e.ResourceIdCol(), e.ResourceId) {
		selfishs = append(selfishs, e.ResourceIdCol())
		values = append(values, *e.ResourceId)
	}
	if fn(e.PlatIdCol(), e.PlatId) {
		selfishs = append(selfishs, e.PlatIdCol())
		values = append(values, *e.PlatId)
	}
	if fn(e.TenantIdCol(), e.TenantId) {
		selfishs = append(selfishs, e.TenantIdCol())
		values = append(values, *e.TenantId)
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
func (e *MyPlatAuthorityResource) Asterisk(fns ...func(string) string) []*iris.Column[MyPlatAuthorityResource] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris.Column[MyPlatAuthorityResource]{
		e.IDCol().Decorate(fn),
		e.AuthorityIdCol().Decorate(fn),
		e.ResourceIdCol().Decorate(fn),
		e.PlatIdCol().Decorate(fn),
		e.TenantIdCol().Decorate(fn),
		e.CreateByCol().Decorate(fn),
		e.CreateAtCol().Decorate(fn),
		e.ModifyByCol().Decorate(fn),
		e.ModifyAtCol().Decorate(fn),
		e.DeletedCol().Decorate(fn),
	}
}

// PKey 主键
func (e *MyPlatAuthorityResource) PKey() *iris.Column[MyPlatAuthorityResource] {
	return e.IDCol()
}

// LogicDelKey 逻辑删除
func (e *MyPlatAuthorityResource) LogicDelKey() *iris.Column[MyPlatAuthorityResource] {
	return e.DeletedCol()
}

// Evaluator 计算器
func (e *MyPlatAuthorityResource) Evaluator() *iris.Evaluator[MyPlatAuthorityResource] {
	if e == nil {
		return nil
	}
	return e.evaluator
}

// Table 表
func (e *MyPlatAuthorityResource) Table() *iris.RefTable {
	return iris.WithTable("my_plat_authority_resource")
}

// Self 原始信息
func (e *MyPlatAuthorityResource) Self() *MyPlatAuthorityResource {
	return e
}
