// Package iris
// @author tabuyos
// @since 2023/9/11
// @description iris
package iris

import (
	"deepsea/old/model/3.0/iris/internal/token"
	"fmt"
	"strings"
	"testing"
)

func TestDeleteEvaluator_Eval(t *testing.T) {
	token.LowerCase()
	// eval := Default[string]()
	eval := WithLogicalEvaluator[string]()
	user := WithTable("user")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)

	eval.Delete().From(user).Where(nc.EQ(12).Or(ac.EQ(24).And(ic.EQ("111000")))).Eval()
	info := eval.EvalInfo()

	fmt.Printf("rs:> %+v\n", info.SQL())
	fmt.Printf("rs:> %+v\n", info.TotalSQL())
	fmt.Printf("rs:> %+v\n", info.Values())
	fmt.Printf("rs:> %+v\n", info.Mappers())
}

func TestPredicate_And(t *testing.T) {
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)

	predicate := nc.EQ(12).And(ac.EQ(24).Or(ic.EQ("111000"))).Or(ac.EQ(24), ic.EQ("111000"))

	sql, _ := predicate.Literal()

	fmt.Println(predicate.mod)
	fmt.Println(predicate.Mixed())
	fmt.Println(sql)
}

func BenchmarkWithPred(b *testing.B) {
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = nc.EQ(12).And(ac.EQ(24).Or(ic.EQ("111000")))
	}
}

func BenchmarkDeleteEvaluator_Eval(b *testing.B) {
	sqlMap := map[string]*EvalInfo[string]{}
	memoryPersist := NewMemoryPersist[string](&lks{m: sqlMap})

	user := WithTable("user")
	nc := WithColumn[string]("name", nil)
	// ac := WithColumn[string]("age", nil)
	// ic := WithColumn[string]("id", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token.LowerCase()
		eval := WithLogicalEvaluator[string]()

		eval.Cache("delete", memoryPersist).Delete().From(user).Where(nc.EQ(12)).Eval()
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	var builder strings.Builder
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.WriteString(token.Between.Literal())
		builder.WriteString(token.SpacePlaceholder.Literal())
		builder.WriteString(token.And.Literal())
		builder.WriteString(token.SpacePlaceholder.Literal())
		_ = builder.String()
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{token.Between.Literal(), token.SpacePlaceholder.Literal(), token.And.Literal(), token.SpacePlaceholder.Literal()}, "")
	}
}

func BenchmarkTokenJoin(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = token.Between.Join(token.SpacePlaceholder).Join(token.And).Join(token.SpacePlaceholder).Literal()
	}
}
