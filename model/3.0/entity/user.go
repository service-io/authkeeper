// Package entity
// @author tabuyos
// @since 2023/9/12
// @description entity
package entity

import (
	"metis/model/3.0/iris"
	"time"
)

type User struct {
	// 主键ID
	ID *int64 `json:"id"`
	// 用户名
	Name *string `json:"name"`
	// 昵称
	NickName *string `json:"nickName"`
	// 真实姓名
	RealName *string `json:"realName"`
	// 证件号
	IdCard *string `json:"idCard"`
	// 性别
	Sex *int `json:"sex"`
	// 电话
	Mobile *string `json:"mobile"`
	// 邮箱
	Email *string `json:"email"`
	// 住址
	Address *string `json:"address"`
	// 备注
	Remark *string `json:"remark"`
	// 用户状态 0-停用;1-启用
	Status *int `json:"status"`
	// 实名状态
	SignStatus *int `json:"signStatus"`
	// 来源
	Source *string `json:"source"`
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

	evaluator *iris.Evaluator[User]
}

// NewUser 初始化
func NewUser() *User {
	return &User{}
}

func (e *User) IDCol() *iris.Column[User] {
	return iris.WithColumn("id", func(rec *User) any {
		return &rec.ID
	})
}
func (e *User) NameCol() *iris.Column[User] {
	return iris.WithColumn("name", func(rec *User) any {
		return &rec.Name
	})
}
func (e *User) NickNameCol() *iris.Column[User] {
	return iris.WithColumn("nick_name", func(rec *User) any {
		return &rec.NickName
	})
}
func (e *User) RealNameCol() *iris.Column[User] {
	return iris.WithColumn("real_name", func(rec *User) any {
		return &rec.RealName
	})
}
func (e *User) IdCardCol() *iris.Column[User] {
	return iris.WithColumn("id_card", func(rec *User) any {
		return &rec.IdCard
	})
}
func (e *User) SexCol() *iris.Column[User] {
	return iris.WithColumn("sex", func(rec *User) any {
		return &rec.Sex
	})
}
func (e *User) MobileCol() *iris.Column[User] {
	return iris.WithColumn("mobile", func(rec *User) any {
		return &rec.Mobile
	})
}
func (e *User) EmailCol() *iris.Column[User] {
	return iris.WithColumn("email", func(rec *User) any {
		return &rec.Email
	})
}
func (e *User) AddressCol() *iris.Column[User] {
	return iris.WithColumn("address", func(rec *User) any {
		return &rec.Address
	})
}
func (e *User) RemarkCol() *iris.Column[User] {
	return iris.WithColumn("remark", func(rec *User) any {
		return &rec.Remark
	})
}
func (e *User) StatusCol() *iris.Column[User] {
	return iris.WithColumn("status", func(rec *User) any {
		return &rec.Status
	})
}
func (e *User) SignStatusCol() *iris.Column[User] {
	return iris.WithColumn("sign_status", func(rec *User) any {
		return &rec.SignStatus
	})
}
func (e *User) SourceCol() *iris.Column[User] {
	return iris.WithColumn("source", func(rec *User) any {
		return &rec.Source
	})
}
func (e *User) CreateByCol() *iris.Column[User] {
	return iris.WithColumn("create_by", func(rec *User) any {
		return &rec.CreateBy
	})
}
func (e *User) CreateAtCol() *iris.Column[User] {
	return iris.WithColumn("create_at", func(rec *User) any {
		return &rec.CreateAt
	})
}
func (e *User) ModifyByCol() *iris.Column[User] {
	return iris.WithColumn("modify_by", func(rec *User) any {
		return &rec.ModifyBy
	})
}
func (e *User) ModifyAtCol() *iris.Column[User] {
	return iris.WithColumn("modify_at", func(rec *User) any {
		return &rec.ModifyAt
	})
}
func (e *User) DeletedCol() *iris.Column[User] {
	return iris.WithColumn("deleted", func(rec *User) any {
		return &rec.Deleted
	})
}

func (e *User) Configure(fn func(*iris.Evaluator[User])) {
	if e.evaluator == nil {
		e.evaluator = iris.WithLogicalEvaluator[User]()
	}
	fn(e.evaluator)
}

func (e *User) Asterisk(fns ...func(string) string) []*iris.Column[User] {
	var fn func(string) string
	if len(fns) > 0 {
		fn = fns[0]
	}
	return []*iris.Column[User]{
		e.IDCol().Decorate(fn),
		e.NameCol().Decorate(fn),
		e.NickNameCol().Decorate(fn),
		e.RealNameCol().Decorate(fn),
		e.IdCardCol().Decorate(fn),
		e.SexCol().Decorate(fn),
		e.MobileCol().Decorate(fn),
		e.EmailCol().Decorate(fn),
		e.AddressCol().Decorate(fn),
		e.RemarkCol().Decorate(fn),
		e.StatusCol().Decorate(fn),
		e.SignStatusCol().Decorate(fn),
		e.SourceCol().Decorate(fn),
		e.CreateByCol().Decorate(fn),
		e.CreateAtCol().Decorate(fn),
		e.ModifyByCol().Decorate(fn),
		e.ModifyAtCol().Decorate(fn),
		e.DeletedCol().Decorate(fn),
	}
}

func (e *User) PKey() *iris.Column[User] {
	return e.IDCol()
}

func (e *User) LogicDelKey() *iris.Column[User] {
	return e.DeletedCol()
}

func (e *User) Evaluator() iris.EvalInfoService[User] {
	return e.evaluator
}

func (e *User) Table() *iris.RefTable {
	return iris.WithTable("account_role")
}

func (e *User) Self() *User {
	return e
}
