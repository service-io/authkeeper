// Package entity
// @author tabuyos
// @since 2023/9/8
// @description entity
package entity

import (
	"fmt"
	"metis/model/bei"
	"metis/model/bei/keyword"
	"testing"
)

func TestAccount_Asterisk(t *testing.T) {
	keyword.RegistryCase(false)
	account := NewAccount()
	account.Configure(func(bei *bei.Evaluator[Account]) {
		bei.Select(account.Asterisk()...).From(account.Table()).Where(account.FID().Eq(111000).And(account.FName().Like("tabuyos"), account.FAge().Ge(25))).Eval()
	})
	persist := account.Evaluator().EvalInfo()
	fmt.Println(persist.SQL())
	fmt.Println(persist.TotalSQL())
	fmt.Println(persist.Values())
	fmt.Println(persist.Mappers())
}

type lks struct {
	m map[string]*bei.EvalInfo[Account]
}

func (l *lks) Put(key string, info *bei.EvalInfo[Account]) {
	l.m[key] = info
}

func (l *lks) Get(key string) *bei.EvalInfo[Account] {
	return l.m[key]
}

func BenchmarkAccount_Asterisk(b *testing.B) {
	sqlMap := map[string]*bei.EvalInfo[Account]{}
	memoryPersist := bei.NewMemoryPersist[Account](&lks{m: sqlMap})
	memoryPersist.Persistence("one", bei.OfEvalInfo[Account]("select id, name, age, birthday, gender, deleted from account where (id = ? and name like ? and age >= ?) and deleted = 0;", "", nil, nil))
	keyword.RegistryCase(false)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		account := NewAccount()
		account.Configure(func(bei *bei.Evaluator[Account]) {
			// bei.Replace(memoryPersist.Lookup("one"))
			bei.Select(account.Asterisk()...).From(account.Table()).Where(account.FID().Eq(111000).And(account.FName().Like("tabuyos"), account.FAge().Ge(25))).WithSQLKey("two").Eval(memoryPersist)
		})
		_ = account.Evaluator().EvalInfo()
	}
}
