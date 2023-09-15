// Package iris
// @author tabuyos
// @since 2023/9/12
// @description iris
package iris

import (
	"deepsea/old/model/3.0/iris/internal/token"
	"fmt"
	"testing"
)

func TestUpdateEvaluator_Eval(t *testing.T) {
	token.LowerCase()
	// eval := Default[string]()
	eval := WithLogicalEvaluator[string]()
	user := WithTable("user")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)

	eval.Update(user).Set(nc, 321).Set(ac, 26).Set(ic, 1).Where(nc.EQ(12).Or(ac.EQ(24).And(ic.EQ("111000")))).Eval()
	info := eval.EvalInfo()

	fmt.Printf("rs:> %+v\n", info.SQL())
	fmt.Printf("rs:> %+v\n", info.TotalSQL())
	fmt.Printf("rs:> %+v\n", info.Values())
	fmt.Printf("rs:> %+v\n", info.Mappers())
}

func BenchmarkUpdateEvaluator_Eval(b *testing.B) {
	sqlMap := map[string]*EvalInfo[string]{}
	memoryPersist := NewMemoryPersist[string](&lks{m: sqlMap})

	token.LowerCase()
	// eval := Default[string]()
	user := WithTable("user")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eval := WithLogicalEvaluator[string]()
		eval.Cache("update", memoryPersist).Update(user).Set(nc, 321).Set(ac, 26).Set(ic, 1).Where(nc.EQ(12).Or(ac.EQ(24).And(ic.EQ("111000")))).Eval()
	}
}
