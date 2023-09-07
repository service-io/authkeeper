// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

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
	return &RefTable{
		name: "account",
	}
}

func TestBaseEntity_QuerySQL(t *testing.T) {
	account := &Account{}
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Select(account.FID(), account.FName()).From(account.Table())

	sql, countSQL, values, mappers := bei.QuerySQL()
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", countSQL)
	fmt.Printf("rs:> %+v\n", values)
	fmt.Printf("rs:> %+v\n", mappers)
}

func TestBaseEntity_QuerySQL0(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Select(account.FID(), account.FName()).From(table).Where(Once(account.FName().Fd(), EQ, 1).And(Once(account.FID().Fd(), EQ, 2), Once(account.FAge().Fd(), EQ, 3)))

	sql, countSQL, values, mappers := bei.QuerySQL()
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", countSQL)
	fmt.Printf("rs:> %+v\n", values)
	fmt.Printf("rs:> %+v\n", mappers)
}

func TestBaseEntity_QuerySQL1(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Select(account.FID(), account.FName()).
		From(table).
		Where(Once(account.FName().Fd(), EQ, 1).And(Once(account.FID().Fd(), EQ, 2), Once(account.FAge().Fd(), EQ, 3))).GroupBy(bei.Group(account.FGender().Fd()))

	sql, countSQL, values, mappers := bei.QuerySQL()
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", countSQL)
	fmt.Printf("rs:> %+v\n", values)
	fmt.Printf("rs:> %+v\n", mappers)
}

func TestBaseEntity_QuerySQL2(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Select(account.FID(), account.FName()).
		From(table).
		Where(Once(account.FName().Fd(), EQ, 1).And(Once(account.FID().Fd(), EQ, 2), Once(account.FAge().Fd(), EQ, 3))).
		GroupBy(bei.Group(account.FGender().Fd())).
		OrderBy(bei.Asc(account.FBirthday().Fd()))

	sql, countSQL, values, mappers := bei.QuerySQL()
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", countSQL)
	fmt.Printf("rs:> %+v\n", values)
	fmt.Printf("rs:> %+v\n", mappers)
}

func TestBaseEntity_QuerySQL3(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Select(account.FID(), account.FName()).
		From(table).
		Where(Once(account.FName().Fd(), EQ, 1).And(Once(account.FID().Fd(), EQ, 2), Once(account.FAge().Fd(), EQ, 3))).
		GroupBy(bei.Group(account.FGender().Fd())).
		Having(Once(account.FName().Fd(), EQ, 1).And(Once(account.FID().Fd(), EQ, 2), Once(account.FAge().Fd(), EQ, 3))).
		OrderBy(bei.Asc(account.FBirthday().Fd())).
		WithLogicDeleted()

	sql, countSQL, values, mappers := bei.QuerySQL()
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", countSQL)
	fmt.Printf("rs:> %+v\n", values)
	fmt.Printf("rs:> %+v\n", mappers)
}

func TestBaseEntity_QuerySQL4(t *testing.T) {
	account := &Account{}
	table := account.Table().As("acc")
	user := &RefTable{
		name: "user",
		as:   "ur",
	}
	table.Ref(user.JoinType(Join).On(table.RefKey(account.FName().Fd()), EQ, user.RefKey("nickname")))
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	wrapFn := bei.WrapFn(table)
	bei.Select(wrapFn(account.FID()), wrapFn(account.FName())).
		From(table).
		Where(Once(wrapFn(account.FName()).Fd(), EQ, 1).And(Once(wrapFn(account.FID()).Fd(), EQ, 2), Once(wrapFn(account.FAge()).Fd(), EQ, 3))).
		GroupBy(bei.Group(wrapFn(account.FGender()).Fd())).
		Having(Once(wrapFn(account.FGender()).Fd(), EQ, 4).And(Once(wrapFn(account.FBirthday()).Fd(), EQ, time.Now()))).
		OrderBy(bei.Asc(wrapFn(account.FBirthday()).Fd())).
		WithLogicDeleted()

	sql, countSQL, values, mappers := bei.QuerySQL()
	fmt.Printf("rs:> %+v\n", bei.String())
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", countSQL)
	fmt.Printf("rs:> %+v\n", values)
	fmt.Printf("rs:> %+v\n", mappers)
}

func TestBaseEntity_Insert(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Insert(account.FID(), account.FName()).Into(table).WithValues(1, 2)

	sql, values := bei.InsertSQL()
	fmt.Printf("rs:> %+v\n", bei.String())
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", values)
}

func TestBaseEntity_Insert0(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Insert(account.FID(), account.FName()).Into(table).WithValues(1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4)

	sql, values := bei.InsertSQL()
	fmt.Printf("rs:> %+v\n", bei.String())
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", values)
}

func TestBaseEntity_DeleteSQL(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Delete().From(table).Where(Once("id", EQ, 1)).WithLogicDeleted()

	sql, values := bei.DeleteSQL()
	fmt.Printf("rs:> %+v\n", bei.String())
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", values)
}

func TestBaseEntity_UpdateSQL(t *testing.T) {
	account := &Account{}
	table := account.Table()
	bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
	bei.Update(table, account.FName(), account.FAge()).Where(Once("id", EQ, 1)).WithLogicDeleted().WithValues(2, 3)

	sql, values := bei.UpdateSQL()
	fmt.Printf("rs:> %+v\n", bei.String())
	fmt.Printf("rs:> %+v\n", sql)
	fmt.Printf("rs:> %+v\n", values)
}

func BenchmarkBaseEntity_QuerySQL(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		account := &Account{}
		table := account.Table().As("acc")
		user := &RefTable{
			name: "user",
			as:   "ur",
		}
		table.Ref(user.JoinType(Join).On(table.RefKey(account.FName().Fd()), EQ, user.RefKey("nickname")))
		bei := &BaseEntity[Account]{deletedKey: "deleted", deletedVal: "1", undeletedVal: "0"}
		wrapFn := bei.WrapFn(table)
		bei.Select(wrapFn(account.FID()), wrapFn(account.FName())).
			From(table).
			Where(Once(wrapFn(account.FName()).Fd(), EQ, 1).And(Once(wrapFn(account.FID()).Fd(), EQ, 2), Once(wrapFn(account.FAge()).Fd(), EQ, 3))).
			GroupBy(bei.Group(wrapFn(account.FGender()).Fd())).
			Having(Once(wrapFn(account.FGender()).Fd(), EQ, 4).And(Once(wrapFn(account.FBirthday()).Fd(), EQ, time.Now()))).
			OrderBy(bei.Asc(wrapFn(account.FBirthday()).Fd())).
			WithLogicDeleted()

		_, _, _, _ = bei.QuerySQL()
		// fmt.Printf("rs:> %+v\n", beiold.String())
		// fmt.Printf("rs:> %+v\n", sql)
		// fmt.Printf("rs:> %+v\n", countSQL)
		// fmt.Printf("rs:> %+v\n", values)
		// fmt.Printf("rs:> %+v\n", mappers)
	}
}
