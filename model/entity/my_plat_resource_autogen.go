package entity

import (
	"deepsea/model/iris"
	"time"
)

type MyPlatResource struct {
	// 主键ID
	ID *int64 `json:"id"`
	// 父ID
	Pid *int64 `json:"pid"`
	// 平台ID
	PlatId *int64 `json:"platId"`
	// 资源标题
	Title *string `json:"title"`
	// 资源名
	Name *string `json:"name"`
	// 图标
	Icon *string `json:"icon"`
	// 资源类型 1-菜单;2-按钮;3-api
	Type *int `json:"type"`
	// 权限标识
	Permission *string `json:"permission"`
	// api请求地址
	Path *string `json:"path"`
	// 前端路由
	Router *string `json:"router"`
	// 排序
	Sort *string `json:"sort"`
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

	evaluator *iris.Evaluator[MyPlatResource]
}

// NewMyPlatResource 初始化
func NewMyPlatResource() *MyPlatResource {
	return &MyPlatResource{}
}

// IDCol ID 列
func (e *MyPlatResource) IDCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("id", func(rec *MyPlatResource) any {
		return &rec.ID
	})
}

// PidCol Pid 列
func (e *MyPlatResource) PidCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("pid", func(rec *MyPlatResource) any {
		return &rec.Pid
	})
}

// PlatIdCol PlatId 列
func (e *MyPlatResource) PlatIdCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("plat_id", func(rec *MyPlatResource) any {
		return &rec.PlatId
	})
}

// TitleCol Title 列
func (e *MyPlatResource) TitleCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("title", func(rec *MyPlatResource) any {
		return &rec.Title
	})
}

// NameCol Name 列
func (e *MyPlatResource) NameCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("name", func(rec *MyPlatResource) any {
		return &rec.Name
	})
}

// IconCol Icon 列
func (e *MyPlatResource) IconCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("icon", func(rec *MyPlatResource) any {
		return &rec.Icon
	})
}

// TypeCol Type 列
func (e *MyPlatResource) TypeCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("type", func(rec *MyPlatResource) any {
		return &rec.Type
	})
}

// PermissionCol Permission 列
func (e *MyPlatResource) PermissionCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("permission", func(rec *MyPlatResource) any {
		return &rec.Permission
	})
}

// PathCol Path 列
func (e *MyPlatResource) PathCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("path", func(rec *MyPlatResource) any {
		return &rec.Path
	})
}

// RouterCol Router 列
func (e *MyPlatResource) RouterCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("router", func(rec *MyPlatResource) any {
		return &rec.Router
	})
}

// SortCol Sort 列
func (e *MyPlatResource) SortCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("sort", func(rec *MyPlatResource) any {
		return &rec.Sort
	})
}

// CreateByCol CreateBy 列
func (e *MyPlatResource) CreateByCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("create_by", func(rec *MyPlatResource) any {
		return &rec.CreateBy
	})
}

// CreateAtCol CreateAt 列
func (e *MyPlatResource) CreateAtCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("create_at", func(rec *MyPlatResource) any {
		return &rec.CreateAt
	})
}

// ModifyByCol ModifyBy 列
func (e *MyPlatResource) ModifyByCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("modify_by", func(rec *MyPlatResource) any {
		return &rec.ModifyBy
	})
}

// ModifyAtCol ModifyAt 列
func (e *MyPlatResource) ModifyAtCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("modify_at", func(rec *MyPlatResource) any {
		return &rec.ModifyAt
	})
}

// DeletedCol Deleted 列
func (e *MyPlatResource) DeletedCol() *iris.Column[MyPlatResource] {
	return iris.WithColumn("deleted", func(rec *MyPlatResource) any {
		return &rec.Deleted
	})
}

// Configure evaluator 配置
func (e *MyPlatResource) Configure(fn func(*iris.Evaluator[MyPlatResource])) {
	if e.evaluator == nil {
		e.evaluator = iris.WithLogicalEvaluator[MyPlatResource]()
	}
	fn(e.evaluator)
}

// ColumnAndValue 列值计算
func (e *MyPlatResource) ColumnAndValue(fns ...func(*iris.Column[MyPlatResource], any) bool) (selfishs []iris.Selfish, values []any) {
	fn := func(*iris.Column[MyPlatResource], any) bool {
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
	if fn(e.TitleCol(), e.Title) {
		selfishs = append(selfishs, e.TitleCol())
		values = append(values, *e.Title)
	}
	if fn(e.NameCol(), e.Name) {
		selfishs = append(selfishs, e.NameCol())
		values = append(values, *e.Name)
	}
	if fn(e.IconCol(), e.Icon) {
		selfishs = append(selfishs, e.IconCol())
		values = append(values, *e.Icon)
	}
	if fn(e.TypeCol(), e.Type) {
		selfishs = append(selfishs, e.TypeCol())
		values = append(values, *e.Type)
	}
	if fn(e.PermissionCol(), e.Permission) {
		selfishs = append(selfishs, e.PermissionCol())
		values = append(values, *e.Permission)
	}
	if fn(e.PathCol(), e.Path) {
		selfishs = append(selfishs, e.PathCol())
		values = append(values, *e.Path)
	}
	if fn(e.RouterCol(), e.Router) {
		selfishs = append(selfishs, e.RouterCol())
		values = append(values, *e.Router)
	}
	if fn(e.SortCol(), e.Sort) {
		selfishs = append(selfishs, e.SortCol())
		values = append(values, *e.Sort)
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
func (e *MyPlatResource) Asterisk(fns ...func(string) string) []*iris.Column[MyPlatResource] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris.Column[MyPlatResource]{
		e.IDCol().Decorate(fn),
		e.PidCol().Decorate(fn),
		e.PlatIdCol().Decorate(fn),
		e.TitleCol().Decorate(fn),
		e.NameCol().Decorate(fn),
		e.IconCol().Decorate(fn),
		e.TypeCol().Decorate(fn),
		e.PermissionCol().Decorate(fn),
		e.PathCol().Decorate(fn),
		e.RouterCol().Decorate(fn),
		e.SortCol().Decorate(fn),
		e.CreateByCol().Decorate(fn),
		e.CreateAtCol().Decorate(fn),
		e.ModifyByCol().Decorate(fn),
		e.ModifyAtCol().Decorate(fn),
		e.DeletedCol().Decorate(fn),
	}
}

// PKey 主键
func (e *MyPlatResource) PKey() *iris.Column[MyPlatResource] {
	return e.IDCol()
}

// LogicDelKey 逻辑删除
func (e *MyPlatResource) LogicDelKey() *iris.Column[MyPlatResource] {
	return e.DeletedCol()
}

// Evaluator 计算器
func (e *MyPlatResource) Evaluator() *iris.Evaluator[MyPlatResource] {
	if e == nil {
		return nil
	}
	return e.evaluator
}

// Table 表
func (e *MyPlatResource) Table() *iris.RefTable {
	return iris.WithTable("my_plat_resource")
}

// Self 原始信息
func (e *MyPlatResource) Self() *MyPlatResource {
	return e
}
