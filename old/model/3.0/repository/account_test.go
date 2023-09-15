// Package repository
// @author tabuyos
// @since 2023/9/13
// @description repository
package repository

import (
	"fmt"
	entity2 "metis/old/model/3.0/entity"
	"metis/old/model/3.0/iris"
	"testing"
)

func TestPage(t *testing.T) {
	config := entity2.NewAccount()
	config.Configure(func(eval *iris.Evaluator[entity2.Account]) {
		eval.Select(config.Asterisk()...).From(config.Table()).Page(20, 0).Eval()
	})
	evalInfo := config.Evaluator().EvalInfo()
	fmt.Println(evalInfo)
	fmt.Println(evalInfo.SQL())
	fmt.Println(evalInfo.TotalSQL())
}

func TestJoin1(t *testing.T) {
	roleConfig := entity2.NewRole()
	shipConfig := entity2.NewAccountRole()
	config := entity2.NewAccount()
	config.Configure(func(eval *iris.Evaluator[entity2.Account]) {
		table := config.Table()
		shipTable := shipConfig.Table()
		shipTable.LeftJoin().OnEQ(config.IDCol().Decorate(table.Decorate), shipConfig.AccountIdCol().Decorate(shipTable.Decorate))
		roleTable := roleConfig.Table()
		roleTable.LeftJoin().OnEQ(roleConfig.IDCol().Decorate(roleTable.Decorate), shipConfig.RoleIdCol().Decorate(shipTable.Decorate))
		eval.Select(config.Asterisk(table.Decorate)...).From(table.Ref(shipTable, roleTable)).Where(shipConfig.RoleIdCol().Decorate(shipTable.Decorate).EQ(321)).Eval()
	})

	evalInfo := config.Evaluator().EvalInfo()

	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	fmt.Println(execSQL)
	fmt.Println(values)
}

func TestJoin2(t *testing.T) {
	accountConfig := entity2.NewAccount()
	shipConfig := entity2.NewAccountRole()
	config := entity2.NewRole()
	config.Configure(func(eval *iris.Evaluator[entity2.Role]) {
		table := config.Table()
		shipTable := shipConfig.Table()
		shipTable.LeftJoin().OnEQ(config.IDCol().Decorate(table.Decorate), shipConfig.RoleIdCol().Decorate(shipTable.Decorate))
		accountTable := accountConfig.Table()
		accountTable.LeftJoin().OnEQ(accountConfig.IDCol().Decorate(accountTable.Decorate), shipConfig.AccountIdCol().Decorate(shipTable.Decorate))
		eval.Select(config.Asterisk(table.Decorate)...).From(table.Ref(shipTable, accountTable)).Where(shipConfig.AccountIdCol().Decorate(shipTable.Decorate).EQ(123)).Eval()
	})

	evalInfo := config.Evaluator().EvalInfo()

	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	fmt.Println(execSQL)
	fmt.Println(values)
}
