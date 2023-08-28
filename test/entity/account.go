package entity

import (
	"time"
)

type Account struct {
	ID       *int64     `json:"id"`
	Name     *string    `json:"name"`
	Age      *uint16    `json:"age"`
	Gender   *string    `json:"gender"`
	Birthday *time.Time `json:"birthday"`

	bei *BaseEntity[Account]
}

func (e *Account) Configure(configFn func(entity *BaseEntity[Account])) {
	if e.bei == nil {
		e.bei = &BaseEntity[Account]{table: "account"}
	}
	configFn(e.bei)
}

func (e *Account) BEI() *BaseEntity[Account] {
	if e.bei == nil {
		e.bei = &BaseEntity[Account]{table: "account"}
	}
	return e.bei
}

func (e *Account) Pkey() *FS[Account] {
	return e.FID()
}

func (e *Account) FID() *FS[Account] {
	return OfFS(
		"id", func(r *Account) any {
			return &r.ID
		},
	)
}

func (e *Account) FName() *FS[Account] {
	return OfFS(
		"name", func(r *Account) any {
			return &r.Name
		},
	)
}

func (e *Account) FAge() *FS[Account] {
	return OfFS(
		"age", func(r *Account) any {
			return &r.Age
		},
	)
}

func (e *Account) FGender() *FS[Account] {
	return OfFS(
		"gender", func(r *Account) any {
			return &r.Gender
		},
	)
}

func (e *Account) FBirthday() *FS[Account] {
	return OfFS(
		"birthday", func(r *Account) any {
			return &r.Birthday
		},
	)
}
