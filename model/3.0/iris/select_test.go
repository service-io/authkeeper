package iris

import (
	"fmt"
	"metis/model/3.0/iris/internal/token"
	"testing"
)

func TestRefTable_Lit(t *testing.T) {
	eq := func(b bool) string { return "=" }
	a := WithTable("Axx")
	b := WithTable("Bxx").JoinType(LeftJoin)
	c := WithTable("Cxx").JoinType(LeftJoin).As("c")
	d := WithTable("Dxx").JoinType(LeftJoin).As("d")
	e := WithTable("Exx").JoinType(LeftJoin).As("e").On("name", eq, "nickname")
	e0 := WithTable("E0").JoinType(LeftJoin).On("name0", eq, "nickname0")
	f := WithTable("Ayy").Ref(e, e0)
	fmt.Printf("rs:> %+v\n", a.Literal())
	fmt.Printf("rs:> %+v\n", b.Literal())
	fmt.Printf("rs:> %+v\n", c.Literal())
	fmt.Printf("rs:> %+v\n", d.Literal())
	fmt.Printf("rs:> %+v\n", e.Literal())
	fmt.Printf("rs:> %+v\n", f.Literal())
}

func TestPredicate_Literal(t *testing.T) {
	name := WithColumn[string]("name", nil)
	age := WithColumn[string]("age", nil)
	gender := WithColumn[string]("gender", nil)
	bir := WithColumn[string]("bir", nil)
	a, av := WithPred(name, EQOpr, "tabuyos").Literal()
	b, bv := WithPred(name, EQOpr, "tabuyos").And(WithPred(age, EQOpr, 23)).Literal()
	c, cv := WithPred(name, EQOpr, "tabuyos").And(WithPred(age, EQOpr, 23), WithPred(gender, EQOpr, 1)).Literal()
	d, dv := WithPred(name, EQOpr, "tabuyos").And(WithPred(age, EQOpr, 23).Or(WithPred(gender, EQOpr, 1))).Literal()
	e, ev := WithPred(name, EQOpr, "tabuyos").
		And(
			WithPred(age, EQOpr, 23).
				Or(
					WithPred(gender, EQOpr, 1),
				),
		).
		Or(
			WithPred(bir, EQOpr, 121).And(
				WithPred(WithColumn[string]("uid0", nil), NQOpr, 32),
				WithPred(WithColumn[string]("uid1", nil), NQOpr, 32),
				WithPred(WithColumn[string]("uid2", nil), NQOpr, 32),
				WithPred(WithColumn[string]("uid3", nil), NQOpr, 32),
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
		_ = WithTable("Axx").Ref(
			WithTable("Ayy").On("name", eq, "nickname"), WithTable("Azz").On("name", eq, "nickname"),
		)
		// _ = WithTable("Bxx").JoinType(LeftJoin)
		// _ = WithTable("Cxx").JoinType(LeftJoin).As("c")
		// _ = WithTable("Dxx").JoinType(LeftJoin).As("d")
		// e := WithTable("Exx").JoinType(LeftJoin).As("e").On("name", eq, "nickname")
		// e0 := WithTable("E0").JoinType(LeftJoin).On("name0", eq, "nickname0")
		// _ = WithTable("Ayy").Ref(e, e0)
	}
}

func BenchmarkPredicate_Literal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WithPred(WithColumn[string]("name", nil), EQOpr, "tabuyos").Literal()
		// WithPred("name", EQOpr, "tabuyos").And(WithPred("age", EQOpr, 23)).Literal()
		// WithPred("name", EQOpr, "tabuyos").And(WithPred("age", EQOpr, 23), WithPred("gender", EQOpr, 1)).Literal()
		// WithPred("name", EQOpr, "tabuyos").And(WithPred("age", EQOpr, 23).Or(WithPred("gender", EQOpr, 1))).Literal()
		// WithPred("name", EQOpr, "tabuyos").
		//	And(
		//		WithPred("age", EQOpr, 23).
		//			Or(
		//				WithPred("gender", EQOpr, 1),
		//			),
		//	).
		//	Or(
		//		WithPred("bir", EQOpr, 121).And(
		//			WithPred("uid0", NQOpr, 32),
		//			WithPred("uid1", NQOpr, 32),
		//			WithPred("uid2", NQOpr, 32),
		//			WithPred("uid3", NQOpr, 32),
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
	token.LowerCase()
	// eval := Default[string]()
	eval := WithLogicalEvaluator[string]()
	user := WithTable("user")
	account := WithTable("account")
	tenant := WithTable("tenant")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)
	and0 := WithPred(ac, EQOpr, "25")
	and1 := WithPred(ic, EQOpr, "111000")
	or0 := WithPred(nc, EQOpr, "25")
	or1 := WithPred(ic, EQOpr, "111000")
	predicate := WithPred(nc, EQOpr, "tabuyos").And(and0, and1).Or(or0, or1)
	having := WithPred(ic, EQOpr, "111000")
	account.JoinType(LeftJoin).OnEQ(nc.Name(), ic.Name())
	tenant.JoinType(LeftJoin).OnEQ(ic.Name(), ac.Name())
	user.Ref(account, tenant)
	eval.Select(nc, ac, ic).Hint(token.Distinct).From(user).Where(predicate).GroupBy(nc, ac).
		Having(having).OrderBy(ac.Asc(), nc.Desc()).Limit(20).Offset(0).Eval()
	sql := eval.EvalInfo().SQL()
	totalSql := eval.EvalInfo().TotalSQL()
	values := eval.EvalInfo().Values()
	mappers := eval.EvalInfo().Mappers()
	fmt.Println(sql)
	fmt.Println(totalSql)
	fmt.Println(values)
	fmt.Println(mappers)
}

func BenchmarkEvaluator_Eval(b *testing.B) {
	token.LowerCase()
	b.ReportAllocs()
	b.ResetTimer()
	user := WithTable("user")
	account := WithTable("account")
	tenant := WithTable("tenant")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)
	for i := 0; i < b.N; i++ {
		eval := &Evaluator[string]{}
		and0 := WithPred(ac, EQOpr, "25")
		and1 := WithPred(ic, EQOpr, "111000")
		or0 := WithPred(ac, EQOpr, "25")
		or1 := WithPred(ic, EQOpr, "111000")
		predicate := WithPred(nc, EQOpr, "tabuyos").And(and0, and1).Or(or0, or1)
		having := WithPred(ic, EQOpr, "111000")
		account.JoinType(LeftJoin).OnEQ(nc.Name(), ic.Name())
		tenant.JoinType(LeftJoin).OnEQ(ic.Name(), ac.Name())
		user.Ref(account, tenant)
		eval.Select(nc, ac, ic).Hint(token.Distinct).From(user).Where(predicate).GroupBy(nc, ac).
			Having(having).OrderBy(ac.Asc(), nc.Desc()).Limit(20).Offset(0).Eval()
	}
}

func TestPredicate_Render(t *testing.T) {
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)
	and0 := WithPred(ac, EQOpr, "25")
	and1 := WithPred(ic, EQOpr, "111000")
	or0 := WithPred(nc, EQOpr, "25")
	or1 := WithPred(ic, EQOpr, "111000")
	predicate := WithPred(nc, EQOpr, "tabuyos").And(and0, and1).Or(or0, or1)

	builder, release := getBuilder()
	defer release()
	predicate.Render(builder)
	fmt.Println(builder.String())
}

type lks struct {
	m map[string]*EvalInfo[string]
}

func (l *lks) Put(key string, info *EvalInfo[string]) {
	l.m[key] = info
}

func (l *lks) Get(key string) *EvalInfo[string] {
	return l.m[key]
}

func TestSelectEvaluator_Persist(t *testing.T) {
	sqlMap := map[string]*EvalInfo[string]{}
	memoryPersist := NewMemoryPersist[string](&lks{m: sqlMap})
	memoryPersist.Persistence(
		"one", WithEvalInfo[string](
			"select distinct name, age, id from user left join account on name = id left join tenant on id = age where (name = ? and age = ? and id = ?) or name = ? or id = ? group by name, age having id = ? order by age, name desc limit = ? offset = ?;",
			"", nil, nil,
		),
	)

	token.LowerCase()
	// eval := Default[string]()
	eval := WithLogicalEvaluator[string]()
	user := WithTable("user")
	account := WithTable("account").As("acc")
	tenant := WithTable("tenant")
	nc := WithColumn[string]("name", nil).Decorate(account.Decorate).As("acc_name")
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)
	and0 := WithPred(ac, EQOpr, "25")
	and1 := WithPred(ic, EQOpr, "111000")
	or0 := WithPred(nc, EQOpr, "25")
	or1 := WithPred(ic, EQOpr, "111000")
	predicate := WithPred(nc, EQOpr, "tabuyos").And(and0, and1).Or(or0, or1)
	having := WithPred(ic, EQOpr, "111000")
	account.JoinType(LeftJoin).OnEQ(nc.Name(), ic.Name())
	tenant.JoinType(LeftJoin).OnEQ(ic.Name(), ac.Name())
	user.Ref(account, tenant)
	eval.Select(nc, ac, ic).Hint(token.Distinct).From(user).Where(predicate).GroupBy(nc, ac).
		Having(having).OrderBy(ac.Asc(), nc.Desc()).Limit(20).Offset(0).Eval()
	sql := eval.EvalInfo().SQL()
	totalSql := eval.EvalInfo().TotalSQL()
	values := eval.EvalInfo().Values()
	mappers := eval.EvalInfo().Mappers()
	fmt.Println(sql)
	fmt.Println(totalSql)
	fmt.Println(values)
	fmt.Println(mappers)
}

func BenchmarkSelectEvaluator_Persist(b *testing.B) {
	sqlMap := map[string]*EvalInfo[string]{}
	memoryPersist := NewMemoryPersist[string](&lks{m: sqlMap})

	token.LowerCase()
	// eval := Default[string]()
	user := WithTable("user")
	account := WithTable("account")
	tenant := WithTable("tenant")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eval := WithLogicalEvaluator[string]()
		and0 := WithPred(ac, EQOpr, "25")
		and1 := WithPred(ic, EQOpr, "111000")
		or0 := WithPred(nc, EQOpr, "25")
		or1 := WithPred(ic, EQOpr, "111000")
		predicate := WithPred(nc, EQOpr, "tabuyos").And(and0, and1).Or(or0, or1)
		having := WithPred(ic, EQOpr, "111000")
		account.JoinType(LeftJoin).OnEQ(nc.Name(), ic.Name())
		tenant.JoinType(LeftJoin).OnEQ(ic.Name(), ac.Name())
		user.Ref(account, tenant)
		eval.Cache("one", memoryPersist).Select(nc, ac, ic).Hint(token.Distinct).From(user).Where(predicate).GroupBy(nc, ac).Having(having).OrderBy(ac.Asc(), nc.Desc()).Limit(20).Offset(0).Eval()
	}
}
