// Package repository
// @author tabuyos
// @since 2023/9/13
// @description repository
package repository

import (
	"metis/helper"
	"metis/helper/database"
	"metis/helper/recorderx"
	entity2 "metis/old/model/3.0/entity"
	iris2 "metis/old/model/3.0/iris"
)

// iAccountRoleAutoGen 该接口自动生成, 请勿修改
type iAccountRoleAutoGen interface {
	SelectManyByConfig(iris2.ConfigService[entity2.Account]) []*entity2.Account
	// SelectAccountByRoleID Query and return the left related entries of XXX id(f key)
	SelectAccountByRoleID(int64) []*entity2.Account
	// SelectRoleByAccountID Query and return the right related entries of XXX id(f key)
	SelectRoleByAccountID(int64) []*entity2.Role
}

// accountRoleAutoGen 该结构体自动生成, 请勿修改
type accountRoleAutoGen struct {
}

func (ag *accountRoleAutoGen) SelectManyByConfig(config iris2.ConfigService[entity2.Account]) []*entity2.Account {
	// TODO implement me
	panic("implement me")
}

func (ag *accountRoleAutoGen) SelectAccountByRoleID(roleID int64) []*entity2.Account {
	recorder := recorderx.DefaultRecorder()
	errorHandler := helper.ErrToLogAndPanic(recorder)

	roleConfig := entity2.NewRole()
	shipConfig := entity2.NewAccountRole()
	config := entity2.NewAccount()
	config.Configure(func(eval *iris2.Evaluator[entity2.Account]) {
		table := config.Table()
		shipTable := shipConfig.Table()
		shipTable.LeftJoin().OnEQ(config.IDCol().Decorate(table.Decorate), shipConfig.AccountIdCol().Decorate(shipTable.Decorate))
		roleTable := roleConfig.Table()
		roleTable.LeftJoin().OnEQ(roleConfig.IDCol().Decorate(roleTable.Decorate), shipConfig.RoleIdCol().Decorate(shipTable.Decorate))
		eval.Select(config.Asterisk(table.Decorate)...).From(table.Ref(shipTable, roleTable)).Where(shipConfig.RoleIdCol().Decorate(shipTable.Decorate).EQ(roleID)).Eval()
	})

	evalInfo := config.Evaluator().EvalInfo()
	if evalInfo == nil {
		return nil
	}

	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	errorHandler(err)
	defer helper.DeferClose(stmt, errorHandler)
	rows, err := stmt.QueryContext(nil, values...)
	errorHandler(err)
	accounts := helper.Rows(rows, func() (*entity2.Account, []any) {
		account := entity2.NewAccount()
		mappers := evalInfo.MapperRows(account)
		return account, mappers
	})

	return accounts
}

func (ag *accountRoleAutoGen) SelectRoleByAccountID(accountID int64) []*entity2.Role {
	recorder := recorderx.DefaultRecorder()
	errorHandler := helper.ErrToLogAndPanic(recorder)

	accountConfig := entity2.NewAccount()
	shipConfig := entity2.NewAccountRole()
	config := entity2.NewRole()
	config.Configure(func(eval *iris2.Evaluator[entity2.Role]) {
		table := config.Table()
		shipTable := shipConfig.Table()
		shipTable.LeftJoin().OnEQ(config.IDCol().Decorate(table.Decorate), shipConfig.RoleIdCol().Decorate(shipTable.Decorate))
		accountTable := accountConfig.Table()
		accountTable.LeftJoin().OnEQ(accountConfig.IDCol().Decorate(accountTable.Decorate), shipConfig.AccountIdCol().Decorate(shipTable.Decorate))
		eval.Select(config.Asterisk(table.Decorate)...).From(table.Ref(shipTable, accountTable)).Where(shipConfig.AccountIdCol().Decorate(shipTable.Decorate).EQ(accountID)).Eval()
	})

	evalInfo := config.Evaluator().EvalInfo()
	if evalInfo == nil {
		return nil
	}

	execSQL := evalInfo.SQL()
	values := evalInfo.Values()

	db := database.FetchDB()
	stmt, err := db.Prepare(execSQL)
	errorHandler(err)
	defer helper.DeferClose(stmt, errorHandler)
	rows, err := stmt.QueryContext(nil, values...)
	errorHandler(err)
	roles := helper.Rows(rows, func() (*entity2.Role, []any) {
		role := entity2.NewRole()
		mappers := evalInfo.MapperRows(role)
		return role, mappers
	})

	return roles
}

var _ iAccountRoleAutoGen = (*accountRoleAutoGen)(nil)
