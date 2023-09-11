package iris

import (
	"fmt"
	"metis/model/3.0/iris/internal/keyword"
	"testing"
)

func TestRefTable_Lit(t *testing.T) {
	eq := func(b bool) string { return "=" }
	a := WithRefTable("Axx")
	b := WithRefTable("Bxx").JoinType(LeftJoin)
	c := WithRefTable("Cxx").JoinType(LeftJoin).As("c")
	d := WithRefTable("Dxx").JoinType(LeftJoin).As("d")
	e := WithRefTable("Exx").JoinType(LeftJoin).As("e").On("name", eq, "nickname")
	e0 := WithRefTable("E0").JoinType(LeftJoin).On("name0", eq, "nickname0")
	f := WithRefTable("Ayy").Ref(e, e0)
	fmt.Printf("rs:> %+v\n", a.Literal())
	fmt.Printf("rs:> %+v\n", b.Literal())
	fmt.Printf("rs:> %+v\n", c.Literal())
	fmt.Printf("rs:> %+v\n", d.Literal())
	fmt.Printf("rs:> %+v\n", e.Literal())
	fmt.Printf("rs:> %+v\n", f.Literal())
}

func TestPredicate_Literal(t *testing.T) {
	a, av := Once("name", EqOp, "tabuyos").Literal()
	b, bv := Once("name", EqOp, "tabuyos").And(Once("age", EqOp, 23)).Literal()
	c, cv := Once("name", EqOp, "tabuyos").And(Once("age", EqOp, 23), Once("gender", EqOp, 1)).Literal()
	d, dv := Once("name", EqOp, "tabuyos").And(Once("age", EqOp, 23).Or(Once("gender", EqOp, 1))).Literal()
	e, ev := Once("name", EqOp, "tabuyos").
		And(
			Once("age", EqOp, 23).
				Or(
					Once("gender", EqOp, 1),
				),
		).
		Or(
			Once("bir", EqOp, 121).And(
				Once("uid0", NqOp, 32),
				Once("uid1", NqOp, 32),
				Once("uid2", NqOp, 32),
				Once("uid3", NqOp, 32),
			),
		).Literal()

	fmt.Printf("rs:> %+v\n  %+v\n", a, av)
	fmt.Printf("rs:> %+v\n  %+v\n", b, bv)
	fmt.Printf("rs:> %+v\n  %+v\n", c, cv)
	fmt.Printf("rs:> %+v\n  %+v\n", d, dv)
	fmt.Printf("rs:> %+v\n  %+v\n", e, ev)
}

func BenchmarkRefTable_Literal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eq := func(b bool) string { return "=" }
		_ = WithRefTable("Axx").Ref(
			WithRefTable("Ayy").On("name", eq, "nickname"), WithRefTable("Azz").On("name", eq, "nickname"),
		)
		//_ = WithRefTable("Bxx").JoinType(LeftJoin)
		//_ = WithRefTable("Cxx").JoinType(LeftJoin).As("c")
		//_ = WithRefTable("Dxx").JoinType(LeftJoin).As("d")
		//e := WithRefTable("Exx").JoinType(LeftJoin).As("e").On("name", eq, "nickname")
		//e0 := WithRefTable("E0").JoinType(LeftJoin).On("name0", eq, "nickname0")
		//_ = WithRefTable("Ayy").Ref(e, e0)
	}
}

func BenchmarkPredicate_Literal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Once("name", EqOp, "tabuyos").Literal()
		//Once("name", EqOp, "tabuyos").And(Once("age", EqOp, 23)).Literal()
		//Once("name", EqOp, "tabuyos").And(Once("age", EqOp, 23), Once("gender", EqOp, 1)).Literal()
		//Once("name", EqOp, "tabuyos").And(Once("age", EqOp, 23).Or(Once("gender", EqOp, 1))).Literal()
		//Once("name", EqOp, "tabuyos").
		//	And(
		//		Once("age", EqOp, 23).
		//			Or(
		//				Once("gender", EqOp, 1),
		//			),
		//	).
		//	Or(
		//		Once("bir", EqOp, 121).And(
		//			Once("uid0", NqOp, 32),
		//			Once("uid1", NqOp, 32),
		//			Once("uid2", NqOp, 32),
		//			Once("uid3", NqOp, 32),
		//		),
		//	).Literal()
	}
}

func BenchmarkAppendNil(b *testing.B) {
	var values = make([]string, 0)
	var name []string = nil
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		values = append(values, name...)
	}
}

func TestSelectEvaluator_Eval(t *testing.T) {
	keyword.RegistryCase(true)
	eval := &Evaluator[string]{}
	user := WithRefTable("user")
	account := WithRefTable("account")
	tenant := WithRefTable("tenant")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)
	and0 := WithPredicate(ac.Literal(), EqOp, "25")
	and1 := WithPredicate(ic.Literal(), EqOp, "111000")
	or0 := WithPredicate(ac.Literal(), EqOp, "25")
	or1 := WithPredicate(ic.Literal(), EqOp, "111000")
	predicate := WithPredicate(nc.Literal(), EqOp, "tabuyos").And(and0, and1).Or(or0, or1)
	having := WithPredicate(ic.Literal(), EqOp, "111000")
	account.JoinType(LeftJoin).OnEQ(nc.Name(), ic.Name())
	tenant.JoinType(LeftJoin).OnEQ(ic.Name(), ac.Name())
	user.Ref(account, tenant)
	eval.Select(nc, ac, ic).Hint(keyword.Distinct).From(user).Where(predicate).GroupBy(nc, ac).
		Having(having).OrderBy(ac.Asc(), nc.Desc()).Limit(20).Offset(0).Eval()
}

func BenchmarkEvaluator_Eval(b *testing.B) {
	keyword.RegistryCase(true)
	b.ReportAllocs()
	b.ResetTimer()
	user := WithRefTable("user")
	account := WithRefTable("account")
	tenant := WithRefTable("tenant")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)
	for i := 0; i < b.N; i++ {
		eval := &Evaluator[string]{}
		and0 := WithPredicate(ac.Literal(), EqOp, "25")
		and1 := WithPredicate(ic.Literal(), EqOp, "111000")
		or0 := WithPredicate(ac.Literal(), EqOp, "25")
		or1 := WithPredicate(ic.Literal(), EqOp, "111000")
		predicate := WithPredicate(nc.Literal(), EqOp, "tabuyos").And(and0, and1).Or(or0, or1)
		having := WithPredicate(ic.Literal(), EqOp, "111000")
		account.JoinType(LeftJoin).OnEQ(nc.Name(), ic.Name())
		tenant.JoinType(LeftJoin).OnEQ(ic.Name(), ac.Name())
		user.Ref(account, tenant)
		eval.Select(nc, ac, ic).Hint(keyword.Distinct).From(user).Where(predicate).GroupBy(nc, ac).
			Having(having).OrderBy(ac.Asc(), nc.Desc()).Limit(20).Offset(0).Eval()
	}
}

func BenchmarkBuilder(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}
}
