// Package autogen
// @author tabuyos
// @since 2023/9/7
// @description autogen
package bei

import (
	"fmt"
	"testing"
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

func (e *Account) Pkey() *FD[Account] {
	return e.FID()
}

func (e *Account) FID() *FD[Account] {
	return OfFD(
		"id", func(r *Account) any {
			return &r.ID
		},
	)
}

func (e *Account) FName() *FD[Account] {
	return OfFD(
		"name", func(r *Account) any {
			return &r.Name
		},
	)
}

func (e *Account) FAge() *FD[Account] {
	return OfFD(
		"age", func(r *Account) any {
			return &r.Age
		},
	)
}

func (e *Account) FGender() *FD[Account] {
	return OfFD(
		"gender", func(r *Account) any {
			return &r.Gender
		},
	)
}

func (e *Account) FBirthday() *FD[Account] {
	return OfFD(
		"birthday", func(r *Account) any {
			return &r.Birthday
		},
	)
}

func (e *Account) Table() *RefTable {
	return OfRef("account")
}

func TestBEI(t *testing.T) {
	account := &Account{}
	dft := OfLogicDeletedDefault[Account]()

	persist := dft.Select(account.FName()).From(account.Table()).Where(account.FName().Eq(123)).Limit(10).Persist()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.TotalSQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
	fmt.Printf("rs:> %+v\n", persist.Mappers())
}

func TestBEI_0(t *testing.T) {
	account := &Account{}
	dft := OfLogicDeletedDefault[Account]()

	persist := dft.Update(account.Table()).Set(account.FName(), 1).Where(account.FName().LikeRight(4321)).Persist()
	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_1(t *testing.T) {
	account := &Account{}
	dft := OfLogicDeletedDefault[Account]()

	persist := dft.Delete().From(account.Table()).Where(account.FID().Eq(4321)).Persist()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_2(t *testing.T) {
	account := &Account{}
	dft := OfLogicDeletedDefault[Account]()

	persist := dft.Insert(account.FName(), account.FAge()).Values(1, 2, 1, 2).Persist()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_3(t *testing.T) {
	account := &Account{}
	dft := OfLogicDeletedDefault[Account]()

	persist := dft.Insert(account.FName(), account.FAge()).Into(account.Table()).Value(1, 2).Persist()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_4(t *testing.T) {
	account := &Account{}
	dft := OfLogicDeletedDefault[Account]()

	persist := dft.Select(account.FName()).From(account.Table()).Where(account.FName().Eq(123)).Limit(10).GroupBy(dft.Group(account.FName().Fd())).OrderBy(dft.Asc("nk")).Persist()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.TotalSQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
	fmt.Printf("rs:> %+v\n", persist.Mappers())
}

func TestBEI_5(t *testing.T) {
	account := &Account{}
	dft := OfLogicDeletedDefault[Account]()

	user := OfRef("user").As("usr")
	acco := OfRef("account").As("acc").JoinType(LeftJoin).On(dft.OfFD("a", nil).Inject(user).Fd(), eqSym, "b")

	persist := dft.Select(account.FName().Inject(acco)).From(user.Ref(acco)).
		Where(
			account.FName().Inject(acco).Eq("张三").
				And(account.FAge().Inject(acco).Eq(33).
					Or(account.FBirthday().Inject(acco).Eq("2023-10-10"),
						dft.OfFD("year", nil).Inject(user).Eq("2023"))),
		).Limit(10).
		GroupBy(dft.Group(account.FName().Fd())).OrderBy(dft.Asc("nk")).Persist()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.TotalSQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
	fmt.Printf("rs:> %+v\n", persist.Mappers())
}
