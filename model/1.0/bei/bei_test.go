// Package autogen
// @author tabuyos
// @since 2023/9/7
// @description autogen
package bei

import (
	"fmt"
	"metis/model/1.0/bei/keyword"
	"testing"
	"time"
)

type Account struct {
	ID       *int64     `json:"id"`
	Name     *string    `json:"name"`
	Age      *uint16    `json:"age"`
	Gender   *string    `json:"gender"`
	Birthday *time.Time `json:"birthday"`

	bei   *Evaluator[Account]
	table *RefTable
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
	if e.table == nil {
		e.table = OfRef("account")
		return e.table
	}
	return e.table
}

func TestBEI(t *testing.T) {
	account := &Account{}
	dft := WithLogical[Account]()

	persist := dft.Select(account.FName()).From(account.Table()).Where(account.FName().Eq(123)).Limit(10).EvalInfo()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.TotalSQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
	fmt.Printf("rs:> %+v\n", persist.Mappers())
}

func TestBEI_0(t *testing.T) {
	account := &Account{}
	dft := WithLogical[Account]()

	persist := dft.Update(account.Table()).Set(account.FName(), 1).Where(account.FName().LikeRight(4321)).EvalInfo()
	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_1(t *testing.T) {
	account := &Account{}
	dft := WithLogical[Account]()

	persist := dft.Delete().From(account.Table()).Where(account.FID().Eq(4321)).EvalInfo()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_2(t *testing.T) {
	account := &Account{}
	dft := WithLogical[Account]()

	persist := dft.Insert(account.FName(), account.FAge()).Values(1, 2, 1, 2).EvalInfo()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_3(t *testing.T) {
	account := &Account{}
	dft := WithLogical[Account]()

	persist := dft.Insert(account.FName(), account.FAge()).Into(account.Table()).Value(1, 2).EvalInfo()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
}

func TestBEI_4(t *testing.T) {
	account := &Account{}
	dft := WithLogical[Account]()

	persist := dft.Select(account.FName()).From(account.Table()).Where(account.FName().Eq(123)).Limit(10).GroupBy(dft.Group(account.FName().Fd())).OrderBy(dft.Asc("nk")).EvalInfo()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.TotalSQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
	fmt.Printf("rs:> %+v\n", persist.Mappers())
}

func TestPredicate_And(t *testing.T) {
	account := &Account{}
	predicate := account.FName().Eq("张三").And(account.FAge().Eq(33).Or(account.FBirthday().Eq("2023-10-10")))

	sql, _ := predicate.SQL()

	fmt.Println(sql)
}

func BenchmarkPredicate_And(b *testing.B) {
	account := &Account{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = account.FName().Eq("张三").And(account.FAge().Eq(33).Or(account.FBirthday().Eq("2023-10-10")))
	}
}

func TestBEI_5(t *testing.T) {
	keyword.RegistryCase(false)
	account := &Account{}
	dft := WithLogical[Account]()

	user := OfRef("user").As("usr")
	acco := OfRef("account").As("acc").JoinType(LeftJoin).On(dft.OfFD("a", nil).Inject(user).Fd(), eqSym, "b")

	persist := dft.Select(account.FName().Inject(acco)).From(user.Ref(acco)).
		Where(
			account.FName().Inject(acco).Eq("张三").
				And(
					account.FAge().Inject(acco).Eq(33).
						Or(
							account.FBirthday().Inject(acco).Eq("2023-10-10"),
							dft.OfFD("year", nil).Inject(user).Eq("2023"),
						),
				),
		).Limit(10).
		GroupBy(dft.Group(account.FName().Fd())).OrderBy(dft.Asc("nk")).EvalInfo()

	fmt.Printf("rs:> %+v\n", dft.String())
	fmt.Printf("rs:> %+v\n", persist.SQL())
	fmt.Printf("rs:> %+v\n", persist.TotalSQL())
	fmt.Printf("rs:> %+v\n", persist.Values())
	fmt.Printf("rs:> %+v\n", persist.Mappers())
}

func BenchmarkBEI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		account := &Account{}
		dft := WithLogical[Account]()
		dft.Select(account.FName(), account.FAge(), account.FID()).From(account.Table()).EvalInfo()

		// user := OfRef("user").As("usr")
		// acco := OfRef("account").As("acc").JoinType(LeftJoin).On(dft.OfFD("a", nil).Inject(user).Fd(), eqSym, "b")
		// // user.Ref(acco)
		//
		// _ = dft.Select(account.FName().Inject(acco)).
		//   From(user.Ref(acco)).
		//   // Where(
		//   // account.FName().Inject(acco).Eq("张三"),
		//   // And(
		//   //   // account.FAge().Inject(acco).Eq(33).
		//   //   // Or(account.FBirthday().Inject(acco).Eq("2023-10-10"),
		//   //   // dft.OfFD("year", nil).Inject(user).Eq("2023"),
		//   //   // ),
		//   // ),
		//   // ).
		//   // Limit(10).
		//   // GroupBy(dft.Group(account.FName().Fd())).
		//   // OrderBy(dft.Asc("nk")).
		//   EvalInfo()
		//
		// // fmt.Printf("rs:> %+v\n", dft.String())
	}
}

func BenchmarkQueryBuilder_Eval(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		account := &Account{}
		dft := WithLogical[Account]()
		dft.Select(account.FName(), account.FAge(), account.FID(), account.FBirthday(), account.FGender()).
			From(account.Table()).
			Where(account.FID().Eq(111000).And(account.FName().Like("tabuyos"), account.FAge().Ge(25))).
			WithSQLKey("one").Eval()
	}
}

func BenchmarkPredicate_SQL(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Once("name", eqSym, "tabuyos").And(Once("age", eqSym, 23), Once("gender", eqSym, 1)).SQL()
	}
}

func BenchmarkEvaluator_Eval(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	user := OfRef("user")
	account := OfRef("account")
	tenant := OfRef("tenant")
	nc := OfFD[string]("name", nil)
	ac := OfFD[string]("age", nil)
	ic := OfFD[string]("id", nil)
	for i := 0; i < b.N; i++ {
		eval := &Evaluator[string]{}
		and0 := Once(ac.Literal(), eqSym, "25")
		and1 := Once(ic.Literal(), eqSym, "111000")
		or0 := Once(ac.Literal(), eqSym, "25")
		or1 := Once(ic.Literal(), eqSym, "111000")
		predicate := Once(nc.Literal(), eqSym, "tabuyos").And(and0, and1).Or(or0, or1)
		having := Once(ic.Literal(), eqSym, "111000")
		account.JoinType(LeftJoin).On(nc.Literal(), eqSym, ic.Literal())
		tenant.JoinType(LeftJoin).On(ic.Literal(), eqSym, ac.Literal())
		user.Ref(account, tenant)
		eval.Select(nc, ac, ic).Hint(keyword.Distinct).From(user).Where(predicate).GroupBy(eval.Group(nc.Fd(), ac.Fd())).
			Having(having).OrderBy(eval.Asc(ac.Fd()), eval.Desc(nc.Fd())).Limit(20).Offset(0).Eval()
		eval.Select(nc, ac, ic).Hint(keyword.Distinct).From(user).Where(predicate).Eval()
	}
}
