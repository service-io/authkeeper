// Package iris
// @author tabuyos
// @since 2023/9/11
// @description iris
package iris

import (
	"deepsea/model/iris/internal/token"
	"fmt"
	"testing"
)

func TestInsertEvaluator_Eval(t *testing.T) {
	token.LowerCase()
	// eval := Default[string]()
	eval := WithLogicalEvaluator[string]()
	user := WithTable("user")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)

	eval.Insert(ac, nc, ic).Into(user).Values("1", "2", "3", "1", "2", "3", "1", "2", "3").Eval()
	info := eval.EvalInfo()

	fmt.Printf("rs:> %+v\n", info.SQL())
	fmt.Printf("rs:> %+v\n", info.TotalSQL())
	fmt.Printf("rs:> %+v\n", info.Values())
	fmt.Printf("rs:> %+v\n", info.Mappers())
}

func BenchmarkInsertEvaluator_Eval(b *testing.B) {
	sqlMap := map[string]*EvalInfo[string]{}
	memoryPersist := NewMemoryPersist[string](&lks{m: sqlMap})

	user := WithTable("user")
	nc := WithColumn[string]("name", nil)
	ac := WithColumn[string]("age", nil)
	ic := WithColumn[string]("id", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token.LowerCase()
		eval := WithLogicalEvaluator[string]()

		eval.Cache("insert", memoryPersist).Insert(ac, nc, ic).Into(user).Values("1", "2", "3", "1", "2", "3", "1", "2", "3").Eval()
	}
}
