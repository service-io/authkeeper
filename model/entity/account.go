// Package entity
// @author tabuyos
// @since 2023/9/8
// @description entity
package entity

import (
	bei2 "metis/model/1.0/bei"
	"time"
)

type Account struct {
	ID       *int64     `json:"id"`
	Name     *string    `json:"name"`
	Age      *uint16    `json:"age"`
	Gender   *string    `json:"gender"`
	Deleted  *int8      `json:"deleted"`
	Birthday *time.Time `json:"birthday"`

	evaluator *bei2.Evaluator[Account]
}

func NewAccount() *Account {
	return &Account{}
}

func (e *Account) FID() *bei2.FD[Account] {
	return bei2.OfFD(
		"id", func(r *Account) any {
			return &r.ID
		},
	)
}

func (e *Account) FName() *bei2.FD[Account] {
	return bei2.OfFD(
		"name", func(r *Account) any {
			return &r.Name
		},
	)
}

func (e *Account) FAge() *bei2.FD[Account] {
	return bei2.OfFD(
		"age", func(r *Account) any {
			return &r.Age
		},
	)
}

func (e *Account) FGender() *bei2.FD[Account] {
	return bei2.OfFD(
		"gender", func(r *Account) any {
			return &r.Gender
		},
	)
}

func (e *Account) FBirthday() *bei2.FD[Account] {
	return bei2.OfFD(
		"birthday", func(r *Account) any {
			return &r.Birthday
		},
	)
}

func (e *Account) FDeleted() *bei2.FD[Account] {
	return bei2.OfFD(
		"deleted", func(r *Account) any {
			return &r.Deleted
		},
	)
}

func (e *Account) Configure(fn func(*bei2.Evaluator[Account])) {
	if e.evaluator == nil {
		e.evaluator = bei2.WithLogical[Account]()
	}
	fn(e.evaluator)
}

func (e *Account) Asterisk() []*bei2.FD[Account] {
	return []*bei2.FD[Account]{e.FID(), e.FName(), e.FAge(), e.FBirthday(), e.FGender(), e.FDeleted()}
}

func (e *Account) PKey() *bei2.FD[Account] {
	return e.FID()
}

func (e *Account) LogicDelKey() *bei2.FD[Account] {
	return e.FDeleted()
}

func (e *Account) Evaluator() bei2.EvalInfoService[Account] {
	return e.evaluator
}

func (e *Account) Table() *bei2.RefTable {
	return bei2.OfRef("account")
}

func (e *Account) Self() *Account {
	return e
}
