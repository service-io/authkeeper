// Package entity
// @author tabuyos
// @since 2023/9/8
// @description entity
package entity

import (
	"metis/old/model/1.0/bei"
	"time"
)

type Account struct {
	ID       *int64     `json:"id"`
	Name     *string    `json:"name"`
	Age      *uint16    `json:"age"`
	Gender   *string    `json:"gender"`
	Deleted  *int8      `json:"deleted"`
	Birthday *time.Time `json:"birthday"`

	evaluator *bei.Evaluator[Account]
}

func NewAccount() *Account {
	return &Account{}
}

func (e *Account) FID() *bei.FD[Account] {
	return bei.OfFD(
		"id", func(r *Account) any {
			return &r.ID
		},
	)
}

func (e *Account) FName() *bei.FD[Account] {
	return bei.OfFD(
		"name", func(r *Account) any {
			return &r.Name
		},
	)
}

func (e *Account) FAge() *bei.FD[Account] {
	return bei.OfFD(
		"age", func(r *Account) any {
			return &r.Age
		},
	)
}

func (e *Account) FGender() *bei.FD[Account] {
	return bei.OfFD(
		"gender", func(r *Account) any {
			return &r.Gender
		},
	)
}

func (e *Account) FBirthday() *bei.FD[Account] {
	return bei.OfFD(
		"birthday", func(r *Account) any {
			return &r.Birthday
		},
	)
}

func (e *Account) FDeleted() *bei.FD[Account] {
	return bei.OfFD(
		"deleted", func(r *Account) any {
			return &r.Deleted
		},
	)
}

func (e *Account) Configure(fn func(*bei.Evaluator[Account])) {
	if e.evaluator == nil {
		e.evaluator = bei.WithLogical[Account]()
	}
	fn(e.evaluator)
}

func (e *Account) Asterisk() []*bei.FD[Account] {
	return []*bei.FD[Account]{e.FID(), e.FName(), e.FAge(), e.FBirthday(), e.FGender(), e.FDeleted()}
}

func (e *Account) PKey() *bei.FD[Account] {
	return e.FID()
}

func (e *Account) LogicDelKey() *bei.FD[Account] {
	return e.FDeleted()
}

func (e *Account) Evaluator() bei.EvalInfoService[Account] {
	return e.evaluator
}

func (e *Account) Table() *bei.RefTable {
	return bei.OfRef("account")
}

func (e *Account) Self() *Account {
	return e
}
